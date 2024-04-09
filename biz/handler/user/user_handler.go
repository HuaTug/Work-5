// Code generated by hertz generator.

package user

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	user "Hertz_refactored/biz/model/user"
	"Hertz_refactored/biz/pack"
	"Hertz_refactored/biz/pkg/logging"
	"Hertz_refactored/biz/pkg/utils"
	user_service "Hertz_refactored/biz/service/user"
)

// UpdateUser .
// @router /v1/user/update/:user_id [POST]
func UpdateUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UpdateUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	v, _ := c.Get("user_id")
	userId := utils.Transfer(v)
	resp := new(user.UpdateUserResponse)
	if err = user_service.NewUserService(ctx).Update(req, userId); err == nil {
		resp.Code = consts.StatusOK
		resp.Msg = "用户更新成功"
		c.JSON(consts.StatusOK, resp)
		return
	} else {
		resp.Code = consts.StatusBadRequest
		resp.Msg = "用户更新失败"
		c.JSON(consts.StatusBadRequest, resp)
	}
}

// DeleteUser .
// @router /v1/user/delete/:user_id [POST]
func DeleteUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.DeleteUserRequest
	err = c.BindAndValidate(&req)
	v, _ := c.Get("user_id")
	userId := utils.Transfer(v)
	resp := new(user.DeleteUserResponse)
	if err = user_service.NewUserService(ctx).Delete(userId); err == nil {
		resp.Code = consts.StatusOK
		resp.Msg = "用户删除成功"
		c.JSON(consts.StatusOK, resp)
	} else {
		resp.Code = consts.StatusBadRequest
		resp.Msg = "删除用户失败"
		c.JSON(consts.StatusBadRequest, resp)
	}
}

// QueryUser .
// @router /v1/user/query/ [POST]
func QueryUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.QueryUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	var User []*user.User
	var total int64
	if User, total, err = user_service.NewUserService(ctx).Query(&req); err != nil {
		c.JSON(consts.StatusBadRequest, err.Error())
		logging.Error(err)
	}
	c.JSON(consts.StatusOK, &user.QueryUserResponse{Code: user.Code_Success, Users: pack.Users(User), Totoal: total})
}

// CreateUser .
// @router /v1/user/create/ [POST]
func CreateUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.CreateUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp := new(user.CreateUserResponse)
	userResp, err := user_service.NewUserService(ctx).CreateUser(req)
	resp.Code = consts.StatusOK
	resp.Msg = "创建用户成功"
	c.JSON(consts.StatusOK, resp)
	c.JSON(consts.StatusOK, user.CreateUserResponse{User: userResp})
}

// LoginUser .
// @router /v1/user/login [POST]
func LoginUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.LoginUserResquest
	err = c.BindAndValidate(&req)
	if err != nil {
		logging.Info(err)
		c.JSON(consts.StatusBadRequest, "登录失败")
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
}

// GetUserInfo .
// @router /v1user/get [GET]
func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.GetUserInfoRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		logging.Error(err)
		return
	}
	v, _ := c.Get("user_id")
	userId := utils.Transfer(v)
	//go CacheIdAndName(utils.Transfer(userId), req.UserName)
	/*
		这是一个Redis缓存的一个逻辑，即为用户先从redis缓存中根据关键字去查询数据，如果存在，
		便可以直接从缓存中获取数据，如果数据不存在，那么便去从数据库中查询数据，如果数据库中查到了，
		那么便将其加载到redis缓存中，以便后来去查询
	*/
	Resp, err := user_service.NewUserService(ctx).GetInfo(userId)
	if err != nil {
		return
	}
	resp := new(user.GetUserInfoResponse)
	resp.Code = consts.StatusOK
	resp.User = Resp
	c.JSON(consts.StatusOK, resp)
}

// ToDo :后续为了安全起见 需要再对其进行一个优化 tips：尝试对登录操作时，就完成映射关系
