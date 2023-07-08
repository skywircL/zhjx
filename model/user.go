/**
 * @Author: lrc
 * @Date: 2023/7/8-11:16
 * @Desc:
 **/

package model

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
