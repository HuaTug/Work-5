package comment

import (
	"context"
	"encoding/json"

	"sync"

	"time"

	"Hertz_refactored/biz/dal/cache"
	"Hertz_refactored/biz/dal/db"
	"Hertz_refactored/biz/dal/db/mq"
	"Hertz_refactored/biz/model/comment"
	"Hertz_refactored/biz/pkg/logging"

	"github.com/sirupsen/logrus"
)

type CommentService struct {
	ctx context.Context
}


var wg sync.WaitGroup


func NewCommentService(ctx context.Context) *CommentService {
	return &CommentService{ctx: ctx}
}
func SendComment(data *comment.Comment) error {
	msg, err := json.Marshal(data)
	if err != nil {
		logrus.Info(err)
	}
	return mq.SendMessageMQ(msg)
}
func (s *CommentService) Create(req comment.CreateCommentRequest, userId int64) error {
	key:="comment_id"
	Id:=cache.GenerateID(key)
	comments := &comment.Comment{
		CommentId: Id,
		VideoId: req.VideoId,
		Comment: req.Comment,
		UserId:  userId,
		Time:    time.Now().Format(time.DateTime),
		IndexId: req.IndexId,
	}
	if req.ActionType == 1 && req.IndexId != 0 { //表示为非一级评论
		if flag := db.Exist(req); flag == false {
			comments.IndexId = 0
			if err := SendComment(comments); err != nil {
				logging.Error(err)
				return err
			}
			logging.Info("新插入一条评论成功")
		} else { //这样else后的逻辑就表示为一级逻辑
			if err := SendComment(comments); err != nil {
				logging.Error(err)
				return err
			}
		}
	} else {
		if err := SendComment(comments); err != nil {
			logging.Error(err)
			return err
		}
	}
	commentId := db.GetMaxId()+1
	go func() {
		err := cache.CacheSetCommentVideo(req.VideoId, commentId)
		if err != nil {
			logging.Error(err)
		}
	}()
	return nil
}

func (s *CommentService) Delete(req comment.CommentDeleteRequest) error {


	wg.Add(1)
	var err error
	go func() {
		defer wg.Done()
		if err = db.DeleteComment(req); err != nil {
			logging.Error(err)
		}
	}()
	if err != nil {
		return err
	}
	wg.Wait()

	return nil
}

func (s *CommentService) List(req comment.ListCommentRequest) ([]*comment.Comment, int64, error) {
	var comments []*comment.Comment
	var total int64
	var err error
	comments, err = cache.CacheGetListComment(req.VideoId)
	if err != nil {
		logging.Error(err)
		
	}else{
		return comments,total,err

	}
	comments, total, err = db.ListComment(req)
	if err != nil {
		return comments, total, err
	} else {
		go cache.CacheSetAllComment(req.VideoId, comments)
		return comments, total, err
	}
}
