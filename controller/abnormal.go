/**
 * @Author: lrc
 * @Date: 2023/7/7-9:38
 * @Desc:
 **/

package controller

import (
	"github.com/gin-gonic/gin"
	"videoStream/service"
	"videoStream/util"
)

func GetAbnormalRecord(c *gin.Context) {
	//调用service,推流返回
	record, err := service.GetAbnormalRecord()
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.OKWithData(c, record)

}
