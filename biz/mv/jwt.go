package mv

import (
	user2 "Hertz_refactored/biz/dal/db/user"
	"Hertz_refactored/biz/model/user"
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
		MaxRefresh:  time.Hour * 24 * 30,
		IdentityKey: identity,
		TokenLookup: "query:token,form:token",
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginRequest user.LoginUserResquest
			if err := c.BindAndValidate(&loginRequest); err != nil {
				return nil, err
			}
			users, err := user2.CheckUser(loginRequest.Username, loginRequest.Password)
			if err != nil {
				panic(err)
			}
			if len(users) == 0 {
				return nil, errors.NewPublic("user already exists or wrong password")
			}
			c.Set("user_id", users[0].UserID)
			return users[0].UserID, nil
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.Set("token", token)
			c.JSON(http.StatusOK, utils.H{
				"code":    code,
				"token":   token,
				"expire":  expire.Format(time.DateTime),
				"message": "success",
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
