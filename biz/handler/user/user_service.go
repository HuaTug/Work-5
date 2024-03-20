// Code generated by hertz generator.

package user

import (
	user2 "Hertz_refactored/biz/dal/db/user"
	user "Hertz_refactored/biz/model/user"
	"Hertz_refactored/biz/pack"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
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
	userId, _ := c.Get("user_id")
	resp := new(user.UpdateUserResponse)
	if uid, ok := userId.(float64); ok {
		User := &user.User{
			UserName: req.Name,
			Password: req.Password,
		}
		if err := user2.UpdateUser(User, int64(uid)); err != nil {
			c.JSON(consts.StatusBadRequest, err.Error())
			panic(err)
		}

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
	userId, _ := c.Get("user_id")
	resp := new(user.DeleteUserResponse)
	if uid, ok := userId.(float64); ok {
		if err = user2.DeleteUser(int64(uid)); err != nil {
			c.JSON(consts.StatusBadRequest, err.Error())
			panic(err)
		}
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
	if User, total, err = user2.QueryUser(req.Keyword, req.Page, req.PageSize); err != nil {
		c.JSON(consts.StatusBadRequest, err.Error())
		panic(err)
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
	User := &user.User{
		UserName: req.Name,
		Password: req.Password,
	}

	if err := user2.CreateUser(User); err != nil {
		panic(err)
	}
	resp := new(user.CreateUserResponse)
	resp.Code = consts.StatusOK
	resp.Msg = "创建用户成功"
	c.JSON(consts.StatusOK, resp)
}

// LoginUser .
// @router /v1/user/login [POST]
func LoginUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.LoginUserResquest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
}