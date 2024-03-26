package db

import (
	"Hertz_refactored/biz/model/comment"
	"github.com/sirupsen/logrus"
)

func CreateComment(comment *comment.Comment) error {
	return Db.Create(comment).Error
}
func Exist(req comment.CreateCommentRequest) bool {
	var count int64
	if Db.Model(&comment.Comment{}).Where("comment_id=? and video_id=?", req.IndexId, req.VideoId).Count(&count); count == 0 {
		logrus.Info("没有对应的视频评论")
		return false
	}
	return true
}
func DeleteComment(comments comment.CommentDeleteRequest) error {
	return Db.Model(&comment.Comment{}).Where("comment_id=? and video_id=?", comments.CommentId, comments.VideoId).Delete(comments).Error
}

func ListComment(List comment.ListCommentRequest) ([]*comment.Comment, int64, error) {
	if List.PageSize == 0 {
		List.PageSize = 15
	}
	var total int64
	var comments []*comment.Comment
	err := Db.Model(&comment.Comment{}).Where("video_id=?", List.VideoId).Count(&total).Limit(int(List.PageSize)).Offset(int((List.PageNum - 1) * List.PageSize)).Find(&comments).Error
	return comments, total, err
}
