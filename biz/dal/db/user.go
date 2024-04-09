package db

import (
	"Hertz_refactored/biz/model/user"
	"Hertz_refactored/biz/pkg/logging"
	"Hertz_refactored/biz/pkg/util"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
)

// ToDo 完成规范形式的三层架构模式
func CreateUser(ctx context.Context, users *user.User) (*user.User, error,bool) {
	err := Db.WithContext(ctx).Create(&users).Error
	if err != nil {
		logging.Error(err)
		return nil, err,true
	}
	return users, err,true
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

func CheckUser(account, password string) (user.User, error,bool) {
	var users user.User
	var count int64
	Db.Model(&user.User{}).Where("user_name =?", account).Count(&count).Find(&users)
	if count==0{
		logrus.Info("正在创建新用户")
		return users,nil,true
	}
	if flag := util.VerifyPassword(password, users.Password); flag == false {
		return users, errors.New("密码错误"),true
	}
	return users, nil,false
}

