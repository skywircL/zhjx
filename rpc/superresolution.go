/**
 * @Author: lrc
 * @Date: 2023/5/18-16:47
 * @Desc:
 **/

package rpc

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
	"videoStream/dao"
	"videoStream/model"
)

const (
	//address = "127.0.0.1:50052"
	address = "10.16.50.17:50052"
)

// SuperRes 非图片库图片得加参数处理
func SuperRes(imageName []string, IsYoloX bool, files []*multipart.FileHeader) error {
	//todo 出发错误返回时得先清空文件夹
	//先将文件夹中资源释放
	_ = deleteFiles("../../demosr/inputs")

	_ = deleteFiles("../../demosr/outputs")

	conn, err := grpc.Dial(address, grpc.WithBlock(), grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := NewSuperResolutionClient(conn)

	for _, v := range imageName {
		if len(imageName) == 1 && v == "" {
			break
		}
		//todo 通过图片库名称来查询图片的存储路径
		var imagePath model.SuperResolution //这个指向的是一个文件夹，把文件夹下的照片复制到input文件夹下
		dao.DB.Model(model.SuperResolution{}).Where("image_name = ?", v).Find(&imagePath)
		//通过imagePath将该路径下的所有文件移到input文件夹下
		err := copyFiles(imagePath.Path, "../../demosr/inputs") //todo 得改参数
		if err != nil {
			return err
		}

		//todo 将图片库中的每一组图片通过for循环调用rpc
		result, err := c.SuperResolutionFunc(context.Background(), &SuperResolutionRequest{})
		if err != nil {
			return err
		}

		if result.Error { //没问题就转移
			//todo 通过返回的调用状态来判断是否应该将图从output转移到query文件夹下
			//还得创建文件夹
			err := os.MkdirAll("../../fast_reid/query/"+v, os.ModePerm)
			if err != nil {
				return err
			}

			err = copyFiles(imagePath.Path, "../../fast_reid/query/"+v)
			if err != nil {
				return err
			}
		}
	}

	//todo 根据参数如果需要非图片库图片处理
	if IsYoloX {
		for _, file := range files {
			src, err := file.Open()
			if err != nil {
				return err
			}
			defer src.Close()

			// 构建目标文件路径
			destinationPath := "../../demosr/inputs/" + file.Filename
			NewFile, err := os.Create(destinationPath)
			if err != nil {
				return err
			}
			defer NewFile.Close()
			content, err := ioutil.ReadAll(src)
			if err != nil {
				return err
			}
			_, err = NewFile.Write(content)
			if err != nil {
				return err
			}

			result, err := c.SuperResolutionFunc(context.Background(), &SuperResolutionRequest{})
			if err != nil {
				return err
			}
			if result.Error { //没问题就转移
				// todo 通过返回的调用状态来判断是否应该将图从output转移到query文件夹下
				//先创建文件夹
				err = os.MkdirAll("../../Yolov5DeepsortFastreid/fast_reid/query/temp", os.ModePerm)
				if err != nil {
					return err
				}

				//等待处理
				time.Sleep(5 * time.Second)

				err := copyFiles("../../demosr/outputs/restored_imgs", "../../Yolov5DeepsortFastreid/fast_reid/query/temp")
				if err != nil {
					return err
				}
			}

		}

	}
	// rpc执行person_bank  这边先不清理，默认行人检测获取的图片只检测一个人，后面有需求在改
	bank, err := c.PersonBank(context.Background(), &PersonBankRequest{})
	if err != nil {
		return err
	}
	if !bank.Error {
		return errors.New(bank.Message)
	}
	return nil
}

// PersonBank rpc调用personBank
func PersonBank() error {
	conn, err := grpc.Dial(address, grpc.WithBlock(), grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := NewSuperResolutionClient(conn)

	bank, err := c.PersonBank(context.Background(), &PersonBankRequest{})
	if err != nil {
		return err
	}
	if !bank.Error {
		return errors.New(bank.Message)
	}
	//生成成功后清空query文件夹

	return nil
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
		err = os.MkdirAll("../../demosr/inputs", os.ModePerm)
		if err != nil {
			return err
		}
		err = os.MkdirAll("../../demosr/outputs", os.ModePerm)
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
