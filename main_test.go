package main

import (
	"Hertz_refactored/biz/config"
	"Hertz_refactored/biz/dal/cache"
	"Hertz_refactored/biz/dal/db"
	"Hertz_refactored/biz/dal/db/mq/script"
	"Hertz_refactored/biz/mv"
	"bytes"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"net/url"
	"testing"
)

func hInit() *server.Hertz {
	config.Init()
	mv.InitJwt()
	db.Init()
	cache.Init()
	script.LoadingScript()
	h := server.Default()
	register(h)
	return h
}

func TestUserRegister(t *testing.T) {
	h := hInit()
	req := `{
		"username":"test",
		"password":"123456789"
	}
	`
	header := ut.Header{
		Key:   "Content-Type",
		Value: "application/x-www-form-urlencoded",
	}
	ut.PerformRequest(h.Engine, "POST", "v1/user/create", &ut.Body{Body: bytes.NewBufferString(req), Len: len(req)},
		header,
	)
}

func TestUserInfo(t *testing.T) {
	token := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTQyMzUzNTAsIm9yaWdfaWF0IjoxNzExNjQzMzUwLCJ1c2VyX2lkIjoxfQ.XrBNo1vlHduekRoWHJplMCoZVm2K5Jldd2EwpzW9qqQ`
	h := hInit()

	// 将 token 放置在 form 表单数据中
	formData := url.Values{}
	formData.Set("token", token)

	// 设置请求的 Content-Type 为 application/x-www-form-urlencoded
	header := ut.Header{
		Key:   "Content-Type",
		Value: "application/x-www-form-urlencoded",
	}

	// 发送 POST 请求，将 form 表单数据作为请求主体发送
	w := ut.PerformRequest(h.Engine, "GET", "/v1/user/get/", &ut.Body{
		Body: bytes.NewBufferString(formData.Encode()), // 编码 form 表单数据
		Len:  len(formData.Encode()),                   // 计算编码后的长度
	}, header)

	// 检查响应状态码是否符合预期
	assert.DeepEqual(t, consts.StatusOK, w.Result().StatusCode())
}
