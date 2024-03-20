package favorite

import (
	"Hertz_refactored/biz/dal/mysql"
	"Hertz_refactored/biz/model/favorite"
	"Hertz_refactored/biz/model/user"
	"github.com/sirupsen/logrus"
)

func Favorite(req favorite.FavoriteRequest, uid int64) error {

	favorites2 := &favorite.Favorite{
		VideoId: req.VideoId,
		UserId:  uid,
	}
	if err := mysql.Db.Model(&favorite.Favorite{}).Create(favorites2).Error; err != nil {
		logrus.Info(err)
		return err
	}
	return nil
}

func UnFavorite(req favorite.FavoriteRequest, uid interface{}) error {
	var userid int64
	if v, ok := uid.(float64); ok {
		userid = int64(v)
	} else {
		logrus.Info("类型断言失败，无法将变量转换为float64类型")
	}
	if err := mysql.Db.Where("user_id=? And video_id=?", userid, req.VideoId).Delete(&favorite.Favorite{}).Error; err != nil {
		logrus.Info(err)
		return err
	}
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