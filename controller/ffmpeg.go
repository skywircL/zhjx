/**
 * @Author: lrc
 * @Date: 2023/5/16-19:57
 * @Desc:
 **/

package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
	"io"
	"log"
	"sync"
	"time"
	"videoStream/dao"
	"videoStream/model"
	"videoStream/rpc"
	"videoStream/service"
	"videoStream/util"
)

const (
	AbnormalAddress = "127.0.0.1:50053"
)

type wsConn struct {
	inChan    chan model.AbnormalRes
	outChan   chan []byte
	closeChan chan []byte
	isClose   bool // 通道closeChan是否已经关闭
	mutex     sync.Mutex
	conn      *websocket.Conn
}

var abnormal rpc.AbnormalDetection_AbnormalClient

func InitWebSocket(conn *websocket.Conn) (ws *wsConn, err error) {
	ws = &wsConn{
		inChan:    make(chan model.AbnormalRes, 1024),
		outChan:   make(chan []byte, 1024),
		closeChan: make(chan []byte, 1024),
		conn:      conn,
	}
	// 完善必要协程：读取客户端数据协程/发送数据协程
	go ws.readMsgLoop()
	go ws.writeMsgLoop()
	return
}

// readMsgLoop TODO:读取客户端发送的数据写入到inChan
func (conn *wsConn) readMsgLoop() {
	for {
		// 确定数据结构
		var (
			err error
		)
		var abnormal model.AbnormalRes
		err = conn.conn.ReadJSON(&abnormal)
		if err != nil {
			log.Println(err)
			return
		}

		if err = conn.InChanWrite(abnormal); err != nil {
			goto ERR
		}

		//// 接受数据
		//if _, data, err = conn.conn.ReadMessage(); err != nil {
		//	goto ERR
		//}
		//// 写入数据

	}
ERR:
	conn.CloseConn()
}

// writeMsgLoop TODO:读取outChan的数据响应给客户端
func (conn *wsConn) writeMsgLoop() {
	for {
		var (
			data []byte
			err  error
		)
		// 读取数据
		if data, err = conn.OutChanRead(); err != nil {
			goto ERR
		}
		// 发送数据
		if err = conn.conn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.CloseConn()
}

// InChanRead TODO:读取inChan的数据
func (conn *wsConn) InChanRead() (data model.AbnormalRes, err error) {
	select {
	case data = <-conn.inChan:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	default:
		data = model.AbnormalRes{}
	}
	return
}

// InChanWrite TODO:inChan写入数据
func (conn *wsConn) InChanWrite(data model.AbnormalRes) (err error) {
	select {
	case conn.inChan <- data:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}
	return
}

// OutChanRead TODO:读取inChan的数据
func (conn *wsConn) OutChanRead() (data []byte, err error) {
	select {
	case data = <-conn.outChan:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}
	return
}

// OutChanWrite TODO:inChan写入数据
func (conn *wsConn) OutChanWrite(data []byte) (err error) {
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}
	return
}

// CloseConn TODO:关闭WebSocket连接
func (conn *wsConn) CloseConn() {
	// 关闭closeChan以控制inChan/outChan策略,仅此一次
	conn.mutex.Lock()
	if !conn.isClose {
		close(conn.closeChan)
		conn.isClose = true
	}
	conn.mutex.Unlock()
	//关闭WebSocket的连接,conn.Close()是并发安全可以多次关闭
	_ = conn.conn.Close()
}

