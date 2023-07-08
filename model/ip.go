/**
 * @Author: lrc
 * @Date: 2023/5/13-10:09
 * @Desc:
 **/

package model

import "gorm.io/gorm"

type DeviceIp struct {
	gorm.Model
	Ip             string `json:"ip"`
	DeviceName     string `json:"device_name"`
	DeviceLocation string `json:"device_location"`
	IsDisplay      int    `json:"is_display"`
	Index          string `json:"index"` //同一ip下的第几个摄像头
}

type Ffmpeg struct {
	gorm.Model
	Location string `json:"location,omitempty"`
	CameraIp string `json:"camera_ip,omitempty"`
	Stream   string `json:"stream,omitempty"`
	SavePath string `json:"save_path,omitempty"`
}

type FfmpegStatus struct {
	gorm.Model
	RtspUrl    string `json:"rtsp_url"`
	CameraName string `json:"camera_name"`
	Pid        string `json:"pid"`
	Status     string `json:"status"`
}

type FfmpegStreamStruct struct {
	CameraName []string `json:"camera_name"`
}
