package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

var KeyFunc = func(token *jwt.Token) (interface{}, error) {
	return []byte(viper.GetString("auth-token")), nil
}

var Method = jwt.SigningMethodHS256
