/**
 * @Author: lrc
 * @Date: 2023/5/13-11:38
 * @Desc:
 **/

package rpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
	"videoStream/dao"
	"videoStream/model"
	"videoStream/util"
)

const (
	Address = "10.16.50.17:50051"
	//Address         = "127.0.0.1:50051"
	abnormalAddress = "10.16.50.17:50053"
	//abnormalAddress = "127.0.0.1:50053"
)

// Ffmpeg todo 将入参改为只用传设备名称即可
func Ffmpeg(tx *gorm.DB, Ffmpeg []string) (pushStreamUrl []string, savePath []string, err error) {

	conn, err := grpc.Dial(Address, grpc.WithBlock(), grpc.WithInsecure())
	var rtspUrls []string
	var rtmpUrls []string
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := NewFfmpegClient(conn)

	//todo 先从数据库里拿数据  deviceName到数据库里拿ip ,stream,loc
	for _, subFfm := range Ffmpeg {
		var device model.DeviceIp
		dao.DB.Model(model.DeviceIp{}).Where("ip = ?", subFfm).Find(&device)

		if device.DeviceLocation == "实验室" {
			rtspUrl := "rtsp://lengjx:ljx588588@" + device.Ip + "/Streaming/Channels/" + device.Index + "?transportmode=multicast"
			rtmpUrl := "rtmp://127.0.0.1:1935/live/reid_stream" + strconv.Itoa(int(device.ID))
			rtspUrls = append(rtspUrls, rtspUrl)
			rtmpUrls = append(rtmpUrls, rtmpUrl)
		}

		savePa := "../yolox/per_img/" + util.GetGID()
		savePath = append(savePath, savePa)

		//FlagResp, err := c.ChangeFfmpegFlag(context.Background(), &FlagParam{
		//	Flag: false,
		//})
		//if err != nil {·
		//	return nil, nil, err
		//}
		//
		//if !FlagResp.Error {
		//	return nil, nil, errors.New(FlagResp.Message)
		//}

		r, err := c.PersonDetection(context.Background(), &CameraIp{
			CameraIp: device.Ip,
			Location: device.DeviceLocation,
			Stream:   device.Index,
			SavePath: savePa,
		},
		)

		if err != nil {
			return nil, nil, err
		}
		fmt.Println(r.Message, r.Error)
	}

	//异步执行这一步,不然会阻塞在这里
	go func() {
		_, err = c.VideoStream(context.Background(), &VideoStreamStruct{
			RtspUrl: rtspUrls,
			RtmpUrl: rtmpUrls,
		})
		if err != nil {
			return
		}
	}()

	time.Sleep(1 * time.Second) //用于等待rpc是否出现异常

	if err != nil {
		log.Println(err)
		return
	}

	//无错误改数据库,这里改数据库是为了记录摄像头推流状态
	for _, name := range Ffmpeg {
		//更新status为1，表示显示状态
		err := tx.Model(model.FfmpegStatus{}).Where("camera_name=?", name).Update("status", "显示").Error
		if err != nil {
			tx.Rollback()
			return nil, nil, err
		}

	}
	//
	//	if pid.ID == 0 {
	//		//todo 不存在，直接推流
	//		pid.CameraName = name
	//		pid.Status = "1"
	//		pid.Pid = streamInfo.Pid[i]
	//		err = tx.Create(&pid).Error
	//		if err != nil {
	//			log.Println(err)
	//			tx.Rollback()
	//			return nil, nil, err
	//		}
	//	}
	//}
	//
	//for _, pid := range quitPid {
	//	//todo 存在，先杀进程，再推流
	//	IntPid, err := strconv.Atoi(pid.Pid)
	//	if err != nil {
	//		tx.Rollback()
	//		return nil, nil, err
	//	}
	//
	//	process, err := os.FindProcess(IntPid)
	//	if err != nil {
	//		tx.Rollback()
	//		return nil, nil, err
	//	}
	//	err = process.Kill()
	//	if err != nil {
	//		tx.Rollback()
	//		return nil, nil, err
	//	}
	//	err = tx.Model(model.FfmpegStatus{}).Where("camera_name=?", pid.CameraName).Delete(&pid).Error
	//	if err != nil {
	//		tx.Rollback()
	//		return nil, nil, err
	//	}
	//
	//}
	tx.Commit()
	pushStreamUrl = rtmpUrls
	return
}

func ChangeFlag(flag bool) (bool, string) {
	conn, err := grpc.Dial(Address, grpc.WithBlock(), grpc.WithInsecure())
	defer conn.Close()
	conn2, err := grpc.Dial(abnormalAddress, grpc.WithBlock(), grpc.WithInsecure())
	defer conn2.Close()
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := NewFfmpegClient(conn)

	c2 := NewAbnormalDetectionClient(conn2)

	changeFlag, err := c.ChangeFfmpegFlag(context.Background(), &FlagParam{
		Flag: flag,
	})
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	changeFlag2, err := c2.AbnormalChangeFlag(context.Background(), &AbnormalFlagParam{
		Flag: flag,
	})
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	if !changeFlag2.Error {
		return changeFlag2.Error, changeFlag2.Message
	}

	return changeFlag.Error, changeFlag.Message
}
