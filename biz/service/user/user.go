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

//
type UserService struct {
	ctx context.Context
}

// context的作用:在网络编程中，如果存在 A 调用 B 的 API，B 调用 C 的  API，那么假如 A 调用 B 取消，那么 B 调用 C 也应该被取消。

// ToDo:通过调用这个函数，满足handler层可以通过这个接口调用service内的业务逻辑

func NewUserService(ctx context.Context) *UserService {
	return &UserService{ctx: ctx}
}

//这是定义一个UserService的类型 然后再通过NewService去创建这个服务 将上下文context层层传递
//这是方法的定义格式，其中括号里面的内容是一个方法接收器 它指定了方法关联到了`UserService`类型，这样当调用某一个方法时，就需要先创建一个`UserService`类型的对象

func (s *UserService) CreateUser(req user.CreateUserRequest, ctx context.Context) (*user.User, error, bool) {
	var err error
	var flag bool
	var wg sync.WaitGroup
	key := "user_id"
	wg.Add(1)
	go func() {
		//防止用户重复注册
		defer wg.Done()
		if _, err, flag = NewUserService(ctx).VerifyUser(req.Name, req.Password); !flag {
			logrus.Info("用户重复注册")
		}
	}()
	
	wg.Wait()
	if !flag {
		return nil, err, flag
	}
	//使用Redis生成全局主键
	id := cache.GenerateID(key)

	password, _ := util.Crypt(req.Password)
	User := &user.User{
		UserID:   id,
		UserName: req.Name,
		Password: password,
	}
	return db.CreateUser(s.ctx, User)

}

func (s *UserService) LoginUser(req user.LoginUserResquest) (err error) {

	return nil

}

// ToDo:这是对JWT 登录认证时候的检验 通过这种切片的方式完成

func (s *UserService) VerifyUser(account, password string) (user.User, error, bool) {
	var users user.User
	var err error
	var flag bool
	if users, err, flag = db.CheckUser(account, password); err != nil {
		logrus.Info(err)
		return users, err, flag
	}
	return users, nil, flag

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
/*  	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = db.UpdateUser(User, userId); err != nil {
			logging.Info(err)
		}
	}()
	if err != nil {
		return err
	}
	wg.Wait()  */
	if err = db.UpdateUser(User, userId); err != nil {
		logging.Info(err)
		return err
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
/* 	wg.Add(1)
	go func() {
		defer wg.Done()
		if User, total, err = db.QueryUser(req.Keyword, req.Page, req.PageSize); err != nil {
			logging.Error(err)
		}

	}()
	if err != nil {
		return User, total, err
	}
	wg.Wait() */
	if User, total, err = db.QueryUser(req.Keyword, req.Page, req.PageSize); err != nil {
		logging.Error(err)
	}
	if err!=nil{
		return User,total,err
	}	
	return User, total, err
}
