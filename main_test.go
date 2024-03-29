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
	h := hInit()
	token := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTQyMzUzNTAsIm9yaWdfaWF0IjoxNzExNjQzMzUwLCJ1c2VyX2lkIjoxfQ.XrBNo1vlHduekRoWHJplMCoZVm2K5Jldd2EwpzW9qqQ`
	// 将 token 放置在 form 表单数据中
	formData := url.Values{}
	formData.Set("token", token)

	// 设置请求的 Content-Type 为 application/x-www-form-urlencoded
	header := ut.Header{
		Key:   "Content-Type",
		Value: "application/x-www-form-urlencoded",
	}

	w := ut.PerformRequest(h.Engine, "GET", "/v1/user/get", &ut.Body{
		Body: bytes.NewBufferString(formData.Encode()), // 编码 form 表单数据
		Len:  len(formData.Encode()),                   // 计算编码后的长度
	}, header)

	// 检查响应状态码是否符合预期
	assert.DeepEqual(t, consts.StatusOK, w.Result().StatusCode())
}

func TestUpdateUser(t *testing.T) {
	h := hInit()
	token := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTQyMzUzNTAsIm9yaWdfaWF0IjoxNzExNjQzMzUwLCJ1c2VyX2lkIjoxfQ.XrBNo1vlHduekRoWHJplMCoZVm2K5Jldd2EwpzW9qqQ`
	// 将 token 放置在 form 表单数据中
	formData := url.Values{}
	formData.Set("token", token)
	formData.Set("name", "Nigtusg")
	formData.Set("password", "123456")
	// 设置请求的 Content-Type 为 application/x-www-form-urlencoded
	header := ut.Header{
		Key:   "Content-Type",
		Value: "application/x-www-form-urlencoded",
	}
	w := ut.PerformRequest(h.Engine, "POST", "/v1/user/update", &ut.Body{
		Body: bytes.NewBufferString(formData.Encode()), // 编码 form 表单数据
		Len:  len(formData.Encode()),                   // 计算编码后的长度
	}, header)
	// 检查响应状态码是否符合预期
	assert.DeepEqual(t, consts.StatusOK, w.Result().StatusCode())
}

func TestDeleteUser(t *testing.T) {
	h := hInit()
	token := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTQyMzUzNTAsIm9yaWdfaWF0IjoxNzExNjQzMzUwLCJ1c2VyX2lkIjoxfQ.XrBNo1vlHduekRoWHJplMCoZVm2K5Jldd2EwpzW9qqQ`
	// 将 token 放置在 form 表单数据中
	formData := url.Values{}
	formData.Set("token", token)

	// 设置请求的 Content-Type 为 application/x-www-form-urlencoded
	header := ut.Header{
		Key:   "Content-Type",
		Value: "application/x-www-form-urlencoded",
	}

	w := ut.PerformRequest(h.Engine, "DELETE", "/v1/user/delete", &ut.Body{
		Body: bytes.NewBufferString(formData.Encode()), // 编码 form 表单数据
		Len:  len(formData.Encode()),                   // 计算编码后的长度
	}, header)

	// 检查响应状态码是否符合预期
	assert.DeepEqual(t, consts.StatusOK, w.Result().StatusCode())
}

func TestVideoFee(t *testing.T) {
	h := hInit()
	//token := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTQyMzUzNTAsIm9yaWdfaWF0IjoxNzExNjQzMzUwLCJ1c2VyX2lkIjoxfQ.XrBNo1vlHduekRoWHJplMCoZVm2K5Jldd2EwpzW9qqQ`
	// 将 token 放置在 form 表单数据中
	lastime := "2024-03-28 13:12"
	formData := url.Values{}
	formData.Set("last_tim", lastime)

	// 设置请求的 Content-Type 为 application/x-www-form-urlencoded
	header := ut.Header{
		Key:   "Content-Type",
		Value: "application/x-www-form-urlencoded",
	}

	w := ut.PerformRequest(h.Engine, "GET", "/v1/feed", &ut.Body{
		Body: bytes.NewBufferString(formData.Encode()), // 编码 form 表单数据
		Len:  len(formData.Encode()),                   // 计算编码后的长度
	}, header)

	// 检查响应状态码是否符合预期
	assert.DeepEqual(t, consts.StatusOK, w.Result().StatusCode())
}

