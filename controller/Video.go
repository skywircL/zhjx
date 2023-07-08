/**
 * @Author: lrc
 * @Date: 2023/6/25-18:56
 * @Desc:
 **/

package controller

import (
	"github.com/gin-gonic/gin"
	"videoStream/service"
	"videoStream/util"
)

func GetVideo(c *gin.Context) {
	//获取参数
	type video struct {
		StartTime    string   `json:"start_time"`
		EndTime      string   `json:"end_time"`
		DeviceName   string   `json:"device_name"`
		PeopleImages []string `json:"people_images"`
	}

	var v video

	err := c.ShouldBindJSON(&v)
	if err != nil {
		util.ParamError(c)
		return
	}
	//调用service,先查寻该参数下有没有异常，根据图片库人名
	rc, err := service.GetVideo(v.DeviceName, v.PeopleImages, v.StartTime, v.EndTime)
	if err != nil {
		util.HandleError(c, err)
		return
	}

	util.OKWithData(c, rc)

}

// GetVideoStream ，拿到设备，时间，推流
func GetVideoStream(c *gin.Context) {
	//获取参数
	type video struct {
		StartTime  string `json:"start_time"`
		EndTime    string `json:"end_time"`
		DeviceName string `json:"device_name"`
	}

	var v video

	err := c.ShouldBindJSON(&v)
	if err != nil {
		util.ParamError(c)
		return
	}
	//调用service,推流返回
	stream, err := service.VideoStream(v.StartTime, v.EndTime, v.DeviceName)
	if err != nil {
		util.HandleError(c, err)
		return
	}

	util.OKWithData(c, stream)

}

// QuitVideoStream 清空接口
func QuitVideoStream(c *gin.Context) {
	//调用rpc，修改flag值
	err := service.QuitVideoStream()
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.OK(c)
}

// GetPersonTrackRecord 查询所有人物跟踪记录
func GetPersonTrackRecord(c *gin.Context) {
	record, err := service.GetPersonTrackRecord()
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.OKWithData(c, record)
}

// GetPersonTrackRecordByCondition 根据人物，设备，时间查询人物跟踪记录
func GetPersonTrackRecordByCondition(c *gin.Context) {

}
