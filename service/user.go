/**
 * @Author: lrc
 * @Date: 2023/7/8-11:21
 * @Desc:
 **/

package service

import (
	"videoStream/dao"
	"videoStream/util"
)

func JudgeUserExist(username string, password string) bool {
	pwd := dao.QueryUserPwd(username)
	verify := util.PasswordVerify(password, pwd.Password)
	return verify
}
