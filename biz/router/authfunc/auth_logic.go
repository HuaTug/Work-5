package authfunc

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	jwt "Hertz_refactored/biz/mv"
)

type Res struct {
	Msg  string
	Code int64
}

func Auth() []app.HandlerFunc {
	return append(make([]app.HandlerFunc, 0),
		DoubleTokenAuthFunc(),
		//jwt.AccessTokenJwtMiddleware.MiddlewareFunc(),
	)
}

func DoubleTokenAuthFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		if !jwt.IsAccessTokenAvailable(ctx, c) {
			if !jwt.IsRefreshTokenAvailable(ctx, c) {
				resp := new(Res)
				resp.Code = consts.StatusBadRequest
				resp.Msg = "Token is Inavailable"
				c.JSON(consts.StatusOK, resp)
				c.Abort()
				return
			}

			//此时表示refresh-token并未过期 在生成一个新的access-token
			//resp:=new(Res)

			//ToDo 
			jwt.GenerateAccessToken(ctx, c)
		}
		c.Next(ctx)
	}
}
