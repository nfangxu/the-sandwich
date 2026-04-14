package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/the-sandwich/backend/internal/domain/user"
)

type Service struct {
	userRepo  user.UserRepository
	jwtSecret string
}

func NewAuthService(userRepo user.UserRepository, jwtSecret string) *Service {
	return &Service{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (s *Service) Register(email, username, password string) (*user.User, error) {
	// Check if user already exists
	existing, _ := s.userRepo.FindByEmail(email)
	if existing != nil {
		return nil, fmt.Errorf("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	u := &user.User{
		Email:    email,
		Username: username,
		Password: string(hashedPassword),
	}

	if err := s.userRepo.Create(u); err != nil {
		return nil, err
	}

	return u, nil
}

func (s *Service) Login(email, password string) (string, error) {
	u, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": u.ID,
		"email":   u.Email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return tokenString, nil
}

func (s *Service) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
