package user

import (
	"Hertz_refactored/biz/dal/db"
	"context"
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"

	"Hertz_refactored/biz/dal/cache"
	"Hertz_refactored/biz/model/user"
	"Hertz_refactored/biz/pkg/logging"
	"Hertz_refactored/biz/pkg/util"
)

type UserService struct {
	ctx context.Context
}

var wg sync.WaitGroup

// ToDo:通过调用这个函数，满足handler层可以通过这个接口调用service内的业务逻辑

func NewUserService(ctx context.Context) *UserService {
	return &UserService{ctx: ctx}
}

func (s *UserService) CreateUser(req user.CreateUserRequest, ctx context.Context) (*user.User, error) {
	var err error
	password, _ := util.Crypt(req.Password)
	User := &user.User{
		UserName: req.Name,
		Password: password,
	}

	wg.Add(1)
	go func() {
		//防止用户重复注册
		defer wg.Done()
		if _, err = NewUserService(ctx).VerifyUser(req.Name, req.Password); err == nil {
			logrus.Info("用户重复注册")
		}
	}()
	if err != nil {
		return nil, err
	}
	wg.Wait()
	return db.CreateUser(s.ctx, User)

}

func (s *UserService) LoginUser(req user.LoginUserResquest) (err error) {

	return nil

}

// ToDo:这是对JWT 登录认证时候的检验 通过这种切片的方式完成

func (s *UserService) VerifyUser(account, password string) (user.User, error) {
	var users user.User
	var err error
	if users, err = db.CheckUser(account, password); err != nil {
		logrus.Info(err)
		return users, err
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
	var err error
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = db.UpdateUser(User, userId); err != nil {
			logging.Info(err)
		}
	}()
	if err != nil {
		return err
	}
	wg.Wait()
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
	wg.Add(1)
	go func() {
		defer wg.Done()
		if User, total, err = db.QueryUser(req.Keyword, req.Page, req.PageSize); err != nil {
			logging.Error(err)
		}
	}()
	if err != nil {
		return User, total, err
	}
	wg.Wait()
	return User, total, err
}
