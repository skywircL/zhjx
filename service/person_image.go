/**
 * @Author: lrc
 * @Date: 2023/5/19-9:01
 * @Desc:
 **/

package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"strings"
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
		split := strings.Split(img.CoverImgPath, "/")
		imgName := split[len(split)-1]
		personName := split[len(split)-2]
		path := "10.16.50.17:8080" + "/img" + "/" + personName + "/" + imgName

		var temp model.ReturnSuperResolution
		temp.ImageName = img.ImageName
		temp.Path = path
		temp.ID = img.ID
		SuperRes = append(SuperRes, temp)
	}
	return
}

func CreateFolder(PersonName string) error {
	err := os.MkdirAll("../../PersonImageDatabase/"+PersonName, os.ModePerm)

	if err != nil {
		fmt.Println("创建文件夹失败:", err)
		return err
	}
	return nil

}

func StoreImage(c *gin.Context, files []*multipart.FileHeader, personName []string) error {
	//先查询数据库中是否有该人物的信息，如果有则返回错误
	for _, file := range files {
		err, m := dao.QueryOnlyImgInfo(file.Filename)
		if err != nil {
			log.Println(err)
			return err
		}
		if m.ID == 0 {
			return errors.New(file.Filename + "该人物名称已存在")
		}

		err = c.SaveUploadedFile(file, "../../PersonImageDatabase/"+personName[0]+"/"+file.Filename)
		if err != nil {
			log.Println(err)
			return err
		}

		//选取第一张图片作为封面，并保存
		coverImg := "../../PersonImageDatabase/" + personName[0] + "/" + files[0].Filename
		path := "../../PersonImageDatabase/" + personName[0] + "/"

		err = dao.StoreImage(personName[0], path, coverImg)
		if err != nil {
			log.Println(err)
			return err
		}

	}

	return nil
}

func DeletePersonData(PersonName []string) error {
	for _, name := range PersonName {
		err := os.RemoveAll("../../PersonImageDatabase/" + name)
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

func CreatePersonBank(files []*multipart.FileHeader, imageName []string, IsYoloX bool) error {
	//先将文件夹中资源释放
	_ = DeleteFiles("../../Yolov5DeepsortFastreid/fast_reid/query")

	_ = os.MkdirAll("../../Yolov5DeepsortFastreid/fast_reid/query", os.ModePerm)

	for _, v := range imageName {
		if len(imageName) == 1 && v == "" {
			break
		}
		//todo 通过图片库名称来查询图片的存储路径
		var imagePath model.SuperResolution //这个指向的是一个文件夹，把文件夹下的照片复制到input文件夹下
		dao.DB.Model(model.SuperResolution{}).Where("image_name = ?", v).Find(&imagePath)
		err := os.MkdirAll("../../Yolov5DeepsortFastreid/fast_reid/query/"+v, os.ModePerm)
		if err != nil {
			return err
		}
		err = copyFiles(imagePath.Path, "../../Yolov5DeepsortFastreid/fast_reid/query/"+v)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	//todo 根据参数如果需要非图片库图片处理
	if IsYoloX {
		for _, file := range files {
			src, err := file.Open()
			if err != nil {
				log.Println(err)
				return err
			}
			defer src.Close()

			// 构建目标文件路径
			_ = os.MkdirAll("../../Yolov5DeepsortFastreid/fast_reid/query/temp", os.ModePerm)
			destinationPath := "../../Yolov5DeepsortFastreid/fast_reid/query/temp/" + file.Filename
			NewFile, err := os.Create(destinationPath)
			if err != nil {
				log.Println(err)
				return err
			}
			defer NewFile.Close()
			content, err := ioutil.ReadAll(src)
			if err != nil {
				log.Println(err)
				return err
			}
			_, err = NewFile.Write(content)
			if err != nil {
				log.Println(err)
				return err
			}
		}

		err := rpc.PersonBank()
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}
