package cache

import (
	"Hertz_refactored/biz/model/video"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"

	"github.com/sirupsen/logrus"
	"strconv"
)

func CacheSetAuthor(videoid, authorid int64) {
	key := strconv.FormatInt(videoid, 10)
	err := CacheHSet("video:"+key, key, authorid)
	if err != nil {
		logrus.Info("Set cache error: ", err)
	}
}

func CacheGetAuthor(videoid int64) (int64, error) {
	key := strconv.FormatInt(videoid, 10)
	data, err := CacheHGet("video:"+key, key)
	if err != nil {
		return 0, nil
	}
	var uid int64
	err = json.Unmarshal(data, &uid)
	if err != nil {
		return 0, err
	}
	return uid, nil
}

func Insert(videos []*video.Video) {
	pipe := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   1,
	}).Pipeline()
	for _, v := range videos {
		err := pipe.Set(fmt.Sprintf("key:%d", v.VideoId), fmt.Sprintf("value:%d", v.FavoriteCount), 24*3600*30)
		if err != nil {
			logrus.Info(err)
		}
	}
	_, err := pipe.Exec()
	if err != nil {
		fmt.Printf("Pipeline exec error: %s", err)
	}
	logrus.Info("PipeLine 执行完毕")
}
