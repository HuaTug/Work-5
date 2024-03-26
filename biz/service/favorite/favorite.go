package favorite

import (
	"Hertz_refactored/biz/dal/cache"
	"Hertz_refactored/biz/dal/db"
	"Hertz_refactored/biz/model/favorite"
	"Hertz_refactored/biz/model/user"
	"Hertz_refactored/biz/service/relation"
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
)

const (
	add = int64(1)
	sub = int64(-1)
)

type FavoriteService struct {
	ctx context.Context
}

func NewFavoriteService(ctx context.Context) *FavoriteService {
	return &FavoriteService{ctx: ctx}
}

func (s *FavoriteService) Favorite(req favorite.FavoriteRequest, uid int64) error {
	favorites := &favorite.Favorite{
		VideoId: req.VideoId,
		UserId:  uid,
	}
	username := cache.CacheGetIdAndName(uid)
	//ToDo :通过查阅资料 可以通过Count计数进行判断所需要查询的数据是否存在 若不用这种方式则会导致即使查询为空 也不会报错 因为err只是查询是否成功 而err==nil只表示查询语法没有出错 不只代表数据查询为空
	//这完善了点赞机制
	if db.Judge(req.VideoId) == false {
		errs := fmt.Sprintf("用户:%s,已经对这个视频点赞过", username)
		return errors.New(errs)
	}
	if err := db.FavoriteAction(favorites); err != nil {
		return err
	}
	//ToDo:实现对视频的点赞缓存操作
	go relation.CacheChangeUserCount(uid, add, "like")
	return nil
}

func (s *FavoriteService) UnFavorite(req favorite.FavoriteRequest, userId int64) error {
	touid, err := cache.CacheGetAuthor(req.VideoId)
	if err != nil {
		logrus.Info(err)
	}
	if err := db.UnFavoriteAction(userId, req.VideoId); err != nil {
		return err
	}
	go relation.CacheChangeUserCount(touid, sub, "unlike")
	return nil
}

// ToDo :让List实现分页查询操作
func (s *FavoriteService) List(uid int64) (favorite.ListFavoriteResponse, error) {
	var favs []*favorite.Favorite
	var users []*favorite.User
	db.Db.Model(&user.User{}).Where("user_id=?", uid).Find(&users)
	db.Db.Model(&favorite.Favorite{}).Where("user_id=?", uid).Find(&favs)
	return Append(favs, users), nil
}

func Append(favs []*favorite.Favorite, users []*favorite.User) favorite.ListFavoriteResponse {
	return favorite.ListFavoriteResponse{
		Code:  200,
		Msg:   "查询列表",
		Favs:  favs,
		Users: users,
	}
}
