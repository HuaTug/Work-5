package comment

import (
	"Hertz_refactored/biz/dal/mysql"
	"Hertz_refactored/biz/model/comment"
)

func CreateComment(comment *comment.Comment) error {
	return mysql.Db.Create(comment).Error
}

func DeleteComment(comments comment.CommentDeleteRequest) error {
	return mysql.Db.Model(&comment.Comment{}).Where("comment_id=? and video_id=?", comments.CommentId, comments.VideoId).Delete(comments).Error
}

func ListComment(List comment.ListCommentRequest, uid int64) ([]*comment.Comment, int64, error) {
	if List.PageSize == 0 {
		List.PageSize = 15
	}
	var total int64
	var comments []*comment.Comment
	err := mysql.Db.Model(&comment.Comment{}).Where("user_id=?", uid).Count(&total).Limit(int(List.PageSize)).Offset(int((List.PageNum - 1) * List.PageSize)).Find(&comments).Error
	return comments, total, err
}
