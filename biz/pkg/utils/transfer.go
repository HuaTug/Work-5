package utils

import (
	"github.com/sirupsen/logrus"
)

func Transfer(v interface{}) int64 {
	var userId int64
	if v, ok := v.(float64); ok {
		userId = int64(v)
	} else {
		logrus.Info("类型断言失败,无法将变量转换为string类型")
	}
	return userId
}
