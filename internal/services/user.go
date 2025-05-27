package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/ochko-b/goapp/generated/sqlc"
	"github.com/ochko-b/goapp/internal/models"
	"github.com/ochko-b/goapp/internal/repository"
)

type UserService struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetByID(ctx context.Context, userID string) (*models.UserResponse, error) {
	parsedID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	pgUUID := pgtype.UUID{
		Bytes: parsedID,
		Valid: true,
	}
	user, err := s.repo.GetUserByID(ctx, pgUUID)
	if err != nil {
		return nil, err
	}

	return &models.UserResponse{
		ID:        user.ID.String(),
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Time.Format(time.RFC3339),
	}, nil
}

func (s *UserService) UpdateProfile(ctx context.Context, userID string, req *models.UpdateProfileRequest) (*models.UserResponse, error) {
	parsedID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	pgUUID := pgtype.UUID{
		Bytes: parsedID,
		Valid: true,
	}

	user, err := s.repo.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:        pgUUID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	})
	if err != nil {
		return nil, err
	}

	return &models.UserResponse{
		ID:        user.ID.String(),
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Time.Format(time.RFC3339),
	}, nil
}

func (s *UserService) List(ctx context.Context, limit, offset int32) ([]*models.UserResponse, error) {
	users, err := s.repo.ListUsers(ctx, sqlc.ListUsersParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	var userResponses []*models.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, &models.UserResponse{
			ID:        user.ID.String(),
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			CreatedAt: user.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt: user.UpdatedAt.Time.Format(time.RFC3339),
		})
	}
	return userResponses, nil
}
