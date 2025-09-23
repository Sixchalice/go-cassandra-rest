package repository

import (
	"context"
	"log/slog"

	"github.com/Sixchalice/go-cassandra-rest/internal/models"
	"github.com/gocql/gocql"
)

type UserRepository struct {
	// You might inject a Cassandra session or other dependency here
	session *gocql.Session
}

func NewUserRepository(session *gocql.Session) *UserRepository {
	return &UserRepository{session: session}
}

func (r *UserRepository) InsertUser(ctx context.Context, id, name string) error {
	// TODO: implement Cassandra insert
	err := r.session.Query("INSERT INTO users (id, name) VALUES (?, ?)", id, name).WithContext(ctx).Exec()
	if err != nil {
		slog.Error("failed to insert user", "error", err)
		return err
	}
	return nil
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]models.User, error) {
	// TODO: implement Cassandra select all
	iter := r.session.Query("SELECT id, name FROM users").WithContext(ctx).Iter()
	var users []models.User
	var u models.User

	for iter.Scan(&u.ID, &u.Name) {
		users = append(users, u)
	}

	if err := iter.Close(); err != nil {
		slog.Error("failed to close iterator", "error", err)
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	var u models.User
	err := r.session.Query("SELECT id, name FROM users WHERE id = ?", id).WithContext(ctx).Scan(&u.ID, &u.Name)
	if err != nil {
		if err == gocql.ErrNotFound {
			return nil, nil // User not found
		}
		slog.Error("failed to get user by ID", "error", err)
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, id, name string) error {
	err := r.session.Query("UPDATE users SET name = ? WHERE id = ?", name, id).WithContext(ctx).Exec()
	if err != nil {
		slog.Error("failed to update user", "error", err)
		return err
	}
	return nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id string) error {
	err := r.session.Query("DELETE FROM users WHERE id = ?", id).WithContext(ctx).Exec()
	if err != nil {
		slog.Error("failed to delete user", "error", err)
		return err
	}
	return nil
}
