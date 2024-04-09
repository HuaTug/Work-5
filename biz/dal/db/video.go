package db

import (
	"Hertz_refactored/biz/model/video"
	"sync"

	"github.com/sirupsen/logrus"
)

var wg sync.WaitGroup

func Feedlist(lateTime string) ([]*video.Video, error) {
	var videos []*video.Video
	if err := Db.Model(&video.Video{}).Where("publish_time<?", lateTime).Find(&videos).Error; err != nil {
		logrus.Info(err)
		return videos, err
	}
	return videos, nil
}

func Videolist(req video.VideoFeedListRequest) ([]*video.Video, int64, error) {
	var videos []*video.Video
	var count int64
	if err := Db.Model(&video.Video{}).Where("author_id=?", req.AuthorId).Count(&count).Limit(int(req.PageSize)).
		Offset(int((req.PageNum - 1) * req.PageSize)).Find(&videos).Error; err != nil {
		logrus.Info(err)
		return videos, count, err
	}
	return videos, count, nil
}

func Videosearch(req video.VideoSearchRequest) ([]*video.Video, int64, error) {
	var video2 []*video.Video
	var count int64
	var err error
	if req.Keyword != "" {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err = Db.Model(&video.Video{}).
				Where("title like ? And publish_time<? &&publish_time>?", "%"+req.Keyword+"%", req.ToDate, req.FromDate).
				Count(&count).
				Limit(int(req.PageSize)).Offset(int((req.PageNum - 1) * req.PageSize)).
				Find(&video2).Error; err != nil {
				logrus.Info(err)

			}
		}()
		if err != nil {
			return video2, count, err
		}
		wg.Wait()
	}
	return video2, count, nil
}

func VideoCreate(videos video.Video) error {
	return Db.Model(&video.Video{}).Create(videos).Error
}

func FindVideo(videoId int64) (videos *video.Video, err error) {
	if err = Db.Model(&video.Video{}).Where("video_id=?", videoId).Find(&videos).Error; err != nil {
		logrus.Error(err)
		return videos, err
	}
	return videos, err

}
