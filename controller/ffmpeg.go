/**
 * @Author: lrc
 * @Date: 2023/5/16-19:57
 * @Desc:
 **/

package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"videoStream/model"
	"videoStream/rpc"
	"videoStream/service"
	"videoStream/util"
)

func FfmpegStream(c *gin.Context) {
	var Ffmpeg model.FfmpegStreamStruct
	err := c.ShouldBindJSON(&Ffmpeg)
	if err != nil {
		util.ParamError(c)
		return
	}

	pushStreamUrl, err := service.FfmpegStream(Ffmpeg)
	if err != nil {
		util.HandleError(c, err)
		return
	}

	util.OKWithData(c, pushStreamUrl)

}

func Quit(c *gin.Context) {
	err := service.QuitFfmpeg()
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.OK(c)
}

type AbnormalJson struct {
	AnomalyScore      float32
	AnomalyCameraName string
}

func Abnormal(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("upgrade error:", err)
		return
	}

	defer conn.Close()
	for {
		// 接收前端传来的数据
		var abnormal model.AbnormalRes
		err := conn.ReadJSON(&abnormal)
		if err != nil {
			log.Println(err)
			return
		}

		err = rpc.Abnormal(conn, abnormal.CameraName, abnormal.RtspUrl)
		if err != nil {
			util.HandleError(c, err)
			return
		}

		util.OK(c)

	}

}
