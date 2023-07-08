/**
 * @Author: lrc
 * @Date: 2023/7/8-11:17
 * @Desc:
 **/

package Middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
	"time"
	"videoStream/model"
	"videoStream/util"
)

var accessSecret = []byte("MFwwDQYJKoZIhvcNAQEBBQAD")

func CreateToken(username string) (error, string) {
	claims := model.MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: time.Now().Unix() + 60*60*24,
			Issuer:    "RedRockTeam",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenSigned, err := token.SignedString(accessSecret)
	if err != nil {
		log.Println(err)
		return err, ""
	}
	return nil, tokenSigned
}

func ParseToken(Token string) (*model.MyClaims, error) {
	token, err := jwt.ParseWithClaims(Token, &model.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return accessSecret, nil
	})
	if err != nil {
		log.Println(err)
		return nil, nil
	}

	if claims, ok := token.Claims.(*model.MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("your token is invalid")
}

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			util.VerifyError(c)
			c.Abort()
			return
		}
		parts := strings.Split(authorization, " ")
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			util.VerifyError(c)
			c.Abort()
			return
		}
		Token, err := ParseToken(parts[1])
		if err != nil {
			util.VerifyError(c)
			c.Abort()
			return
		}
		c.Set("username", Token.Username)
		c.Next()
	}
}
