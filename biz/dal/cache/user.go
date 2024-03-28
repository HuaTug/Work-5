package cache //nolint:gofmt

import (
	"encoding/json"
	"strconv"

	"github.com/sirupsen/logrus"

	"Hertz_refactored/biz/model/user"
)

func CacheIdAndName(u int64, username string) {
	uid := strconv.FormatInt(u, 10)
	err := CacheHSet("Map:"+uid, uid, username)
	if err != nil {
		logrus.Info(err)
		return
	}
}

func CacheGetIdAndName(u int64) string {
	uid := strconv.FormatInt(u, 10)
	v, err := CacheHGet2("Map:"+uid, uid)
	if err != nil {
		logrus.Info(err)
	}
	return v
}

func CacheSetUser(u *user.User) {
	key := strconv.FormatInt(u.UserID, 10)
	err := CacheSet("user:"+key, u)
	if err != nil {
		logrus.Info("Set cache error: ", err)
	}
}

func CacheGetUser(id int64) (*user.User, error) {
	key := strconv.FormatInt(id, 10)
	data, err := CacheGet("user:" + key)
	users := &user.User{}
	if err != nil {
		logrus.Info(err)
		return users, err
	}
	_ = json.Unmarshal(data, &users)
	if err != nil {
		logrus.Info(err)
		return users, err
	}
	return users, nil
}
