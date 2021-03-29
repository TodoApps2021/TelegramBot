package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/todoapps2021/telegrambot/internal/repository"
)

const (
	SALT = "saltcode"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (as *AuthService) Login(username, password string, telegramId int) error {
	hashPassword, err := generatePasswordHash(password)
	if err != nil {
		return err
	}

	password = hashPassword
	return as.repo.Login(username, password, telegramId)
}

func generatePasswordHash(password string) (string, error) {
	hash := sha1.New()
	if _, err := hash.Write([]byte(password)); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum([]byte(SALT))), nil
}

func (as *AuthService) CheckUser(username, password string) (bool, error) {
	hashPassword, err := generatePasswordHash(password)
	if err != nil {
		return false, err
	}

	password = hashPassword
	logrus.Info(password)
	return as.repo.CheckUser(username, password)
}

func (as *AuthService) GetUserIdByTgId(telegram_id int) (int, error) {
	return as.repo.GetUserIdByTgId(telegram_id)
}
