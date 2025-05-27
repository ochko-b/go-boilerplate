package services

import (
	"context"
	"fmt"
	"time"

	"github.com/ochko-b/goapp/generated/sqlc"
	"github.com/ochko-b/goapp/internal/config"
	"github.com/ochko-b/goapp/internal/models"
	"github.com/ochko-b/goapp/internal/repository"
	"github.com/ochko-b/goapp/internal/utils"
)

type AuthService struct {
	repo      *repository.Repository
	jwtConfig config.JWTConfig
}

func NewAuthService(repo *repository.Repository, jwtConfig config.JWTConfig) *AuthService {
	return &AuthService{
		repo:      repo,
		jwtConfig: jwtConfig,
	}
}

func (s *AuthService) Register(ctx context.Context, req *models.RegisterRequest) (*models.UserResponse, string, error) {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, "", err
	}

	user, err := s.repo.CreateUser(ctx, sqlc.CreateUserParams{
		Email:        req.Email,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		PasswordHash: hashedPassword,
	})
	if err != nil {
		return nil, "", err
	}

	duration, _ := time.ParseDuration(s.jwtConfig.ExpiresIn)
	token, err := utils.GenerateToken(user.ID.String(), user.Email, s.jwtConfig.Secret, duration)
	if err != nil {
		return nil, "", err
	}

	userResponse := &models.UserResponse{
		ID:        user.ID.String(),
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Time.Format(time.RFC3339),
	}

	return userResponse, token, nil
}

func (s *AuthService) Login(ctx context.Context, req *models.LoginRequest) (*models.UserResponse, string, error) {
	user, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, "", err
	}

	if !utils.CheckPassword(req.Password, user.PasswordHash) {
		return nil, "", fmt.Errorf("invalid credentials")
	}

	duration, _ := time.ParseDuration(s.jwtConfig.ExpiresIn)
	token, err := utils.GenerateToken(user.ID.String(), user.Email, s.jwtConfig.Secret, duration)
	if err != nil {
		return nil, "", err
	}

	userResponse := &models.UserResponse{
		ID:        user.ID.String(),
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Time.Format(time.RFC3339),
	}

	return userResponse, token, nil
}

func (s *AuthService) RefreshToken(userID, email string) (string, error) {
	duration, _ := time.ParseDuration(s.jwtConfig.ExpiresIn)
	return utils.GenerateToken(userID, email, s.jwtConfig.Secret, duration)
}
