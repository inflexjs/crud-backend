package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/inflexjs/crud-backend/internal/models"
	"github.com/inflexjs/crud-backend/internal/storage"
)

const salt = "idfgeh27348hfg390gh3"

type AuthService struct {
	store storage.Authorization
}

func NewAuthService(store storage.Authorization) *AuthService {
	return &AuthService{store: store}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.store.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
