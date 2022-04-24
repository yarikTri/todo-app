package service

import (
	"crypto/sha256"
	"fmt"

	"github.com/yarikTri/todo-app"
	"github.com/yarikTri/todo-app/package/repository"
)

// 'salt' is set of random symbols to make hash-password stronger
const salt = "m98c*#PO4qx^&B*DXn7r3)"

// AuthService ...
type AuthService struct {
	repo repository.Authorization
}

// NewAuthService ...
func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// CreateUser ...
func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = s.generateHashPassword(user.Password)

	return s.repo.CreateUser(user)
}

func (s *AuthService) generateHashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
