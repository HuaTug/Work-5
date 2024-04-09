package favorite

import (
	"context"
	"sync"

	"github.com/sirupsen/logrus"

	"Hertz_refactored/biz/dal/cache"
	"Hertz_refactored/biz/dal/db"
	"Hertz_refactored/biz/model/favorite"
	"Hertz_refactored/biz/model/user"
	"Hertz_refactored/biz/pkg/logging"
	"Hertz_refactored/biz/service/relation"
)

const (
	add = int64(1)
	sub = int64(-1)
)

var wg sync.WaitGroup

type FavoriteService struct {
	ctx context.Context
}

func NewFavoriteService(ctx context.Context) *FavoriteService {
	return &FavoriteService{ctx: ctx}
}

func (s *FavoriteService) Favorite(req favorite.FavoriteRequest, uid int64) error {
	//ToDo 如何实现使用redis将视频id与评论id相对应起来 例如对应着同一个视频id具有多条评论
	var err error
	videoId, err := cache.CacheGetCommentVideo(req.CommentId)
	favorites := &favorite.Favorite{
		VideoId:   videoId,
		UserId:    uid,
		CommentId: req.CommentId,
	}
	if err != nil {
		logging.Info(err)
	}
	/*	username := cache.CacheGetIdAndName(uid)
		//ToDo :通过查阅资料 可以通过Count计数进行判断所需要查询的数据是否存在 若不用这种方式则会导致即使查询为空 也不会报错 因为err只是查询是否成功 而err==nil只表示查询语法没有出错 不只代表数据查询为空
		//这完善了点赞机制
		if db.Judge(req.VideoId) == false {
			errs := fmt.Sprintf("用户:%s,已经对这个视频点赞过", username)
			return errors.New(errs)
		}*/
	wg.Add(1)
	go func() {
		err = db.FavoriteAction(favorites)
	}()
	if err != nil {
		return err
	}
	wg.Wait()

	//ToDo:实现对视频的点赞缓存操作
	go relation.CacheChangeUserCount(uid, add, "like")
	return nil
}

func (s *FavoriteService) UnFavorite(req favorite.FavoriteRequest, userId int64) error {
	var err error
	VideoId, _ := cache.CacheGetCommentVideo(req.CommentId)
	touid, err := cache.CacheGetAuthor(VideoId)
	if err != nil {
		logrus.Info(err)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = db.UnFavoriteAction(userId, VideoId)
	}()
	if err != nil {
		logrus.Info(err)
		return err
	}
	wg.Wait()
	go relation.CacheChangeUserCount(touid, sub, "unlike")
	return nil
}

// ToDo :让List实现分页查询操作
func (s *FavoriteService) List(uid int64) (favorite.ListFavoriteResponse, error) {
	var favs []*favorite.Favorite
	var users []*favorite.User
	wg.Add(2)
	go func() {
		defer wg.Done()
		db.Db.Model(&user.User{}).Where("user_id=?", uid).Find(&users)
	}()
	go func() {
		defer wg.Done()
		db.Db.Model(&favorite.Favorite{}).Where("user_id=?", uid).Find(&favs)
	}()
	wg.Wait()

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
