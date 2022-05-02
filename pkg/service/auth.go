package service

import (
	"crypto/ecdsa"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/amper-pw/auth/pkg/model"
	"github.com/amper-pw/auth/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"os"
	"time"
)

const (
	salt     = "hjqrhjqw124617ajfhajs"
	tokenTTL = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}

type AuthService struct {
	repository repository.UserRepositoryInterface
}

func (s *AuthService) ParseToken(accessToken string) (string, error) {

	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, errors.New("invalid signing method")
		}

		key, _ := ioutil.ReadFile(os.Getenv("PUBLIC_KEY_PATH"))

		var ecdsaKey *ecdsa.PublicKey
		var err error
		if ecdsaKey, err = jwt.ParseECPublicKeyFromPEM(key); err != nil {
			return "", errors.New(fmt.Sprintf("Unable to parse ECDSA private key: %v", err))
		}

		return ecdsaKey, nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserId, nil
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repository.FindUserByUsernameAndPassword(username, generatePasswordHash(password))

	if err != nil {
		return "", err
	}

	key, _ := ioutil.ReadFile(os.Getenv("PRIVATE_KEY_PATH"))

	var ecdsaKey *ecdsa.PrivateKey
	if ecdsaKey, err = jwt.ParseECPrivateKeyFromPEM(key); err != nil {
		return "", errors.New(fmt.Sprintf("Unable to parse ECDSA private key: %v", err))
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id.String(),
	})
	return token.SignedString(ecdsaKey)
}

func (s *AuthService) RegisterUser(username, password string) (*model.User, error) {
	return s.repository.Create(username, generatePasswordHash(password))
}

func BuildAuthService(repos repository.UserRepositoryInterface) *AuthService {
	return &AuthService{repository: repos}
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
