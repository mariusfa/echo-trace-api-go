package biz

import (
	"echo/biz/domain"
	"errors"
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository UserRepositoryContract
}

func (us *UserService) Register(userRequest domain.UserRequest) error {
	_, err := us.UserRepository.GetByName(userRequest.Username)
	if err == nil {
		return errors.New("User already exists")
	}

	hashedPassword, err := hashPassword(userRequest.Password)
	if err != nil {
		return err
	}

	user := domain.User{
		Name:           userRequest.Username,
		HashedPassword: hashedPassword,
		ApiToken:       generateApiToken(),
	}
	us.UserRepository.Insert(user)
	return nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func generateApiToken() string {
	allowedChars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 32)
	for i := range b {
		b[i] = allowedChars[rand.Intn(len(allowedChars))]
	}
	return string(b)
}

func (us *UserService) Login(userRequest domain.UserRequest) error {
	user, err := us.UserRepository.GetByName(userRequest.Username)
	if err != nil {
		return errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(userRequest.Password))
	if err != nil {
		return errors.New("invalid credentials")
	}

	// TODO: return token

	return nil
}
