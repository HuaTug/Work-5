package publish

import (
	"Hertz_refactored/biz/dal/mysql"
	"Hertz_refactored/biz/model/comment"
	"Hertz_refactored/biz/model/publish"
	"Hertz_refactored/biz/model/video"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
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
	minioClient, err := minio.New("127.0.0.1:9000", &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})
	bucketName := req.BucketName
	objectName := req.ObjectName + "." + req.ContentType
	fmt.Println(objectName)
	if err != nil {
		exists, err3 := minioClient.BucketExists(context.Background(), bucketName)
		if err3 == nil && exists {
			log.Printf("Bucket %s already exists\n", bucketName)
		} else {
			log.Fatalln("Failed to create bucket:", err)
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
	_, err = minioClient.PutObject(context.Background(), bucketName, objectName, src, -1, minio.PutObjectOptions{})
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
	logrus.Info("视频文件上传成功")
	mysql.Db.Model(&video.Video{}).Create(publishs)
	return nil
}
