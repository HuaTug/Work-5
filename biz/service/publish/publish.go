package publish

import (
	"Hertz_refactored/biz/dal/db"
	"Hertz_refactored/biz/model/comment"
	"Hertz_refactored/biz/model/publish"
	"Hertz_refactored/biz/model/video"
	"Hertz_refactored/biz/pkg/logging"
	"context"
	"fmt"
	minios "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
	"log"
	"mime/multipart"
	"os"
	"time"
)

func UploadFile(file *multipart.FileHeader, req publish.UpLoadVideoRequest, uid int64) error {
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	minioClient, err := minios.New("127.0.0.1:9000", &minios.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		logging.Error(err)
		return err
	}
	bucketName := req.BucketName
	objectName := req.ObjectName + "." + req.ContentType
	fmt.Println(objectName)

	exists, err3 := minioClient.BucketExists(context.Background(), bucketName)
	if err3 == nil && exists {
		logging.Info("Bucket %s already exists\n", bucketName)
	} else {
		err = minioClient.MakeBucket(context.Background(), bucketName, minios.MakeBucketOptions{})
		if err != nil {
			log.Fatalln(err)
		}
	}

	filePath := "C:\\Users\\0\\Downloads\\Video\\" + file.Filename
	fmt.Println(filePath)
	src, err := os.Open(filePath)
	if err != nil {
		logrus.Info("Open文件出错" + "err")
		return err
	}
	defer src.Close()
	_, err = minioClient.PutObject(context.Background(), bucketName, objectName, src, -1, minios.PutObjectOptions{})
	if err != nil {
		logrus.Info(err)
		return err
	}

	publishs := video.Video{
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
