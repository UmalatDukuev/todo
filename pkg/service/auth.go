package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/UmalatDukuev/todo"
	"github.com/UmalatDukuev/todo/pkg/repository"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

type AuthService struct {
	repo repository.Authorization
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
