package db

import (
	"Hertz_refactored/biz/model/user"
	"Hertz_refactored/biz/pkg/errno"
	"Hertz_refactored/biz/pkg/logging"
	"context"
	"github.com/sirupsen/logrus"
)

// ToDo 完成规范形式的三层架构模式
func CreateUser(ctx context.Context, users *user.User) (*user.User, error) {
	err := Db.WithContext(ctx).Create(&users).Error
	if err != nil {
		logging.Error(err)
		return nil, err
	}
	return users, err
}

func DeleteUser(userId int64) error {
	return Db.Where("user_id = ?", userId).Delete(&user.User{}).Error
}

func UpdateUser(user *user.User, uid int64) error {
	return Db.Where("user_id=?", uid).Updates(user).Error
}

func QueryUser(keyword *string, page, pageSize int64) ([]*user.User, int64, error) {
	db := Db.Model(user.User{})
	if keyword != nil && len(*keyword) != 0 {
		db = db.Where("user_name like ?", "%"+*keyword+"%")
	}
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, total, err
	}
	var res []*user.User
	if err := db.Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Find(&res).Error; err != nil {
		return nil, total, err
	}
	return res, total, nil
}

func GetUser(userid int64) (*user.User, error) {
	var users *user.User
	if err := Db.Model(user.User{}).Where("user_id=?", userid).Find(&users).Error; err != nil {
		logrus.Info(err)
		return users, err
	}
	return users, nil
}

func VerifyUser(ctx context.Context, users *user.User) error {
	if err := Db.WithContext(ctx).Where("user_name=? AND password=?", users.UserName, users.Password).Error; err != nil {
		panic(err)
	}
	if users.UserID == 0 {
		err := errno.PasswordIsNotVerified
		logging.Error(err)
		return err
	}
	return nil
}

func CheckUserExistById(userId int64) (bool, error) {
	var users user.User
	if err := Db.Where("id=?", userId).Find(&users).Error; err != nil {
		return false, err
	}
	if users == (user.User{}) {
		return false, nil
	}
	return true, nil
}
