package utils

import (
	"github.com/sirupsen/logrus"
)

func Transfer(v interface{}) int64 {
	var userid int64
	if v, ok := v.(float64); ok {
		userid = int64(v)
	} else {
		logrus.Info("类型断言失败,无法将变量转换为float64类型")
	}
	return userid
}
