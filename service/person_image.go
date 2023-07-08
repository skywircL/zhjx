/**
 * @Author: lrc
 * @Date: 2023/5/19-9:01
 * @Desc:
 **/

package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"os"
	"videoStream/dao"
	"videoStream/model"
	"videoStream/rpc"
)

func GetPersonImage() (err error, SuperRes []model.ReturnSuperResolution) {
	err, imgInfo := dao.PersonImgInfo()
	if err != nil {
		log.Println(err)
		return err, nil
	}
	for _, img := range imgInfo {
		fileByte, err := os.ReadFile(img.CoverImgPath)
		if err != nil {
			log.Println(err)
			return err, nil
		}
		var temp model.ReturnSuperResolution
		temp.ImageName = img.ImageName
		temp.CoverImg = fileByte
		temp.Path = img.Path
		SuperRes = append(SuperRes, temp)
	}
	return
}

func CreateFolder(PersonName string) error {
	err := os.MkdirAll("../PersonImageDatabase/"+PersonName, os.ModePerm)

	if err != nil {
		fmt.Println("创建文件夹失败:", err)
		return err
	}
	return nil

}

func StoreImage(c *gin.Context, files []*multipart.FileHeader, personName []string) error {

	for _, file := range files {
		err := c.SaveUploadedFile(file, "../PersonImageDatabase/"+personName[0]+"/"+file.Filename)
		if err != nil {
			log.Println(err)
			return err
		}

	}

	//选取第一张图片作为封面，并保存
	coverImg := "../PersonImageDatabase/" + personName[0] + "/" + files[0].Filename
	path := "../PersonImageDatabase/" + personName[0] + "/"

	err := dao.StoreImage(personName[0], path, coverImg)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func DeletePersonData(PersonName []string) error {
	for _, name := range PersonName {
		err := os.RemoveAll("../PersonImageDatabase/" + name)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	for _, name := range PersonName {
		err := dao.DeleteImage(name)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}

func QuitPersonTrack(PersonName string) error {
	//将query文件夹下的名称为PersonName的文件夹删除
	err := os.RemoveAll("../Yolov5DeepsortFastreid/fast_reid/query/" + PersonName)
	if err != nil {
		log.Println(err)
		return err
	}

	//todo 重新生成names.npy和另外一个npy文件
	err = rpc.PersonBank()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
