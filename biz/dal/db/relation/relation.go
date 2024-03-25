package relation

import (
	user2 "Hertz_refactored/biz/dal/db/user"
	"Hertz_refactored/biz/dal/mysql"
	"fmt"

	"Hertz_refactored/biz/model/relation"
	"Hertz_refactored/biz/model/user"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/sirupsen/logrus"
)

const (
	add = int64(1)
	sub = int64(-1)
)

func Following(req relation.RelationServiceRequest, id int64) error {
	Relation := &relation.Relation{
		FollowId: req.ToUserId,
		UserId:   id,
	}
	if req.ActionType == 1 {
		if err := mysql.Db.Model(&relation.Relation{}).Create(Relation).Error; err != nil {
			logrus.Info(err)
			return err
		}
		go CacheChangeUserCount(id, add, "follower")
		go CacheChangeUserCount(req.ToUserId, add, "follow")
	} else {
		if err := mysql.Db.Model(&relation.Relation{}).Where("follow_id=?", req.ToUserId).Delete(&Relation).Error; err != nil {
			logrus.Info(err)
			return err
		}
		//这个表示被关注操作
		go CacheChangeUserCount(id, sub, "follower")
		//这个表示为关注操作
		go CacheChangeUserCount(req.ToUserId, sub, "follow")
	}
	return nil
}

func FollowList(req relation.RelationServicePageRequest, uid interface{}) relation.RelationServicePageResponse {
	var lists []*relation.Relation
	var users []*relation.User
	var count int64
	var userId int64
	if v, ok := uid.(float64); ok {
		userId = int64(v)
	} else {
		logrus.Info("转化出错")
	}
	if err := mysql.Db.Model(&relation.Relation{}).Where("user_id=?", userId).Count(&count).
		Limit(int(req.PageSize)).Offset(int((req.PageNum - 1) * req.PageSize)).Find(&lists).Error; err != nil {
		logrus.Info("分页查询出错")
		logrus.Info(err)
		return relation.RelationServicePageResponse{
			Codeh:     consts.StatusBadRequest,
			Msg:       "展示列表失败",
			Relaitons: nil,
			Users:     nil,
		}
	}
	if err := mysql.Db.Model(&user.User{}).Where("user_id=?", userId).Find(&users).Error; err != nil {
		logrus.Info(err)
		return relation.RelationServicePageResponse{
			Codeh:     consts.StatusBadRequest,
			Msg:       "展示列表失败",
			Relaitons: nil,
			Users:     nil,
		}
	}
	return relation.RelationServicePageResponse{
		Codeh:     consts.StatusOK,
		Msg:       "展示列表",
		Relaitons: lists,
		Users:     users,
	}
}

func FollowList2(req relation.RelationServicePageRequest, uid interface{}) error {
	var list []relation.Relation
	var count int64
	if err := mysql.Db.Model(&relation.Relation{}).Where("relation=?", uid.(int64)).Count(&count).
		Limit(int(req.PageSize)).Offset(int((req.PageNum - 1) * req.PageSize)).Find(&list).Error; err != nil {
		logrus.Info("分页查询出错")
		logrus.Info(err)
	}
	return nil
}

// ToDo 遇到的缓存识别问题 (已解决) ->注意Map下的redis映射关系
func CacheChangeUserCount(userid, op int64, types string) {
	username := user2.CacheGetIdAndName(userid)
	result := username[1 : len(username)-1]
	fmt.Println(result)
	users, err := user2.CacheGetUser(result)
	//fmt.Println(users.UserID, users.UserName, users.FollowCount)
	if err != nil {
		logrus.Printf("user:%v miss cache", userid)
		return
	}
	switch types {
	case "follow":
		users.FollowCount += op
	case "follower":
		users.FollowerCount += op
	case "like":
		users.FavoriteCount += op
	case "unlike":
		users.FavoriteCount += op
	}
	user2.CacheSetUser(users)
}
