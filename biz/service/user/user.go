package user

import (
	"Hertz_refactored/biz/dal/cache"
	"Hertz_refactored/biz/dal/db"
	"Hertz_refactored/biz/model/user"
	"Hertz_refactored/biz/pkg/logging"
	"Hertz_refactored/biz/pkg/util"
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	ctx context.Context
}

// ToDo:通过调用这个函数，满足handler层可以通过这个接口调用service内的业务逻辑
func NewUserService(ctx context.Context) *UserService {
	return &UserService{ctx: ctx}
}

func (s *UserService) CreateUser(req user.CreateUserRequest) (users *user.User, err error) {
	password, err := util.Crypt(req.Password)
	User := &user.User{
		UserName: req.Name,
		Password: password,
	}
	return db.CreateUser(s.ctx, User)

}

func (s *UserService) LoginUser(req user.LoginUserResquest) (err error) {
	User := &user.User{
		UserName: req.Username,
		Password: req.Password,
	}
	if err := db.VerifyUser(s.ctx, User); err != nil {
		return err
	}
	return nil
}

// ToDo:这是对JWT 登录认证时候的检验 通过这种切片的方式完成
func CheckUser(account, password string) (user.User, error) {
	var users user.User
	db.Db.Model(&user.User{}).Where("user_name =?", account).Find(&users)
	if flag := util.VerifyPassword(password, users.Password); flag == false {
		return users, errors.New("密码错误")
	}
	return users, nil

}

func (s *UserService) GetInfo(userid int64) (User *user.User, err error) {
	users, err := cache.CacheGetUser(userid)
	if err != nil {
		User, err = db.GetUser(userid)
		fmt.Println(User.UserID)
		if err != nil {
			logrus.Info("该用户不存在")
			return nil, err
		}
		go cache.CacheSetUser(User)
		return User, nil
	}
	return users, err
}

func (s *UserService) Update(req user.UpdateUserRequest, userId int64) error {
	password, _ := util.Crypt(req.Password)
	User := &user.User{
		UserName: req.Name,
		Password: password,
	}
	if err := db.UpdateUser(User, userId); err != nil {
		logging.Info(err)
	}
	return nil
}

func (s *UserService) Delete(userId int64) error {
	if err := db.DeleteUser(userId); err != nil {
		logging.Error(err)
		return err
	}
	return nil
}

func (s *UserService) Query(req *user.QueryUserRequest) ([]*user.User, int64, error) {
	var User []*user.User
	var err error
	var total int64
	if User, total, err = db.QueryUser(req.Keyword, req.Page, req.PageSize); err != nil {
		logging.Error(err)
		return User, total, err
	}
	return User, total, err
}