func TestListVide(t *testing.T) {
	h := hInit()
	token := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTQyMzUzNTAsIm9yaWdfaWF0IjoxNzExNjQzMzUwLCJ1c2VyX2lkIjoxfQ.XrBNo1vlHduekRoWHJplMCoZVm2K5Jldd2EwpzW9qqQ`
	// 将 token 放置在 form 表单数据中
	formData := url.Values{}
	formData.Set("token", token)
	formData.Set("name", "Nigtusg")
	formData.Set("password", "123456")
	// 设置请求的 Content-Type 为 application/x-www-form-urlencoded
	header := ut.Header{
		Key:   "Content-Type",
		Value: "application/x-www-form-urlencoded",
	}
	w := ut.PerformRequest(h.Engine, "POST", "/v1/user/update", &ut.Body{
		Body: bytes.NewBufferString(formData.Encode()), // 编码 form 表单数据
		Len:  len(formData.Encode()),                   // 计算编码后的长度
	}, header)
	// 检查响应状态码是否符合预期
	assert.DeepEqual(t, consts.StatusOK, w.Result().StatusCode())
}

func TestFeedList(t *testing.T) {
	h := hInit()
	// token := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTQyMzUzNTAsIm9yaWdfaWF0IjoxNzExNjQzMzUwLCJ1c2VyX2lkIjoxfQ.XrBNo1vlHduekRoWHJplMCoZVm2K5Jldd2EwpzW9qqQ`
	// 将 token 放置在 form 表单数据中
	pagenum := "1"
	pagesize := "5"
	authorid := "1"
	formData := url.Values{}
	formData.Set("page_size", pagesize)
	formData.Set("page_num", pagenum)
	formData.Set("author_id", authorid)
	// 设置请求的 Content-Type 为 application/x-www-form-urlencoded
	header := ut.Header{
		Key:   "Content-Type",
		Value: "application/x-www-form-urlencoded",
	}

	w := ut.PerformRequest(h.Engine, "GET", "/v1/video/list", &ut.Body{
		Body: bytes.NewBufferString(formData.Encode()), // 编码 form 表单数据
		Len:  len(formData.Encode()),                   // 计算编码后的长度
	}, header)

	// 检查响应状态码是否符合预期
	assert.DeepEqual(t, consts.StatusOK, w.Result().StatusCode())
}

// ToDo 网络测试
func BenchmarkVideoSearch(b *testing.B) {
	h := hInit()
	pagenum := "1"
	fromdate := "2024-03-16 13:39"
	todate := "2024-03-18 13:39"
	pagesize := "5"
	keyword := "心"
	formData := url.Values{}
	formData.Set("from_date", fromdate)
	formData.Set("page_num", pagenum)
	formData.Set("to_date", todate)
	formData.Set("page_size", pagesize)
	formData.Set("keyword", keyword)
	header := ut.Header{
		Key:   "Content-Type",
		Value: "application/x-www-form-urlencoded",
	}
	b.StopTimer()
	//当你调用 b.StopTimer() 时，它会暂停基准测试的计时器。这意味着在此调用之后执行的代码时间不会计入基准测试的持续时间。
	b.StartTimer()
	//在调用 b.StopTimer() 后，你可以使用 b.StartTimer() 来恢复基准测试的计时器。它表示实际进行基准测试的代码即将运行，计时器应该重新开始计时。
	// Perform the request and check the response status code
	for i := 0; i < b.N; i++ {
		w := ut.PerformRequest(h.Engine, "POST", "/v1/video/search", &ut.Body{
			Body: bytes.NewBufferString(formData.Encode()),
			Len:  len(formData.Encode()),
		}, header)
		if w.Result().StatusCode() != consts.StatusOK {
			b.Fatalf("Unexpected status code: %d", w.Result().StatusCode())
		}
	}
}

