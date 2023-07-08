/**
 * @Author: lrc
 * @Date: 2023/7/8-11:18
 * @Desc:
 **/

package controller

import (
	"github.com/gin-gonic/gin"
	"videoStream/Middleware"
	"videoStream/service"
	"videoStream/util"
)

func Login(c *gin.Context) {
	var UserJson struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := c.ShouldBindJSON(&UserJson)
	if err != nil {
		util.ParamError(c)
		return
	}

	username := UserJson.Username
	password := UserJson.Password
	if username == "" || password == "" {
		util.ParamError(c)
		return
	}

	exist := service.JudgeUserExist(username, password)

	var token string
	if exist {
		err, token = Middleware.CreateToken(username)
		if err != nil {
			util.HandleError(c, err)
			return
		}
	} else {
		util.VerifyError(c)
		return
	}

	util.OKWithData(c, token)

}
