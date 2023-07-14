/**
 * @Author: lrc
 * @Date: 2023/6/26-11:46
 * @Desc:
 **/

package model

import (
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Time       string `json:"time"`
	DeviceName string `json:"device_name"`
	PersonName string `json:"person_name"`
	Location   string `json:"location"`
}

type AbImgRecord struct {
	gorm.Model
	DeviceName string `json:"device_name"`
	ImgPath    string `json:"img_path"`
	Status     int    `json:"status"`
	VadScore   int    `json:"vad_score"`
}

type AbnormalRcReturn struct {
	DeviceName string `json:"device_name"`
	Time       string `json:"time"`
}

type QuitFfmpeg struct {
	gorm.Model
	RtmpUrl string `json:"rtmp_url"`
	PId     int    `json:"p_id"`
}
