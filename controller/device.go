/**
 * @Author: lrc
 * @Date: 2023/5/28-9:50
 * @Desc:
 **/

package controller

import (
	"github.com/gin-gonic/gin"
	"videoStream/service"
	"videoStream/util"
)

func AddDevice(c *gin.Context) {

	var Device struct {
		Ip             string `json:"ip" binding:"required"`
		DeviceName     string `json:"device_name" binding:"required"`
		DeviceLocation string `json:"device_location" binding:"required"`
		Index          string `json:"index" binding:"required"`
	}

	err := c.ShouldBindJSON(&Device)
	if err != nil {
		util.ParamError(c)
		return
	}

	err = service.AddDevice(Device.DeviceName, Device.DeviceLocation, Device.Ip, Device.Index)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.OK(c)
}

func GetDevice(c *gin.Context) {
	device, err := service.SearchDevice()
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.OKWithData(c, device)
}

func DeleteDevice(c *gin.Context) {
	var Device struct {
		DeviceName string `json:"device_name" binding:"required"`
	}
	err := c.ShouldBindJSON(&Device)
	if err != nil {
		util.ParamError(c)
		return
	}

	err = service.DeleteDevice(Device.DeviceName)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.OK(c)
}

func ModifyDevice(c *gin.Context) {
	var Device struct {
		Ip             string `json:"ip" binding:"required"`
		DeviceName     string `json:"device_name" binding:"required"`
		DeviceLocation string `json:"device_location" binding:"required"`
		Index          string `json:"index" binding:"required"`
	}
	err := c.ShouldBindJSON(&Device)
	if err != nil {
		util.ParamError(c)
		return
	}

	err = service.UpdateDevice(Device.DeviceName, Device.DeviceLocation, Device.Ip, Device.Index)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.OK(c)
}
