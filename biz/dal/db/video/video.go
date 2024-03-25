package video

import (
	"Hertz_refactored/biz/dal/mysql"
	"Hertz_refactored/biz/dal/redis"
	"Hertz_refactored/biz/model/video"
	"Hertz_refactored/biz/pkg/logging"
	"github.com/sirupsen/logrus"
)

func FeedList(req video.FeedServiceRequest) ([]*video.Video, error) {
	lateTime := req.LastTime
	var videos []*video.Video
	if err := mysql.Db.Model(&video.Video{}).Where("publish_time<?", lateTime).Find(&videos).Error; err != nil {
		logrus.Info(err)
		return nil, err
	}
	//用这个redis自带的函数进行排序 完成一个排行榜的操作
	for _, v := range videos {
		err := redis.RangeAdd(v.FavoriteCount, v.VideoId)
		if err != nil {
			logging.Error(err)
		}
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
