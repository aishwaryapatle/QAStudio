package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/aishwaryapatle/qastudio/internal/auth/model"
	"github.com/aishwaryapatle/qastudio/internal/auth/repository"
	"github.com/aishwaryapatle/qastudio/internal/config"
)

type AuthService struct {
	userRepo *repository.UserRepository
	jwtKey   []byte
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		jwtKey:   []byte(config.Load().JWTSecret),
	}
}

// Signup
func (s *AuthService) Signup(user *model.User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashed)
	return s.userRepo.CreateUser(user)
}

// Login
func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role_id": user.RoleID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(s.jwtKey)
}
