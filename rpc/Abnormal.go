/**
 * @Author: lrc
 * @Date: 2023/5/24-17:41
 * @Desc:
 **/

package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
	"io"
	"log"
	"videoStream/dao"
	"videoStream/model"
)

const (
	AbnormalAddress = "127.0.0.1:50053"
)

type returnData struct {
	AnomalyScore      float32
	AnomalyCameraName string
	ImgUrl            string
}

// Abnormal 异常检测
func Abnormal(connect *websocket.Conn, CameraName []string, RtspUrl []string) error {
	conn, err := grpc.Dial(AbnormalAddress, grpc.WithBlock(), grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := NewAbnormalDetectionClient(conn)
	abnormalCli, err := c.Abnormal(context.Background(), &AbnormalRequest{
		CameraName: CameraName,
		RtspUrl:    RtspUrl,
	})
	if err != nil {
		return err
	}

	for {
		resp, err := abnormalCli.Recv()
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
				return err
			}
			err = connect.WriteMessage(websocket.TextMessage, dataJson)
			if err != nil {
				return err
			}

		}
		if err == io.EOF {
			fmt.Println(err)
			break
		}
		if err != nil {
			return err
		}
	}

	return nil
}

// 修改异常检测flag
func ChangeAbnormalFlag(flag bool) error {
	conn, err := grpc.Dial(AbnormalAddress, grpc.WithBlock(), grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := NewAbnormalDetectionClient(conn)
	_, err = c.AbnormalChangeFlag(context.Background(), &AbnormalFlagParam{
		Flag: flag,
	})
	if err != nil {
		return err
	}

	return nil
}
