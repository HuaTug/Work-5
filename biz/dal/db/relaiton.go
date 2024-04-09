package db

import (
	"Hertz_refactored/biz/model/relation"
	"github.com/sirupsen/logrus"
)

func Follow(Relation *relation.Relation) error {
	if err := Db.Model(&relation.Relation{}).Create(Relation).Error; err != nil {
		logrus.Info(err)
		return err
	}
	return nil
}

func UnFollow(Relation *relation.Relation) error {
	if err := Db.Model(&relation.Relation{}).Create(Relation).Error; err != nil {
		logrus.Info(err)
		return err
	}
	return nil
}

func FollowList(PageSize, PageNum, userId int64) error {
	var list []relation.Relation
	var count int64
	if err := Db.Model(&relation.Relation{}).Where("relation=?", userId).Count(&count).
		Limit(int(PageSize)).Offset(int((PageNum - 1) * PageSize)).Find(&list).Error; err != nil {
		logrus.Info("分页查询出错")
		logrus.Info(err)
	}
	return nil
}
