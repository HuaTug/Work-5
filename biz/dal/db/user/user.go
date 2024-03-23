package user

import (
	"Hertz_refactored/biz/dal/mysql"
	"Hertz_refactored/biz/dal/redis"
	"Hertz_refactored/biz/model/user"
	"Hertz_refactored/biz/pkg/errno"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"strconv"
)

func CreateUser(users *user.User) error {
	return mysql.Db.Create(users).Error
}

func DeleteUser(userId int64) error {
	return mysql.Db.Where("user_id = ?", userId).Delete(&user.User{}).Error
}

func UpdateUser(user *user.User, uid int64) error {
	return mysql.Db.Where("user_id=?", uid).Updates(user).Error
}

func QueryUser(keyword *string, page, pageSize int64) ([]*user.User, int64, error) {
	db := mysql.Db.Model(user.User{})
	if keyword != nil && len(*keyword) != 0 {
		db = db.Where("user_name like ?", "%"+*keyword+"%")
	}
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var res []*user.User
	if err := db.Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Find(&res).Error; err != nil {
		return nil, 0, err
	}
	return res, total, nil
}

func GetUser(username string) (user.User, error) {
	var users user.User
	if err := mysql.Db.Model(user.User{}).Where("user_name=?", username).Find(&users).Error; err != nil {
		logrus.Info(err)
		return users, err
	}
	return users, nil
}

func LoginUser(req user.LoginUserResquest) (err error) {
	if err := mysql.Db.Model(user.User{}).Where("user_name=? AND password=?", req.Username, req.Password).Error; err != nil {
		return err
	}
	return nil
}

func VerifyUser(username, password string) (int64, error) {
	var user user.User
	if err := mysql.Db.Where("user_name=? AND password=?", username, password).Find(&user).Error; err != nil {
		panic(err)
	}
	if user.UserID == 0 {
		err := errno.PasswordIsNotVerified
		return user.UserID, err
	}
	return user.UserID, nil
}

func CheckUserExistById(user_id int64) (bool, error) {
	var users user.User
	if err := mysql.Db.Where("id=?", user_id).Find(&users).Error; err != nil {
		return false, err
	}
	if users == (user.User{}) {
		return false, nil
	}
	return true, nil
}

// ToDo:这是对JWT 登录认证时候的检验 通过这种切片的方式完成
func CheckUser(account, password string) ([]*user.User, error) {
	res := make([]*user.User, 0)
	if err := mysql.Db.Model(&user.User{}).Where("user_name =? AND password =?", account, password).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func CacheIdAndName(u int64, username string) {
	uid := strconv.FormatInt(u, 10)
	err := redis.CacheHSet("Map:"+uid, uid, username)
	if err != nil {
		logrus.Info(err)
		return
	}
}

func CacheGetIdAndName(u int64) string {
	uid := strconv.FormatInt(u, 10)
	v, err := redis.CacheHGet2("Map:"+uid, uid)
	if err != nil {
		logrus.Info(err)
	}
	return v
}

func CacheSetUser(u user.User) {
	key := u.UserName
	err := redis.CacheSet("user:"+key, u)
	if err != nil {
		logrus.Info("Set cache error: ", err)
	}
}

func CacheGetUser(username string) (user.User, error) {
	key := username
	data, err := redis.CacheGet("user:" + key)
	user := user.User{}
	if err != nil {
		logrus.Info(err)
		return user, err
	}
	_ = json.Unmarshal(data, &user)
	if err != nil {
		logrus.Info(err)
		return user, err
	}
	return user, nil
}
