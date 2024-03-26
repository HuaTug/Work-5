package comment

import (
	"Hertz_refactored/biz/dal/cache"
	"Hertz_refactored/biz/dal/db"
	"Hertz_refactored/biz/model/comment"
	"Hertz_refactored/biz/pkg/logging"
	"context"
	"errors"
	"time"
)

type CommentService struct {
	ctx context.Context
}

func NewCommentService(ctx context.Context) *CommentService {
	return &CommentService{ctx: ctx}
}

func (s *CommentService) Create(req comment.CreateCommentRequest, userId int64) error {
	comments := &comment.Comment{
		VideoId: req.VideoId,
		Comment: req.Comment,
		UserId:  userId,
		Time:    time.Now().Format(time.DateTime),
		IndexId: req.IndexId,
	}
	if req.ActionType == 1 {
		if flag := db.Exist(req); flag == false {
			return errors.New("无法找到对应的评论")
		} else {
			if err := db.CreateComment(comments); err != nil {
				logging.Error(err)
				return err
			}
		}
	} else {
		if err := db.CreateComment(comments); err != nil {
			logging.Error(err)
			return err
		}
	}
	return nil
}

func (s *CommentService) Delete(req comment.CommentDeleteRequest) error {
	if err := db.DeleteComment(req); err != nil {
		logging.Error(err)
		return err
	}
	return nil
}

func (s *CommentService) List(req comment.ListCommentRequest) ([]*comment.Comment, int64, error) {
	var comments []*comment.Comment
	var total int64
	var err error
	comments, err = cache.CacheGetListComment(req.VideoId)
	if err != nil {
		logging.Error(err)
	}
	comments, total, err = db.ListComment(req)
	if err != nil {
		return comments, total, err
	} else {
		go cache.CacheSetAllComment(req.VideoId, comments)
		return comments, total, err
	}
}
