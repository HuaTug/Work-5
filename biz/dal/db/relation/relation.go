package relation

import (
	"Hertz_refactored/biz/dal/mysql"
	"Hertz_refactored/biz/model/relation"
	"Hertz_refactored/biz/model/user"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/sirupsen/logrus"
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
	} else {
		if err := mysql.Db.Model(&relation.Relation{}).Where("follow_id=?", req.ToUserId).Delete(&Relation).Error; err != nil {
			logrus.Info(err)
			return err
		}
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
