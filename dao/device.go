/**
 * @Author: lrc
 * @Date: 2023/6/1-16:48
 * @Desc:
 **/

package dao

import "videoStream/model"

func AddDevice(name string, location string, ip string, index string) error {
	return DB.Model(&model.DeviceIp{}).Create(&model.DeviceIp{
		DeviceName:     name,
		DeviceLocation: location,
		Ip:             ip,
		Index:          index,
	}).Error
}

func SearchDevice() (device []model.DeviceIp, err error) {
	err = DB.Model(&model.DeviceIp{}).Find(&device).Error
	return
}

func DeleteDevice(name string) error {
	return DB.Model(&model.DeviceIp{}).Where("device_name = ?", name).Delete(&model.DeviceIp{}).Error
}

func ChangeDevice(name string, location string, ip string, index string) error {
	return DB.Model(&model.DeviceIp{}).Where("ip = ?", ip).UpdateColumn("device_location", location).UpdateColumn("device_name", name).UpdateColumn("index", index).Error
}
