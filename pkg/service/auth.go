package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	models "github.com/Njrctr/javacode_test_golang_junior/models"
	"github.com/Njrctr/javacode_test_golang_junior/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "jh23hriuhq9ah9823334wer{?>wre"
	signingKey = "fsdglk;asd3j41lqwwkj42134jzxca"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId  int  `json:"user_id"`
	IsAdmin bool `json:"is_admin"`
}

type AuthService struct {
	repo repository.Autorization
}

func NewAuthService(repo repository.Autorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.SignUpInput) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateJWTToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
		user.IsAdmin,
	})
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accesToken string) (int, bool, error) {
	token, err := jwt.ParseWithClaims(accesToken, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, false, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, false, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, claims.IsAdmin, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
