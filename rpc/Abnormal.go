/**
 * @Author: lrc
 * @Date: 2023/5/24-17:41
 * @Desc:
 **/

package rpc

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

const (
	AbnormalAddress = "127.0.0.1:50053"
)

type returnData struct {
	AnomalyScore      float32
	AnomalyCameraName string
	ImgUrl            string
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
