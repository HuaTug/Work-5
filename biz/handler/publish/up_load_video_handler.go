// Code generated by hertz generator.

package publish

import (
	publish "Hertz_refactored/biz/model/publish"
	publish2 "Hertz_refactored/biz/service/publish"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/sirupsen/logrus"
)

// UploadVideo .
// @router /v1/video/upload [POST]
func UploadVideo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req publish.UpLoadVideoRequest

	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	userId, _ := c.Get("user_id")

	if v, ok := userId.(float64); ok { //ToDo :先将接口类型转化为interface接口类型后，再去转化为int64类型，而非直接断言从interface接口类型转化为int64类型
		uid := int64(v)
		file, err := c.FormFile("file")
		if err != nil {
			logrus.Info(err)
			c.JSON(consts.StatusBadRequest, "获取文件出错")
		}

		if err := publish2.UploadFile(file, req, uid); err != nil {
			logrus.Info("上传文件出错")
			c.JSON(consts.StatusBadRequest, "文件上传出错")
			return
		}
		resp := new(publish.UpLoadVideoResponse)
		resp.Msg = "上传成功"
		resp.Code = consts.StatusOK
		c.JSON(consts.StatusOK, resp)
	}
}
