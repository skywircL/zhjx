/**
 * @Author: lrc
 * @Date: 2023/5/14-14:36
 * @Desc:
 **/

package util

import "github.com/satori/go.uuid"

func GetGID() string {
	return uuid.NewV4().String()
}
