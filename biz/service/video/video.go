package video

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/sirupsen/logrus"

	"Hertz_refactored/biz/dal/cache"
	"Hertz_refactored/biz/dal/db"
	"Hertz_refactored/biz/model/video"
	es "Hertz_refactored/biz/mv/Es"
	"Hertz_refactored/biz/pkg/logging"
)

type VideoService struct {
	ctx context.Context
}

func NewVideoService(ctx context.Context) *VideoService {
	return &VideoService{ctx: ctx}
}
func (s *VideoService) FeedList(req video.FeedServiceRequest) ([]*video.Video, error) {
	lateTime := req.LastTime
	if videos, err := db.Feedlist(lateTime); err != nil {
		logging.Error(err)
		return videos, err
	} else {
		//ToDo 对其进行优化操作 使用批量插入的方式 而非这种方式
		cache.Insert(videos)
		//用这个redis自带的函数进行排序 完成一个排行榜的操作
		//ToDo: 对redis进行批量插入的操作
		for _, v := range videos {
			err := cache.RangeAdd(v.FavoriteCount, v.VideoId)
			if err != nil {
				logging.Error(err)
			}
		}
		return videos, nil
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
}

func (s *VideoService) VideoList(req video.VideoFeedListRequest) ([]*video.Video, int64, error) {
	var videos []*video.Video
	var count int64
	var err error
	if videos, count, err = db.Videolist(req); err != nil {
		return videos, count, err
	}
	temp := es.VideoIndex{Index: "videos"}
	for _, v := range videos {
		err = temp.CreateVideoDoc(es.Video{
			AuthorId: v.AuthorId,
			VideoId:  v.VideoId,
			Info: es.VideoOtherData{
				PlayUrl:       v.PlayUrl,
				CoverUrl:      v.CoverUrl,
				FavoriteCount: v.FavoriteCount,
				CommentCount:  v.CommentCount,
				//PublishTime:   v.PublishTime,
				Title: v.Title,
			},
		})
		if err != nil {
			hlog.Info("Create Info Err: ", err)
		}
	}
	return videos, count, err
}

func (s *VideoService) VideoSearch(req video.VideoSearchRequest) ([]*video.Video, int64, error) {
	var video []*video.Video
	var count int64
	var err error
	//NewVideoService(context.Background()).SpecialSearch()
	NewVideoService(context.Background()).KeywordSearch(req.Keyword)
	if video, count, err = db.Videosearch(req); err != nil {
		logging.Error(err)
		return video, count, err
	}
	return video, count, err
}

// 使用es进行测试单条数据查询  (根据唯一标识符查找)
func (s *VideoService) SpecialSearch() *es.Video {
	var video *es.Video
	var err error
	temp := es.VideoIndex{Index: "videos"}
	video, err = temp.SearchVideoDoc(2)
	if err != nil {
		logrus.Info(err)
		return nil
	}
	hlog.Info(video.Info.Title)
	return video

}

func (s *VideoService) KeywordSearch(keywords string) []*es.Video {
	var videos []*es.Video
	var err error
	var cnt int64
	temp := es.VideoIndex{Index: "videos"}
	videos, cnt, err = temp.SearchVideoDocDefault(keywords)
	if err != nil {
		hlog.Info(err)
	}
	if len(videos) >= 1 {
		for i, _ := range videos {
			hlog.Info(videos[i].Info.Title)
		}
	}
	hlog.Info(cnt)
	return videos
}
func (s *VideoService) VideoPopular() (videos []*video.Video, err error) {
	//resp := new(video.VideoPopularResponse)
	//ToDo :排行的显示功能有问题 可以在redis内直接查看
	res, err := cache.RangeList("Rank")
	if err != nil {
		logging.Error(err)
		return
	}
	var temp *video.Video
	for i := 0; i < len(res); i++ {
		v, err := strconv.Atoi(res[i])
		if err != nil {
			logrus.Info(err)
			return videos, err
		}
		temp, _ = db.FindVideo(int64(v))
		videos = append(videos, temp)

	}
	return videos, nil
}