func FfmpegStream(c *gin.Context) {
	var Ffmpeg model.FfmpegStreamStruct
	err := c.ShouldBindJSON(&Ffmpeg)
	log.Println(Ffmpeg)
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

type returnData struct {
	AnomalyScore      float32
	AnomalyCameraName string
	ImgUrl            string
}

func Abnormal(c *gin.Context) {
	var (
		err  error
		conn *websocket.Conn
		ws   *wsConn
	)
	conn, err = upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("upgrade error:", err)
		return
	}
	if ws, err = InitWebSocket(conn); err != nil {
		log.Println("init websocket error:", err)
		return
	}

	defer conn.Close()
	co, err := grpc.Dial(AbnormalAddress, grpc.WithBlock(), grpc.WithInsecure())
	log.Println("测试点3")
	if err != nil {
		log.Printf("did not connect: %v", err)
	}

	con := rpc.NewAbnormalDetectionClient(co)

	defer co.Close()
	time.Sleep(2 * time.Second)

	for {
		var data model.AbnormalRes
		if data, err = ws.InChanRead(); err != nil {
			log.Println(err)
			return
		}
		if data == (model.AbnormalRes{}) {
			resp, err := abnormal.Recv()
			if err == nil {
				data := returnData{
					AnomalyScore:      resp.GetAnomalyScore(),
					AnomalyCameraName: resp.GetAnomalyCameraName(),
					ImgUrl:            "",
				}
				imagePath := ""
				if data.AnomalyScore > 30 {
					//查数据库，拿出图片路径
					dao.DB.Model(model.AbImgRecord{}).Where("device_name = ?", data.AnomalyCameraName).Last(&imagePath)
					data.ImgUrl = "127.0.0.1:8080/static/" + imagePath
				}

				dataJson, err := json.Marshal(data)
				if err != nil {
					log.Println(err)
				}
				if err = ws.OutChanWrite(dataJson); err != nil {
					log.Println(err)
					return
				}
			}
			if err == io.EOF {
				fmt.Println(err)
				break
			}
			if err != nil {
				log.Println(err)
			}
		} else {
			//todo 先停止 flag
			_, err = con.AbnormalChangeFlag(context.Background(), &rpc.AbnormalFlagParam{
				Flag: true,
			})
			log.Println("测试点1")
			if err != nil {
				log.Println(err)
				return
			}

			abnormal, err = con.Abnormal(context.Background(), &rpc.AbnormalRequest{
				CameraName: data.CameraName,
				RtspUrl:    data.RtspUrl,
			})
			log.Println(abnormal)
			if err != nil {
				log.Println(err)
				return
			}

		}

	}

	//for {
	//	// 接收前端传来的数据
	//	var abnormal model.AbnormalRes
	//	err := conn.ReadJSON(&abnormal)
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//
	//	abnormalCli, err := con.Abnormal(context.Background(), &rpc.AbnormalRequest{
	//		CameraName: abnormal.CameraName,
	//		RtspUrl:    abnormal.RtspUrl,
	//	})
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//
	//	for {
	//		select {
	//		case <-c.Request.Context().Done():
	//
	//		default:
	//			resp, err := abnormalCli.Recv()
	//			if err == nil {
	//				data := returnData{
	//					AnomalyScore:      resp.GetAnomalyScore(),
	//					AnomalyCameraName: resp.GetAnomalyCameraName(),
	//					ImgUrl:            "",
	//				}
	//				imagePath := ""
	//				if data.AnomalyScore > 30 {
	//					//查数据库，拿出图片路径
	//					dao.DB.Model(model.AbImgRecord{}).Where("device_name = ?", data.AnomalyCameraName).Last(&imagePath)
	//					data.ImgUrl = "127.0.0.1:8080/static/" + imagePath
	//				}
	//
	//				dataJson, err := json.Marshal(data)
	//				if err != nil {
	//					log.Println(err)
	//				}
	//				err = conn.WriteMessage(websocket.TextMessage, dataJson)
	//				if err != nil {
	//					log.Println(err)
	//				}
	//
	//			}
	//			if err == io.EOF {
	//				fmt.Println(err)
	//				break
	//			}
	//			if err != nil {
	//				log.Println(err)
	//			}
	//		}
	//
	//	}

	//err = rpc.Abnormal(conn, abnormal.CameraName, abnormal.RtspUrl)
	//if err != nil {
	//	util.HandleError(c, err)
	//	return
	//}

}

//device, err := dao.QueryDeviceDisPlaying()
//if err != nil {
//	return
//}
//var rtspUrls []string
//var CameraName []string
//for _, v := range device {
//	rtsp := ""
//	if v.DeviceLocation == "实验室" {
//		rtsp = "rtsp://lengjx:ljx588588@" + v.Ip + "/Streaming/Channels/" + v.Index + "?transportmode=multicast"
//	}
//	if v.DeviceLocation == "学校" {
//		// todo 记得要改
//		rtsp = "rtsp://lengjx:ljx588588@" + v.Ip + "/Streaming/Channels/1502?transportmode=multicast"
//	}
//	rtspUrls = append(rtspUrls, rtsp)
//	CameraName = append(CameraName, v.DeviceName)
//}

//err = rpc.Abnormal(conn, CameraName, rtspUrls)
//if err != nil {
//	util.HandleError(c, err)
//	return
//}
