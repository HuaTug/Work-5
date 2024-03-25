package favorite

import (
	"Hertz_refactored/biz/dal/db/relation"
	user2 "Hertz_refactored/biz/dal/db/user"
	"Hertz_refactored/biz/dal/mysql"
	"Hertz_refactored/biz/handler/video"
	"Hertz_refactored/biz/model/favorite"
	"Hertz_refactored/biz/model/user"
	video2 "Hertz_refactored/biz/model/video"
	"github.com/sirupsen/logrus"
	"log"
)

const (
	add = int64(1)
	sub = int64(-1)
)

func Favorite(req favorite.FavoriteRequest, uid int64) error {

	favorites2 := &favorite.Favorite{
		VideoId: req.VideoId,
		UserId:  uid,
	}
	username := user2.CacheGetIdAndName(uid)
	var count int64
	//ToDo :通过查阅资料 可以通过Count计数进行判断所需要查询的数据是否存在 若不用这种方式则会导致即使查询为空 也不会报错 因为err只是查询是否成功 而err==nil只表示查询语法没有出错 不只代表数据查询为空
	//这完善了点赞机制
	if mysql.Db.Model(&video2.Video{}).Where("video_id=? And favorite_count!=?", req.VideoId, 0).Count(&count); count > 0 {
		log.Printf("用户:%s,已经对这个视频点赞过", username)
		return nil
	}
	if err := mysql.Db.Model(&favorite.Favorite{}).Create(favorites2).Error; err != nil {
		logrus.Info(err)
		return err
	}
	//ToDo:实现对视频的点赞缓存操作
	go relation.CacheChangeUserCount(uid, add, "like")
	return nil
}

func UnFavorite(req favorite.FavoriteRequest, uid interface{}) error {
	var userid int64
	if v, ok := uid.(float64); ok {
		userid = int64(v)
	} else {
		logrus.Info("类型断言失败，无法将变量转换为float64类型")
	}
	touid, err := video.CacheGetAuthor(req.VideoId)
	if err != nil {
		logrus.Info(err)
	}
	if err := mysql.Db.Where("user_id=? And video_id=?", userid, req.VideoId).Delete(&favorite.Favorite{}).Error; err != nil {
		logrus.Info(err)
		return err
	}
	go relation.CacheChangeUserCount(touid, sub, "unlike")
	return nil
}

// ToDo :让List实现分页查询操作
func List(uid interface{}) (favorite.ListFavoriteResponse, error) {
	var favs []*favorite.Favorite
	var users []*favorite.User
	mysql.Db.Model(&user.User{}).Where("user_id=?", uid).Find(&users)
	mysql.Db.Model(&favorite.Favorite{}).Where("user_id=?", uid).Find(&favs)
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
