/**
 * @Author: lrc
 * @Date: 2023/4/13-10:42
 * @Desc:
 **/

package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"videoStream/controller"
)

func RunRouter() {
	r := gin.Default()
	r.StaticFS("/static", http.Dir("../absaveimg"))
	r.StaticFS("/img", http.Dir("../../PersonImageDatabase"))
	r.POST("/monitor/super/resolution", controller.SuRes) //超分接口
	//一个非超分的接口

	r.POST("/monitor/add/images", controller.AddPersonData)         //添加人物
	r.GET("/monitor/images", controller.GetImageDatabase)           //返回图像库,这个是查数据库
	r.DELETE("/monitor/delete/images", controller.DeletePersonData) //删除选中人物库

	r.GET("/monitor/portrait", controller.WsGETImage)                   //ws 返回人像图片
	r.GET("/monitor/ffmpeg", controller.FfmpegStream)                   //推流，异步进行截图ws返回图片（在这里异步调用返回图片的rpc服务
	r.DELETE("/monitor/ffmpeg/quit", controller.Quit)                   //停止推流和截图
	r.DELETE("/monitor/ffmpeg/person/quit", controller.QuitPersonTrack) //停止追踪人物

	//异常检测
	r.GET("/monitor/abnormal", controller.Abnormal) //获取异常检测的结果

	//添加设备
	r.POST("/monitor/add/device", controller.AddDevice)
	//获取设备列表
	r.GET("/monitor/device", controller.GetDevice)
	//删除设备
	r.DELETE("/monitor/delete/device", controller.DeleteDevice)
	//修改设备
	r.PUT("/monitor/modify/device", controller.ModifyDevice)

	//搜索设备 要使用sdk

	//r.DELETE("/monitor/ffmpeg/quit", controller.Quit) //注销登录，释放资源

	//回溯视频筛选
	r.GET("/monitor/video", controller.GetVideo)

	//回溯视频推流
	r.GET("/monitor/video/stream", controller.GetVideoStream)

	//回溯视频重置，关停回溯资源
	r.DELETE("/monitor/video/stream/quit", controller.QuitVideoStream)

	//异常记录查询
	r.GET("/monitor/abnormal/record", controller.GetAbnormalRecord)

	//人物跟踪记录查询
	r.GET("/monitor/person/track/record", controller.GetPersonTrackRecord)

	//登录接口
	r.POST("/user/login", controller.Login)

	//退出登录释放资源接口
	r.DELETE("/user/logout", controller.Logout)

	//非超分生成personBank接口
	r.POST("/monitor/add/personBank", controller.AddPersonBank)

	_ = r.Run(":8080")
}
