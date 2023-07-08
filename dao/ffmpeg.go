/**
 * @Author: lrc
 * @Date: 2023/5/21-11:37
 * @Desc:
 **/

package dao

import "videoStream/model"

func GetStreamPid() (err error, Pid []model.FfmpegStatus) {
	err = DB.Model(&model.FfmpegStatus{}).Where("status=1").Find(&Pid).Error
	return

}

func GetPersonTrackRecordByCondition(deviceName string, personName string, startTime string, endTime string) (Record []model.Video, err error) {
	//todo 通过时间和设备名称查询数据库，返回结果
	err = DB.Model(model.Video{}).Where(" device_name = ? and person_name=? and time between ? and ?", deviceName, personName, startTime, endTime).Find(&Record).Error
	return

}
