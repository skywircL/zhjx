/**
 * @Author: lrc
 * @Date: 2023/7/8-11:26
 * @Desc:
 **/

package dao

import "videoStream/model"

func QueryUserPwd(username string) model.User {
	var User model.User
	DB.Model(model.User{}).Where("username = ?", username).Find(&User)
	return User
}
