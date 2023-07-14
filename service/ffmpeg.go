/**
 * @Author: lrc
 * @Date: 2023/5/13-21:50
 * @Desc:
 **/

package service

import (
	"errors"
	"log"
	"os"
	"videoStream/dao"
	"videoStream/model"
	"videoStream/rpc"
)

//func FfmpegStream(ffmpeg model.FfmpegStreamStruct) (pushStreamUrl []string, err error) {
//	//todo 先查库，然后重复的保留，多的推流，没有的直接kill  ip加1502查
//
//	var status []model.FfmpegStatus
//	err = dao.DB.Model(model.FfmpegStatus{}).Where("status=?", 1).Find(&status).Error
//	if err != nil {
//		log.Println(err)
//		return
//	}
//
//	var quitPid []model.FfmpegStatus
//	for _, qName := range status {
//		flag := 0
//		for _, FName := range ffmpeg.CameraName {
//			if qName.CameraName == FName {
//				flag = 1
//			}
//		}
//		if flag == 0 {
//			//没找到直接给他停了
//			quitPid = append(quitPid, qName)
//		}
//	}
//
//	var pids []string
//
//	for _, name := range ffmpeg.CameraName {
//		var pid model.FfmpegStatus
//		err = dao.DB.Model(model.FfmpegStatus{}).Where("camera_name=?", name).Find(&pid).Error
//		if pid.ID == 0 {
//			//todo 不存在，直接推流
//			pids = append(pids, name)
//		}
//
//		if err != nil && err != gorm.ErrRecordNotFound {
//			log.Println(err)
//			return
//		}
//	}
//	////todo 接下来存数据库，修改状态，推流的为1，没有推流的为0，不在ffmpeg中的name全部停止并改status为0  先推流再改数据库
//	////先停ws,和没有选择的流，rpc传一个flag改全局变量flag
//	//ok, msg := rpc.ChangeFlag(true)
//	//if !ok {
//	//	log.Println(msg)
//	//	err = errors.New(msg)
//	//	return
//	//}
//
//	pushStreamUrl, _, err = rpc.Ffmpeg(pids, quitPid)
//	if err != nil {
//		//todo 调用释放资源接口全给他停了 ，调用释放资源
//
//		log.Println(err)
//		return
//	}
//
//	return
//
//}
//
//func QuitFfmpeg() error {
//	//得通过这个来清除文件夹中的图片
//	//先将ws停止
//
//	ok, msg := rpc.ChangeFlag(true)
//	if !ok {
//		log.Println(msg)
//		return errors.New(msg)
//	}
//	//每次调用行人检测rpc算法他会清空文件，所以这边不用管，只需要清空fast_reid文件夹下的query里的所有文件即可
//	err := os.RemoveAll("../Yolov5-Deepsort-Fastreid/fast_reid/query")
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//	//还得停异常检测
//
//	//数据库里拿进程pid停止推流
//	err, pids := dao.GetStreamPid()
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//	//停止推流
//	for _, pid := range pids {
//		IntPid, err := strconv.Atoi(pid.Pid)
//		if err != nil {
//			log.Println(err)
//			return err
//		}
//
//		process, err := os.FindProcess(IntPid)
//		if err != nil {
//			log.Println(err)
//			return err
//		}
//		err = process.Kill()
//		if err != nil {
//			log.Println(err)
//			return err
//		}
//
//	}
//	//数据库里的数据改
//	err = dao.DB.Model(model.FfmpegStatus{}).Where("status=?", 1).UpdateColumn("status", 0).Error
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//	return nil
//
//}

func FfmpegStream(ffmpeg model.FfmpegStreamStruct) (pushStreamUrl []string, err error) {
	//先全部关闭，修改flag,和数据库的status
	//ok, msg := rpc.ChangeFlag(true)
	//if !ok {
	//	log.Println(msg)
	//	err = errors.New(msg)
	//	return
	//}
	//
	//ok2, msg2 := rpc.ChangeFlag(false)
	//if !ok2 {
	//	log.Println(msg2)
	//	err = errors.New(msg2)
	//	return
	//}
	log.Println("修改flag成功")
	//修改数据库的status
	tx := dao.DB.Begin()
	//数据库里的数据改
	for _, name := range ffmpeg.CameraName {
		err = tx.Model(model.DeviceIp{}).Where("is_display=? and device_name=?", 0, name).UpdateColumn("is_display", 1).Error
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return
		}
	}

	//重新推流
	pushStreamUrl, _, err = rpc.Ffmpeg(tx, ffmpeg.CameraName)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return
	}

	return

}

func QuitFfmpeg() error {
	//得通过这个来清除文件夹中的图片
	//先将ws停止
	ok, msg := rpc.ChangeFlag(true)
	if !ok {
		log.Println(msg)
		return errors.New(msg)
	}

	//查库，得到所有的pid，停止推流
	err, pids := dao.GetStreamPid()
	if err != nil {
		log.Println(err)
		return err

	}
	//停止推流
	for _, pid := range pids {
		process, err := os.FindProcess(pid.PId)
		if err != nil && err != os.ErrProcessDone {
			log.Println(err)
			return err
		}
		if err != os.ErrProcessDone {
			err = process.Signal(os.Interrupt) // 发送中断信号（SIGINT）给子进程
			if err != nil {
				log.Println(err)
				return err
			}
		}
		//数据库删除
		err = dao.DB.Unscoped().Where("p_id=?", pid.PId).Delete(model.QuitFfmpeg{}).Error
		if err != nil {
			log.Println(err)
			return err
		}

	}

	//数据库里的数据改
	err = dao.DB.Model(model.DeviceIp{}).Where("is_display=?", 1).UpdateColumn("is_display", 0).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil

}

// ReleaseResources 释放资源接口
func ReleaseResources() error {
	//得通过这个来清除文件夹中的图片
	//先将ws停止

	ok, msg := rpc.ChangeFlag(true)
	if !ok {
		log.Println(msg)
		return errors.New(msg)
	}

	//查库，得到所有的pid，停止推流
	err, pids := dao.GetStreamPid()
	if err != nil {
		log.Println(err)
		return err

	}

	//停止推流
	for _, pid := range pids {

		process, err := os.FindProcess(pid.PId)
		if err != nil {
			log.Println(err)
			return err
		}
		err = process.Signal(os.Interrupt) // 发送中断信号（SIGINT）给子进程
		if err != nil {
			log.Println(err)
			return err
		}

		//数据库删除
		err = dao.DB.Unscoped().Where("p_id=?", pid.PId).Delete(model.QuitFfmpeg{}).Error
		if err != nil {
			log.Println(err)
			return err
		}

	}

	//每次调用行人检测rpc算法他会清空文件，所以这边不用管，只需要清空fast_reid文件夹下的query里的所有文件即可
	err = os.RemoveAll("../../Yolov5DeepsortFastreid/fast_reid/query")
	if err != nil {
		log.Println(err)
		return err
	}

	//数据库里的数据改
	err = dao.DB.Model(model.DeviceIp{}).Where("is_display=?", 1).UpdateColumn("is_display", 0).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
