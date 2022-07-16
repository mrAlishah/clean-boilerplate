package services

import (
	"boilerplate/apps/user/DTO"
	repositories "boilerplate/apps/user/repositories/gorm"
	"boilerplate/core/infrastructures"
	"boilerplate/core/interfaces"
	"boilerplate/core/models"
	"errors"
	"github.com/golang-jwt/jwt"
	"os"
	"strconv"
	"time"
)

// UserService -> struct
type AuthService struct {
	env            *infrastructures.Env
	logger         interfaces.Logger
	userRepository *repositories.UserRepository
	encryption     *infrastructures.Encryption
}

// NewAuthService -> creates a new AuthService
func NewAuthService(
	env *infrastructures.Env,
	logger *infrastructures.Logger,
	userRepository *repositories.UserRepository,
	encryption *infrastructures.Encryption,
) *AuthService {
	return &AuthService{
		env:            env,
		logger:         logger,
		userRepository: userRepository,
		encryption:     encryption,
	}
}

func (s AuthService) CreateAccessToken(user models.User, exp int64, secret string) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["userId"] = user.ID
	atClaims["exp"] = exp
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s AuthService) CreateRefreshToken(user models.User, exp int64, secret string) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["userId"] = user.ID
	atClaims["exp"] = exp
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s AuthService) CreateTokens(user models.User) (map[string]string, error) {

	accessSecret := "access" + os.Getenv("Secret")
	expAccessToken := time.Now().Add(time.Minute * 15).Unix()
	accessToken, err := s.CreateAccessToken(user, expAccessToken, accessSecret)

	refreshSecret := "refresh" + os.Getenv("Secret")
	expRefreshToken := time.Now().Add(time.Hour * 24 * 60).Unix()
	refreshToken, err := s.CreateRefreshToken(user, expRefreshToken, refreshSecret)

	return map[string]string{
		"refreshToken":    refreshToken,
		"accessToken":     accessToken,
		"expRefreshToken": strconv.Itoa(int(expRefreshToken)),
		"expAccessToken":  strconv.Itoa(int(expAccessToken)),
	}, err
}

func (s AuthService) DecodeToken(tokenString string, secret string) (bool, jwt.MapClaims, error) {

	Claims := jwt.MapClaims{}

	key := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			ErrUnexpectedSigningMethod := errors.New("unexpected signing method")
			return nil, ErrUnexpectedSigningMethod
		}
		return []byte(secret), nil
	}

	token, err := jwt.ParseWithClaims(tokenString, Claims, key)
	var valid bool
	if token == nil {
		valid = false
	} else {
		valid = token.Valid
	}
	return valid, Claims, err
}

func (s AuthService) CreateUser(userData DTO.RegisterRequest) (err error) {
	var user models.User
	encryptedPassword := s.encryption.SaltAndSha256Encrypt(userData.Password, userData.Email)
	user.Password = encryptedPassword

	user.FirstName = userData.FirstName
	user.LastName = userData.LastName
	user.Email = userData.Email
	err = s.userRepository.Create(&user)
	if err != nil {
		s.logger.Fatal("Failed to create registered user ", err.Error())
		return
	}
	return
}
