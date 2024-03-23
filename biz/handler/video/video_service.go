// Code generated by hertz generator.

package video

import (
	video2 "Hertz_refactored/biz/dal/db/video"
	"Hertz_refactored/biz/dal/redis"
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"strconv"

	video "Hertz_refactored/biz/model/video"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// FeedService .
// @router /v1/video/ [GET]
func FeedService(ctx context.Context, c *app.RequestContext) {
	var err error
	var req video.FeedServiceRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	var videos []*video.Video
	if videos, err = video2.FeedList(req); err != nil {
		logrus.Info(err)
	}
	for _, v := range videos {
		CacheSetAuthor(v.VideoId, v.AuthorId)
	}
	c.JSON(consts.StatusOK, video.FeedServiceResponse{
		Code:      consts.StatusOK,
		Msg:       "展示视频流列表",
		VideoList: videos,
	})

}

// VideoFeedList .
// @router /v1/video/list [GET]
func VideoFeedList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req video.VideoFeedListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	var videos []*video.Video
	var count int64
	if videos, count, err = video2.VideoList(req); err != nil {
		logrus.Info(err)
	}
	c.JSON(consts.StatusOK, video.VideoFeedListResponse{
		Code:      consts.StatusOK,
		Msg:       "展示上传视频列表",
		VideoList: videos,
		Count:     count,
	})
}

// VideoSearch .
// @router /v1/video/search [POST]
func VideoSearch(ctx context.Context, c *app.RequestContext) {
	var err error
	var req video.VideoSearchRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	var videos []*video.Video
	var count int64
	if videos, count, err = video2.VideoSearch(req); err != nil {
		logrus.Info(err)
	}

	c.JSON(consts.StatusOK, video.VideoSearchResponse{
		Code:        consts.StatusOK,
		Msg:         "视频搜索结果",
		VideoSearch: videos,
		Count:       count,
	})
}

// VideoPopular .
// @router /v1/video/popular [GET]
func VideoPopular(ctx context.Context, c *app.RequestContext) {
	var err error
	var req video.VideoPopularRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(video.VideoPopularResponse)

	c.JSON(consts.StatusOK, resp)
}

func CacheSetAuthor(videoid, authorid int64) {
	key := strconv.FormatInt(videoid, 10)
	err := redis.CacheHSet("video:"+key, key, authorid)
	if err != nil {
		logrus.Info("Set cache error: ", err)
	}
}

func CacheGetAuthor(videoid int64) (int64, error) {
	key := strconv.FormatInt(videoid, 10)
	data, err := redis.CacheHGet("video:"+key, key)
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
