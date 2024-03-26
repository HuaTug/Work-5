package pack

import (
	"Hertz_refactored/biz/model/user"
)

// 这是一个将切片的数据格式改为可以用于返回的整体数据结构
func Users(models []*user.User) []*user.User {
	users := make([]*user.User, 0, len(models))
	for _, m := range models {
		if u := User(m); u != nil {
			users = append(users, u)
		}
	}
	return users
}

func User(model *user.User) *user.User {
	if model == nil {
		return nil
	}
	return &user.User{
		UserID:   model.UserID,
		UserName: model.UserName,
		Password: model.Password,
	}
}
