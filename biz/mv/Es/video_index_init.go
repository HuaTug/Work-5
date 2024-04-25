package es

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"strconv"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/olivere/elastic/v7"
)

type VideoOtherData struct {
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	//PublishTime   string `json:"publish_time"`
	Title string `json:"title"`
}
type Video struct {
	VideoId  int64          `json:"video_id'`
	AuthorId int64          `json:"author_id"`
	Info     VideoOtherData `json:"info`
}

const indexs = "videos"

const videoMapping = `{
	"mappings":{
	  "properties":{
		"video_id":{
			"type":"long"
		},
		"author_id":{
			"type":"long"
		},
		"info":{
			"properties":{
			"play_url":{
				"type":"text"
			},
			"cover_url":{
				"type":"text"
			},
			"favorite_count":{
				"type":"long"
			},
			"comment_count":{
				"type":"long"
			},
			"title":{
				"type":"text"
			}				
			}
		}
	  }
	}
}`

type VideoIndex struct {
	Index   string
	Mapping string
}

func NewVideoIndex() (*VideoIndex, error) {
	video := VideoIndex{
		Index:   indexs,
		Mapping: videoMapping,
	}
	if err := video.Create(); err != nil {
		hlog.Info("Fail to Create!")
		return nil, err
	} else {
		return &video, nil
	}
}

func (v *VideoIndex) Create() error {
	ctx := context.Background()

	exist, err := Client.IndexExists(indexs).Do(ctx)
	if err != nil {
		hlog.Info(err)
		return err
	}
	if !exist {
		create, err := Client.CreateIndex(indexs).Body(v.Mapping).Do(ctx)
		if err != nil {
			hlog.Info(err)
			return err
		}

		if create.Acknowledged {
			hlog.Info("Elasticsearch index[video] initialized")
		}
	}
	return nil
}

func (v *VideoIndex) CreateVideoDoc(data Video) error {
	video_id := strconv.FormatInt(int64(data.VideoId), 10)
	_, err := Client.Index().Index(v.Index).Id(video_id).BodyJson(data).Do(context.Background())
	if err != nil {
		hlog.Info("Create doc failed: ", err)
		return err
	}
	return nil
}

func (v *VideoIndex) DeleteVideoDoc(vid int64) error {
	videoId := strconv.FormatInt(int64(vid), 10)
	_, err := Client.Delete().Index(v.Index).Id(videoId).Do(context.Background())
	if err != nil {
		hlog.Info("Delete doc failed: ", err)
		return err
	}
	return nil
}

func (v *VideoIndex) SearchVideoDoc(vid int64) (*Video, error) {
	videoId := strconv.FormatInt(int64(vid), 10)
	resp, err := Client.Get().Index(v.Index).Id(videoId).Do(context.Background())
	if err != nil {
		hlog.Info("Search doc failed: ", err)
		return nil, err
	}
	if resp.Found {
		var video Video
		json.Unmarshal(resp.Source, &video)
		return &video, nil
	} else {
		return nil, errors.New("Data not found")
	}
}

func (v *VideoIndex) UpdateVideoDoc(video Video) error {
	videoId := strconv.FormatInt(int64(video.VideoId), 10)
	hlog.Info(videoId)
	_, err := Client.Update().Index(v.Index).Id(videoId).Doc(video).Do(context.Background())
	if err != nil {
		hlog.Info("Update doc failed: ", err)
		return err
	}
	return nil
}

func (v *VideoIndex) searchRespCovert(resp *elastic.SearchResult) ([]*Video, int64) {
	dataList := make([]*Video, 0)
	for _, item := range resp.Each(reflect.TypeOf(Video{})) {
		data, ok := item.(Video)
		if !ok {
			continue
		}
		temp := Video{
			VideoId:  data.VideoId,
			AuthorId: data.AuthorId,
			Info:     data.Info,
		}
		dataList = append(dataList, &temp)
	}
	hits := resp.TotalHits()
	return dataList, hits
}

func (v *VideoIndex) SearchVideoDocDefault(keyword string) ([]*Video, int64, error) {
	resp, err := Client.Search().Index(v.Index).
		//同时需要注意复合字段的匹配规则 形式也有所不同 
		//.Should表示只需满足其中的一个查询条件即可
		Query(elastic.NewBoolQuery().Should(
			elastic.NewMultiMatchQuery(keyword, "Info.title"),
			//使用通配符进行字段数据匹配
			elastic.NewWildcardQuery("Info.title", "*"+keyword+"*"),
		)).
		//Sort("publish_time", true).
		From(0).Size(10).
		Do(context.Background())
	if err != nil {
		return nil, -1, err
	}
	data, hits := v.searchRespCovert(resp)
	return data, hits, nil
}

/* func SearchVideoDoc(keywords,username,fromDate,toDate string, pageNum,pageSize int64)([]*Video,int64,error){

} */
