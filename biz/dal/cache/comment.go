package cache

import (
	"Hertz_refactored/biz/model/comment"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"strconv"
)

func CacheSetAllComment(videoId int64, c []*comment.Comment) {
	vid := strconv.FormatInt(videoId, 10)
	err := CacheHSet("comment:"+vid, vid, c)
	if err != nil {
		logrus.Info("Set Cache error: ", err)
	}
}
func CacheGetListComment(videoId int64) ([]*comment.Comment, error) {
	key := strconv.FormatInt(videoId, 10)
	data, err := CacheHGet("comment:"+key, key)
	var comments []*comment.Comment
	if err != nil {
		return comments, err
	}
	err = json.Unmarshal(data, &comments)
	return comments, nil
}
