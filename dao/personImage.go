/**
 * @Author: lrc
 * @Date: 2023/5/19-9:05
 * @Desc:
 **/

package dao

import (
	"videoStream/model"
)

func PersonImgInfo() (err error, findImg []model.SuperResolution) {
	err = DB.Model(&model.SuperResolution{}).Find(&findImg).Error
	return
}

func StoreImage(personName string, Path string, coverImg string) error {
	err := DB.Create(&model.SuperResolution{
		Path:         Path,
		ImageName:    personName,
		CoverImgPath: coverImg,
	}).Error
	return err
}

func DeleteImage(PersonName string) error {
	err := DB.Where("image_name=?", PersonName).Delete(&model.SuperResolution{}).Error
	return err
}

func QueryOnlyImgInfo(personName string) (err error, findImg model.SuperResolution) {
	err = DB.Model(&model.SuperResolution{}).Find(&findImg).Where("image_name=?", personName).Error
	return
}
