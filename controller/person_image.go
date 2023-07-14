/**
 * @Author: lrc
 * @Date: 2023/5/13-21:06
 * @Desc:
 **/

package controller

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"time"
	"videoStream/service"
	"videoStream/util"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1048576,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsGETImage(c *gin.Context) {
	// 升级连接
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("upgrade error:", err)
		return
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("error:", err)
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					// 发送新图片数据给前端
					data, err := readFile(event.Name)
					log.Println(event.Name)

					if err != nil {
						log.Println(err)
						continue
					}
					conn.WriteMessage(websocket.BinaryMessage, data)
				}
			case err := <-watcher.Errors:
				log.Println(1111)
				fmt.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("../../yolox/per_img")

	if err != nil {
		return
	}
	if err != nil {
		fmt.Println("error:", err)
	}

}

func readFile(path string) ([]byte, error) {
	time.Sleep(500 * time.Millisecond) //等待文件创建完成
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	//data, err := os.ReadFile(path)
	//if err != nil {
	//	return nil, err
	//}

	return bytes, nil
}

func GetImageDatabase(c *gin.Context) {
	err, info := service.GetPersonImage()
	if err != nil {
		util.HandleError(c, err)
		return
	}

	//直接返回byte[],还是来拿到参数后再请求图片
	util.OKWithData(c, info)
}

func AddPersonData(c *gin.Context) {
	form, _ := c.MultipartForm()
	personName := form.Value["personName"]
	files := form.File["images"]

	err := service.CreateFolder(personName[0])
	if err != nil {
		util.HandleError(c, err)
		return
	}
	log.Println(files)
	err = service.StoreImage(c, files, personName)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.OK(c)
}

func DeletePersonData(c *gin.Context) {
	type ImageJson struct {
		PersonName []string `json:"person_name" binding:"required"`
	}
	var img ImageJson

	err := c.ShouldBind(&img)
	if err != nil {
		util.ParamError(c)
		return
	}

	err = service.DeletePersonData(img.PersonName)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.OK(c)
}

func QuitPersonTrack(c *gin.Context) {
	type PersonJson struct {
		PersonName string `json:"person_name"`
	}
	var img PersonJson

	err := c.ShouldBind(&img)
	if err != nil {
		util.ParamError(c)
		return
	}

	err = service.QuitPersonTrack(img.PersonName)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.OK(c)

}
