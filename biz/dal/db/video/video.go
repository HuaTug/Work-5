package video

import (
	"Hertz_refactored/biz/dal/mysql"
	"Hertz_refactored/biz/dal/redis"
	"Hertz_refactored/biz/model/video"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"strconv"
)

func FeedList(req video.FeedServiceRequest) ([]*video.Video, error) {
	lateTime := req.LastTime
	var videos []*video.Video
	if err := mysql.Db.Model(&video.Video{}).Where("publish_time<?", lateTime).Find(&videos).Error; err != nil {
		logrus.Info(err)
		return nil, err
	}
	//ToDo:这个业务逻辑如何处理
	/*	for i, v := range videos {
			author, err := GetUserInfo(v.AuthorId)
			CacheSetAuthor(v.Id, v.AuthorId)
			if err != nil {
				return videos, err
			}
		}
		author, err := GetUserInfo(2)
		if err != nil {
			logrus.Info(err)
		}
		fmt.Println(author)*/
	return videos, nil
}

func VideoList(req video.VideoFeedListRequest) ([]*video.Video, int64, error) {
	var videos []*video.Video
	var count int64
	if err := mysql.Db.Model(&video.Video{}).Where("author_id=?", req.AuthorId).Count(&count).Limit(int(req.PageSize)).
		Offset(int((req.PageNum - 1) * req.PageSize)).Find(&videos).Error; err != nil {
		logrus.Info(err)
		return nil, 0, err
	}
	return videos, count, nil

}

func VideoSearch(req video.VideoSearchRequest) ([]*video.Video, int64, error) {
	var video2 []*video.Video
	var count int64
	if req.Keyword != "" {
		if err := mysql.Db.Model(&video.Video{}).Where("title like ?", "%"+req.Keyword+"%").Or("publish_time<? &&publish_time>?", req.ToDate, req.FormDate).Count(&count).
			Limit(int(req.PageSize)).Offset(int((req.PageNum - 1) * req.PageSize)).Find(&video2).Error; err != nil {
			logrus.Info(err)
			return video2, count, err
		}
	}

	return video2, count, nil
}

func CacheSetAuthor(videoid, authorid int64) {
	key := strconv.FormatInt(videoid, 10)
	err := redis.CacheHSet("video", key, authorid)
	if err != nil {
		logrus.Errorf("set cache error:%+v", err)
	}
}

func CacheGetAuthor(videoid int64) (int64, error) {
	key := strconv.FormatInt(videoid, 10)
	data, err := redis.CacheHGet("video", key)
	if err != nil {
		return 0, err
	}
	uid := int64(0)
	err = json.Unmarshal(data, &uid)
	if err != nil {
		return 0, err
	}
	return uid, nil
}
