// Code generated by hertz generator.

package favorite

import (
	"Hertz_refactored/biz/router/authfunc"

	"github.com/cloudwego/hertz/pkg/app"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _likeMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _favoriteserviceMw() []app.HandlerFunc {
	return authfunc.Auth()
}

func _listfavoriteMw() []app.HandlerFunc {
	return authfunc.Auth()
}
