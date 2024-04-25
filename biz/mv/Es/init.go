package es

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

var Client *elastic.Client

func Init() {
	var err error
	Client, err = elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetInfoLog(log.New(ioutil.Discard, "", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)),  //debug as os.stdout
		elastic.SetErrorLog(log.New(ioutil.Discard, "", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)), //debug as os.stderr
		elastic.SetTraceLog(log.New(ioutil.Discard, "", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)),
	)
	logrus.Info("Es initialize success")

	if err != nil {
		panic(err)
	}
	NewVideoIndex()
	index := &VideoIndex{Index: indexs}
	/* 	video := Video{
		VideoId:  1,
		AuthorId: 123,
		Info: VideoOtherData{
			PlayUrl:       "http://example.com/play",
			CoverUrl:      "http://example.com/cover",
			FavoriteCount: 100,
			CommentCount:  50,
			PublishTime:   "2024-04-26",
			Title:         "Test Video",
		},
	} */
	test := Video{
		VideoId:  2,
		AuthorId: 123,
		Info: VideoOtherData{
			PlayUrl:       "http://google.com",
			CoverUrl:      "http://google.com",
			FavoriteCount: 20,
			CommentCount:  10,
			//PublishTime:   TimeTranslate("2024-04-25 10:40:30"),
			Title:         "Update Video",
		},
	}
	index.CreateVideoDoc(test)
	//index.UpdateVideoDoc(test)
}

func TimeTranslate(data string) string {
	templete := "2006-01-02 15:04:05"
	t, err := time.Parse(templete, data)
	if err != nil {
		hlog.Info(err)
		return ""
	}
	return t.Format(templete)
}
