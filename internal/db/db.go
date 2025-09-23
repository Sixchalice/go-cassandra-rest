package db

import (
	"fmt"
	"log/slog"

	"github.com/Sixchalice/go-cassandra-rest/internal/constants"
	"github.com/gocql/gocql"
)

func InitDB(databaseUri string, keyspace string) (*gocql.Session, func(), error) {

	cluster := gocql.NewCluster(databaseUri)
	cluster.Keyspace = constants.SystemKeyspace
	cluster.Consistency = gocql.Quorum

	session, err := cluster.CreateSession()
	if err != nil {
		slog.Error("could not connect to Cassandra", "error", err)
		return nil, nil, err
	}
	defer session.Close()

	createKeyspace := fmt.Sprintf(`CREATE KEYSPACE IF NOT EXISTS %s 
		WITH replication = {'class': 'SimpleStrategy', 
		'replication_factor': '1'}`,
		keyspace)

	if err := session.Query(createKeyspace).Exec(); err != nil {
		slog.Error("failed to create keyspace", "error", err)
		return nil, nil, err
	}

	slog.Info("Keyspace created or already exists", "keyspace", keyspace)

	cluster.Keyspace = keyspace
	keyspaceSession, err := cluster.CreateSession()
	if err != nil {
		slog.Error("could not connect to new keyspace", "error", err)
		return nil, nil, err
	}
	cleanup := func() {
		keyspaceSession.Close()
	}

	fmt.Println("Connected to new keyspace successfully.")

	return keyspaceSession, cleanup, nil
}
