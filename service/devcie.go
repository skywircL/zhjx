/**
 * @Author: lrc
 * @Date: 2023/6/1-16:52
 * @Desc:
 **/

package service

import (
	"log"
	"videoStream/dao"
	"videoStream/model"
)

func AddDevice(name string, location string, ip string, index string) error {
	err := dao.AddDevice(name, location, ip, index)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func SearchDevice() (device []model.DeviceIp, err error) {
	device, err = dao.SearchDevice()
	if err != nil {
		log.Println(err)
	}
	return
}

func DeleteDevice(name string) error {
	err := dao.DeleteDevice(name)
	if err != nil {
		log.Println(err)
	}
	return err
}

func UpdateDevice(name string, location string, ip string, index string) error {
	err := dao.ChangeDevice(name, location, ip, index)
	if err != nil {
		log.Println(err)

	}
	return err
}
