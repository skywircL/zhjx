/**
 * @Author: lrc
 * @Date: 2023/6/25-18:57
 * @Desc:
 **/

package service

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"videoStream/dao"
	"videoStream/model"
	"videoStream/rpc"
	"videoStream/util"
)

func GetVideo(deviceName string, personName []string, startTime string, endTime string) (Rc []model.AbnormalRcReturn, err error) {
	//异常查询

	for _, v := range personName {
		record, err := dao.GetPersonTrackRecordByCondition(deviceName, v, startTime, endTime)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		for _, re := range record {
			Rc = append(Rc, model.AbnormalRcReturn{
				DeviceName: deviceName,
				Time:       re.Time,
			})
		}
	}

	//清空query文件夹
	err = deleteFiles("../Yolov5-Deepsort-Fastreid/fast_reid/query")
	if err != nil {
		log.Println(err)
		return
	}

	//还得创建文件夹
	err = os.MkdirAll("../Yolov5-Deepsort-Fastreid/fast_reid/query", os.ModePerm)
	if err != nil {
		return
	}

	//移动图片到query文件夹下，查数据库
	for _, name := range personName {
		//todo 通过图片库名称来查询图片的存储路径
		var imagePath model.SuperResolution //这个指向的是一个文件夹，把文件夹下的照片复制到input文件夹下
		dao.DB.Model(model.SuperResolution{}).Where("image_name = ?", name).Find(&imagePath)
		//通过imagePath将该路径下的所有文件移到input文件夹下
		err = copyFiles(imagePath.Path, "../demosr/inputs") //todo 得改参数
		if err != nil {
			return
		}

	}

	//拿图片重新生成，调用person bank
	err = rpc.PersonBank()
	if err != nil {
		log.Println(err)
		return
	}
	//根据传过来的时间和设备选择对应视频文件调用rpc，异步，还得同步调用异常检查，这个可以前端同时调用，合成成一个视频

	return
}

func VideoStream(StartTime string, EndTime string, DeviceName string) (rtmpUrl string, err error) {
	//根据时间和设备生成视频文件存储路径

	var path []string

	//传递path推流
	conn, err := grpc.Dial("127.0.0.1:50054", grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return
	}
	gid := util.GetGID()
	rtmpUrl = "rtmp://127.0.0.1:1935/live/reid_streamvideo" + gid

	defer conn.Close()
	c := rpc.NewBacktrackingClient(conn)
	_, err = c.Backtracking(context.Background(), &rpc.BacktrackingRequest{
		VideoPath:  path,
		CameraName: DeviceName,
		Rmtp:       rtmpUrl,
	})

	if err != nil {
		log.Println(err)
		return
	}
	// todo 得等待处理推流

	return
}

func copyFiles(sourcePath, destinationPath string) error {
	err := filepath.Walk(sourcePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil // 忽略目录
		}

		// 检查文件扩展名是否为图片格式
		ext := strings.ToLower(filepath.Ext(path))
		if ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif" {
			// 构建目标文件路径
			destinationFile := filepath.Join(destinationPath, info.Name())

			// 打开源文件
			sourceFile, err := os.Open(path)
			if err != nil {
				return err
			}
			defer sourceFile.Close()

			// 创建目标文件
			destinFile, err := os.Create(destinationFile)
			if err != nil {
				return err
			}
			defer destinFile.Close()

			// 拷贝文件内容
			_, err = io.Copy(destinFile, sourceFile)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func deleteFiles(folderPath string) error {
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			err := os.RemoveAll(path)
			if err != nil {
				return err
			}
		}
		err = os.MkdirAll("../demosr/inputs", os.ModePerm)
		if err != nil {
			return err
		}
		err = os.MkdirAll("../demosr/outputs", os.ModePerm)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func QuitVideoStream() (err error) {
	err = rpc.ChangeAbnormalFlag(true)
	return
}

func GetPersonTrackRecord() (Record []model.Video, err error) {
	//todo 通过时间和设备名称查询数据库，返回结果
	err = dao.DB.Model(model.Video{}).Find(&Record).Error
	return
}

// GetAbnormalRecord 异常数据查询接口
func GetAbnormalRecord() (Record []model.AbImgRecord, err error) {
	//todo 通过时间和设备名称查询数据库，返回结果
	Record, err = dao.GetAbnormalRecord()
	if err != nil {
		log.Println(err)
		return
	}
	return
}
