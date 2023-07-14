/**
 * @Author: lrc
 * @Date: 2023/5/18-16:53
 * @Desc:
 **/

package model

import "gorm.io/gorm"

type SuperResolution struct {
	gorm.Model
	Path         string `json:"path"`
	ImageName    string `json:"image_name"`
	CoverImgPath string `json:"cover_img_path"`
}

type ReturnSuperResolution struct {
	gorm.Model
	Path      string `json:"path"`
	ImageName string `json:"image_name"`
}
