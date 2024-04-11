package publish

import (
	"Hertz_refactored/biz/dal/cache"
	"Hertz_refactored/biz/dal/db"
	"Hertz_refactored/biz/model/comment"
	"Hertz_refactored/biz/model/publish"
	"Hertz_refactored/biz/model/video"
	"Hertz_refactored/biz/pkg/logging"
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"sync"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
)

func UploadFile(file *multipart.FileHeader, req publish.UpLoadVideoRequest, uid int64) error {
	var wg sync.WaitGroup
	key := "video_id"
	Id := cache.GenerateID(key)
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	minioClient, err := minio.New("localhost:9000", &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		logging.Error(err)
		return err
	}
	bucketName := req.BucketName
	var filePath string
	var objectName string
	wg.Add(2)
	go func() {
		defer wg.Done()
		exists, err3 := minioClient.BucketExists(context.Background(), bucketName)
		if err3 == nil && exists {
			logrus.Printf("Bucket %s already exists\n", bucketName)
		} else {
			err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
			if err != nil {
				log.Fatalln(err)
			}
		}
	}()
	switch req.ContentType {
	case "video/mp4":
		filePath = "/home/xuzh/Videos/" + file.Filename
		objectName = req.ObjectName + ".mp4"
	case "png", "jpg", "jpeg":
		filePath = "/home/xuzh/Pictures/" + file.Filename
		objectName = req.ObjectName + ".jpg"
	}

	fmt.Println(filePath)
	src, err := os.Open(filePath)
	if err != nil {
		logrus.Info("Open文件出错" + "err")
		return err
	}
	defer src.Close()
	go func() {
		_, err = minioClient.PutObject(context.Background(), bucketName, objectName, src, -1, minio.PutObjectOptions{})
		if err != nil {
			logrus.Info(err)
		}
		wg.Done()
	}()
	wg.Wait()

	publishs := video.Video{
		VideoId:     Id,
		PlayUrl:     filePath,
		CoverUrl:    comment.Comment{}.Comment,
		PublishTime: time.Now().Format(time.DateTime),
		Title:       file.Filename,
		AuthorId:    uid,
	}
	if err := db.VideoCreate(publishs); err != nil {
		logging.Error(err)
		return err
	}
	logrus.Info("文件上传成功")
	return nil
}
