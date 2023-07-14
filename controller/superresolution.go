/**
 * @Author: lrc
 * @Date: 2023/5/18-23:49
 * @Desc:
 **/

package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"videoStream/service"
	"videoStream/util"
)

func SuRes(c *gin.Context) {
	form, _ := c.MultipartForm()
	IsYoloX := form.Value["IsYoloX"]
	files := form.File["images"]
	imageName := form.Value["imageName"]
	isYoloX, err := strconv.ParseBool(IsYoloX[0])
	if err != nil {
		util.ParamError(c)
		return
	}
	err = service.SuperRes(files, imageName, isYoloX)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.OK(c)
}

func AddPersonBank(c *gin.Context) {
	form, _ := c.MultipartForm()
	personName := form.Value["personName"]
	files := form.File["images"]
	IsYoloX := form.Value["IsYoloX"]

	isYoloX, err := strconv.ParseBool(IsYoloX[0])
	if err != nil {
		util.ParamError(c)
		return
	}

	err = service.CreatePersonBank(files, personName, isYoloX)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.OK(c)

}
