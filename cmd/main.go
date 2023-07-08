/**
 * @Author: lrc
 * @Date: 2023/5/13-11:58
 * @Desc:
 **/

package main

import (
	"videoStream/dao"
	"videoStream/router"
)

func main() {
	dao.MySQLInit()
	router.RunRouter()
}