func TestVideoSea(t *testing.T) {
	h := hInit()
	// token := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTQyMzUzNTAsIm9yaWdfaWF0IjoxNzExNjQzMzUwLCJ1c2VyX2lkIjoxfQ.XrBNo1vlHduekRoWHJplMCoZVm2K5Jldd2EwpzW9qqQ`
	// 将 token 放置在 form 表单数据中
	pagenum := "1"
	fromdate := "2024-03-16 13:39"
	todate := "2024-03-18 13:39"
	pagesize := "5"
	keyword := "心"
	formData := url.Values{}
	formData.Set("from_date", fromdate)
	formData.Set("page_num", pagenum)
	formData.Set("to_date", todate)
	formData.Set("page_size", pagesize)
	formData.Set("keyword", keyword)
	// 设置请求的 Content-Type 为 application/x-www-form-urlencoded
	header := ut.Header{
		Key:   "Content-Type",
		Value: "application/x-www-form-urlencoded",
	}

	w := ut.PerformRequest(h.Engine, "POST", "/v1/video/search", &ut.Body{
		Body: bytes.NewBufferString(formData.Encode()), // 编码 form 表单数据
		Len:  len(formData.Encode()),                   // 计算编码后的长度
	}, header)

	// 检查响应状态码是否符合预期
	assert.DeepEqual(t, consts.StatusOK, w.Result().StatusCode())
}

func TestCommentCreate(t *testing.T) {
	h := hInit()
	token := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTQyMzUzNTAsIm9yaWdfaWF0IjoxNzExNjQzMzUwLCJ1c2VyX2lkIjoxfQ.XrBNo1vlHduekRoWHJplMCoZVm2K5Jldd2EwpzW9qqQ`
	// 将 token 放置在 form 表单数据中
	formData := url.Values{}
	formData.Set("token", token)
	formData.Set("action_type", "1")
	formData.Set("index_id", "5")
	// 设置请求的 Content-Type 为 application/x-www-form-urlencoded
	header := ut.Header{
		Key:   "Content-Type",
		Value: "application/x-www-form-urlencoded",
	}
	w := ut.PerformRequest(h.Engine, "POST", "/v1/comment/publish", &ut.Body{
		Body: bytes.NewBufferString(formData.Encode()), // 编码 form 表单数据
		Len:  len(formData.Encode()),                   // 计算编码后的长度
	}, header)
	// 检查响应状态码是否符合预期
	assert.DeepEqual(t, consts.StatusOK, w.Result().StatusCode())
}

func TestFollow(t *testing.T) {
	h := hInit()
	token := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTQyMzUzNTAsIm9yaWdfaWF0IjoxNzExNjQzMzUwLCJ1c2VyX2lkIjoxfQ.XrBNo1vlHduekRoWHJplMCoZVm2K5Jldd2EwpzW9qqQ`
	// 将 token 放置在 form 表单数据中
	formData := url.Values{}
	formData.Set("token", token)
	formData.Set("action_type", "1")
	formData.Set("comment_id", "5")
	// 设置请求的 Content-Type 为 application/x-www-form-urlencoded
	header := ut.Header{
		Key:   "Content-Type",
		Value: "application/x-www-form-urlencoded",
	}
	w := ut.PerformRequest(h.Engine, "POST", "/like/action", &ut.Body{
		Body: bytes.NewBufferString(formData.Encode()), // 编码 form 表单数据
		Len:  len(formData.Encode()),                   // 计算编码后的长度
	}, header)
	// 检查响应状态码是否符合预期
	assert.DeepEqual(t, consts.StatusOK, w.Result().StatusCode())
}

func TestFavorite(t *testing.T) {
	h := hInit()
	token := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTQyMzUzNTAsIm9yaWdfaWF0IjoxNzExNjQzMzUwLCJ1c2VyX2lkIjoxfQ.XrBNo1vlHduekRoWHJplMCoZVm2K5Jldd2EwpzW9qqQ`
	// 将 token 放置在 form 表单数据中
	formData := url.Values{}
	formData.Set("token", token)
	formData.Set("action_type", "1")
	formData.Set("to_user_id", "5")
	// 设置请求的 Content-Type 为 application/x-www-form-urlencoded
	header := ut.Header{
		Key:   "Content-Type",
		Value: "application/x-www-form-urlencoded",
	}
	w := ut.PerformRequest(h.Engine, "POST", "/v1/relation/action", &ut.Body{
		Body: bytes.NewBufferString(formData.Encode()), // 编码 form 表单数据
		Len:  len(formData.Encode()),                   // 计算编码后的长度
	}, header)
	// 检查响应状态码是否符合预期
	assert.DeepEqual(t, consts.StatusOK, w.Result().StatusCode())
}
