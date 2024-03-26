package mv

import (
	"Hertz_refactored/biz/model/user"
	"Hertz_refactored/biz/pkg/logging"
	user_service "Hertz_refactored/biz/service/user"
	utils2 "Hertz_refactored/biz/utils"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
	"net/http"
	"time"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	identity      = "user_id"
)

func InitJwt() {
	var err error

	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour * 24 * 30,
		MaxRefresh:  time.Second * 2,
		IdentityKey: identity,
		TokenLookup: "query:token,form:token",
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginRequest user.LoginUserResquest
			if err := c.BindAndValidate(&loginRequest); err != nil {
				return nil, err
			}
			users, err := user_service.CheckUser(loginRequest.Username, loginRequest.Password)
			if err != nil {
				c.JSON(http.StatusBadRequest, "登录失败")
				logging.Error(err)
				return nil, err
			}
			if users.UserName == "" || users.Password == "" {
				return nil, errors.NewPublic("user already exists or wrong password")
			}
			c.Set("user_id", users.UserID)
			//生成refresh_token , 并且设置键值对映射
			_, refretoken, _ := utils2.GenerateToken(users.UserID, users.UserName)
			c.Set("refresh", refretoken)
			return users.UserID, nil
		},

		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {

			c.Set("token", token)
			v, _ := c.Get("refresh")
			refresh := v.(string)
			c.JSON(http.StatusOK, utils.H{
				"code":          code,
				"access_token":  token,
				"refresh_token": refresh,
				"expire":        expire.Format(time.DateTime),
				"message":       "success",
			})
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					identity: v,
				}
			}
			return jwt.MapClaims{}
		},

		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt biz err = %+v", e.Error())
			logging.Error(e)
			return e.Error()
		},

		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"code":    code,
				"message": message,
			})
		},
	})
	if err != nil {
		panic(err)
	}
}
