package cache

import (
	"Hertz_refactored/biz/model/comment"
	"Hertz_refactored/biz/pkg/logging"
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

func CacheSetCommentVideo(videoId, commentId int64) error {
	key := strconv.FormatInt(commentId, 10)
	err := CacheHSet("convert:"+key, key, videoId)
	if err != nil {
		logging.Error(err)
		return err
	}
	return err
}

func CacheGetCommentVideo(commentId int64) (videoId int64, err error) {
	key := strconv.FormatInt(commentId, 10)
	data, err := CacheHGet("convert:"+key, key)
	if err != nil {
		return videoId, err
	}
	err = json.Unmarshal(data, &videoId)
	return videoId, nil
}
