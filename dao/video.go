/**
 * @Author: lrc
 * @Date: 2023/7/7-12:02
 * @Desc:
 **/

package dao

import "videoStream/model"

func GetAbnormalRecord() (Record []model.AbImgRecord, err error) {
	//todo 通过时间和设备名称查询数据库，返回结果
	err = DB.Model(model.AbImgRecord{}).Find(&Record).Error
	return
}
