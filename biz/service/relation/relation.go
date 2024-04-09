package relation

import (
	"context"

	"github.com/sirupsen/logrus"

	"Hertz_refactored/biz/dal/cache"
	"Hertz_refactored/biz/dal/db"
	"Hertz_refactored/biz/model/relation"
)

const (
	add = int64(1)
	sub = int64(-1)
)

type RelationService struct {
	ctx context.Context
}

func NewRelationService(ctx context.Context) *RelationService {
	return &RelationService{ctx: ctx}
}
func (s *RelationService) Following(req relation.RelationServiceRequest, id int64) error {
	Relation := &relation.Relation{
		FollowId: req.ToUserId,
		UserId:   id,
	}
	if req.ActionType == 1 {
		if err := db.Follow(Relation); err != nil {
			return err
		}
		go CacheChangeUserCount(id, add, "follow")
		go CacheChangeUserCount(req.ToUserId, add, "follower")
	} else {
		if err := db.UnFollow(Relation); err != nil {
			return err
		}
		//这个表示被关注操作
		go CacheChangeUserCount(id, sub, "follow")
		//这个表示为关注操作
		go CacheChangeUserCount(req.ToUserId, sub, "follower")
	}
	return nil
}

func (s *RelationService) FollowList(req relation.RelationServicePageRequest, userId int64) error {
	if err := db.FollowList(req.PageSize, req.PageNum, userId); err != nil {
		return err
	}
	return nil
}

/*func FollowList(req relation.RelationServicePageRequest, v interface{}) relation.RelationServicePageResponse {
	var lists []*relation.Relation
	var users []*relation.User
	var err error
	userId:=utils.Transfer(v)
	if err := db.FollowList(req.PageSize,req.PageNum,userId);err!=nil{
		return relation.RelationServicePageResponse{
			Codeh:     consts.StatusBadRequest,
			Msg:       "展示列表失败",
			Relaitons: nil,
			Users:     nil,
		}
	}
	if _,err = db.GetUser(userId); err != nil {
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
*/

// ToDo 遇到的缓存识别问题 (已解决) ->注意Map下的redis映射关系
func CacheChangeUserCount(userid, op int64, types string) {
	users, err := cache.CacheGetUser(userid)
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
	cache.CacheSetUser(users)
}
