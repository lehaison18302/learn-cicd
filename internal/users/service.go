package users

import (
	"context"
	"fmt"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// GetUsersWithPagination returns paginated users
func (s *Service) GetUsersWithPagination(ctx context.Context, page, pageSize int) ([]User, error) {
	return s.repo.GetUsers(ctx, page, pageSize)
}

func (s *Service) CreateNewUser(ctx context.Context, user User) error {
	return s.repo.CreateUser(ctx, user)
}

func (s *Service) UpdateUser(ctx context.Context, user User) error {
	return s.repo.UpdateUser(ctx, user)
}

func (s *Service) DeleteUser(ctx context.Context, id int) error {
	return s.repo.DeleteUser(ctx, id)
}

func (s *Service) Login(ctx context.Context, username, password string) (*User, error) {
	user, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if user.Password != password {
		return nil, fmt.Errorf("incorrect password")
	}
	return user, nil
}
