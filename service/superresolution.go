/**
 * @Author: lrc
 * @Date: 2023/5/18-23:51
 * @Desc:
 **/

package service

import (
	"log"
	"mime/multipart"
	"videoStream/rpc"
)

func SuperRes(files []*multipart.FileHeader, imageName []string, IsYoloX bool)error {
	err := rpc.SuperRes(imageName, IsYoloX, files)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

