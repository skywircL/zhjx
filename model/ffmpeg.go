/**
 * @Author: lrc
 * @Date: 2023/5/25-10:39
 * @Desc:
 **/

package model

type AbnormalRes struct {
	CameraName []string `json:"camera_name"`
	RtspUrl    []string `json:"rtsp_url"`
}
