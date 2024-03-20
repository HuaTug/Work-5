package pack

import (
	"Hertz_refactored/biz/model/user"
)

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
