package db

import (
	"Hertz_refactored/biz/model/favorite"
	"Hertz_refactored/biz/model/video"

	"github.com/sirupsen/logrus"
)

func FavoriteAction(favorites *favorite.Favorite) error {
	if err := Db.Model(&favorite.Favorite{}).Create(favorites).Error; err != nil {
		logrus.Info(err)
		return err
	}
	return nil
}

func UnFavoriteAction(userId, videoId int64) error {
	if err := Db.Where("user_id=? And video_id=?", userId, videoId).Delete(&favorite.Favorite{}).Error; err != nil {
		logrus.Info(err)
		return err
	}
	return nil
}

func Judge(VideoId int64) bool {
	var count int64
	//ToDo :通过查阅资料 可以通过Count计数进行判断所需要查询的数据是否存在 若不用这种方式则会导致即使查询为空 也不会报错 因为err只是查询是否成功 而err==nil只表示查询语法没有出错 不只代表数据查询为空
	//这完善了点赞机制
	//ToDo question:如何让不同的人对视频进行点赞操作

	if Db.Model(&video.Video{}).Where("video_id=? And favorite_count!=?", VideoId, 0).Count(&count); count > 0 {
		return false
	}
	return true
}
