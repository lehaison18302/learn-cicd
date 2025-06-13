package users

import (
	"context"
	"fmt"
	"my_project/internal/database"
)

// Repository provides access to user storage.
type Repository struct {
	db database.Service
}

func NewRepository(db database.Service) *Repository {
	return &Repository{db: db}
}

// GetUsers returns a paginated list of users (without password)
func (r *Repository) GetUsers(ctx context.Context, page, pageSize int) ([]User, error) {
	offset := (page - 1) * pageSize
	query := fmt.Sprintf("SELECT id, username, fullname FROM users LIMIT %d OFFSET %d", pageSize, offset)
	rows, err := r.db.GetDB().QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Username, &u.Fullname); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

// create user
func (r *Repository) CreateUser(ctx context.Context, user User) error {
	query := "INSERT INTO users (username, fullname) VALUES (?, ?)"
	_, err := r.db.GetDB().ExecContext(ctx, query, user.Username, user.Fullname)
	if err != nil {
		return err
	}
	return nil
}

// Update user
func (r *Repository) UpdateUser(ctx context.Context, user User) error {
	query := "UPDATE users SET username = ?, fullname = ? WHERE id = ?"
	_, err := r.db.GetDB().ExecContext(ctx, query, user.Username, user.Fullname, user.ID)
	if err != nil {
		return err
	}
	return nil
}

// Delete user
func (r *Repository) DeleteUser(ctx context.Context, id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.db.GetDB().ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

// GetUserByUsername returns a user by username (for login)
func (r *Repository) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	query := "SELECT id, username, fullname, password FROM users WHERE username = ? LIMIT 1"
	row := r.db.GetDB().QueryRowContext(ctx, query, username)
	var u User
	var password string
	if err := row.Scan(&u.ID, &u.Username, &u.Fullname, &password); err != nil {
		return nil, err
	}
	u.Password = password
	return &u, nil
}
