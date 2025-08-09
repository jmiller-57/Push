package auth

import (
	"errors"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	secret string
	mu     sync.Mutex
	users  map[string]*user
}

type user struct {
	ID    string
	Email string
	Hash  []byte
}

func NewService(secret string) *Service {
	return &Service{secret: secret, users: map[string]*user{}}
}

func (s *Service) Register(email, password string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.users[email]; ok {
		return "", errors.New("email already exists")
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	u := &user{ID: "u_" + email, Email: email, Hash: hash}
	s.users[email] = u
	return s.sign(u.ID)
}

func (s *Service) Login(email, password string) (string, error) {
	s.mu.Lock()
	u, ok := s.users[email]
	s.mu.Unlock()
	if !ok || bcrypt.CompareHashAndPassword(u.Hash, []byte(password)) != nil {
		return "", errors.New("invalid credentials")
	}
	return s.sign(u.ID)
}

func (s *Service) sign(userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(s.secret))
}

func (s *Service) Verify(token string) (string, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil
	})
	if err != nil || !t.Valid {
		return "", errors.New("invalid token")
	}
	if claims, ok := t.Claims.(jwt.MapClaims); ok {
		if sub, ok := claims["sub"].(string); ok {
			return sub, nil
		}
	}
	return "", errors.New("no sub")
}
