package pkg

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type MetaToken struct {
	ID            uint
	Username      string
	Password      string
	ExpiredAt     time.Time
	Authorization bool
}

type AccessToken struct {
	Claims MetaToken
}

func Sign(Data map[string]interface{}, SecretPublicKeyEnvName string, ExpiredAt time.Duration, userId uint) (string, error) {

	expiredAt := time.Now().Add(time.Duration(time.Minute) * ExpiredAt).Unix()

	jwtSecretKey := GodotEnv(SecretPublicKeyEnvName)

	claims := jwt.MapClaims{}
	claims["exp"] = expiredAt
	claims["ID"] = userId
	claims["authorization"] = true

	for i, v := range Data {
		claims[i] = v
	}

	to := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := to.SignedString([]byte(jwtSecretKey))

	if err != nil {
		logrus.Error(err.Error())
		return accessToken, err
	}

	return accessToken, nil
}

func VerifyTokenHeader(ctx *gin.Context, SecretPublicKeyEnvName string) (*jwt.Token, error) {
	tokenHeader := ctx.GetHeader("Authorization")
	accessToken := strings.SplitAfter(tokenHeader, "Bearer")[1]
	jwtSecretKey := GodotEnv(SecretPublicKeyEnvName)

	token, err := jwt.Parse(strings.Trim(accessToken, " "), func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return token, nil
}

func VerifyToken(accessToken, SecretPublicKeyEnvName string) (*jwt.Token, error) {
	jwtSecretKey := GodotEnv(SecretPublicKeyEnvName)

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return token, nil
}

func DecodeToken(accessToken *jwt.Token) AccessToken {
	var token AccessToken
	stringify, _ := json.Marshal(&accessToken)
	json.Unmarshal([]byte(stringify), &token)

	return token
}
