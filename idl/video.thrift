namespace go video


enum Code1 {
     Success         = 1
     ParamInvalid    = 2
     DBErr           = 3
}

enum Gender {
    Unknown = 0
    Male    = 1
    Female  = 2
}

struct Video{
    1:i64 VideoId
    2:i64 AuthorId
    3:string PlayUrl
    4:string CoverUrl
    5:i64 FavoriteCount
    6:i64 CommentCount
    7:string PublishTime
    8:string Title
}
struct FeedServiceRequest{
    1:string LastTime  (api.body="last_time",api.form="last_time")
}
struct FeedServiceResponse{
   1: Code1 code
   2: string msg
   3: list<Video> video_list
}

struct VideoFeedListRequest{
    1: i64  AuthorId     (api.body="author_id", api.form="author_id")
    2: i64   PageNum      (api.body="page_num", api.form="page_num")
    3: i64 PageSize (api.body="page_size", api.form="page_size")
}

struct VideoFeedListResponse{
   1: Code1 code
   2: string msg
   3:list<Video> video_list
   4:i64 count
}

struct VideoSearchRequest{
   1: string Keyword (api.body="keyword", api.form="keyword",api.query="keyword")
   2: i64   PageNum     (api.body="page_num", api.form="page_num")
   3: i64 PageSize (api.body="page_size", api.form="page_size")
   4: string FromDate (api.body="from_date",api.form="from_date")
   5: string ToDate    (api.body="to_date",api.form="to_date")
}
struct VideoSearchResponse{
   1: Code1 code
   2: string msg
   3:list<Video> video_search
   4:i64 count
}

struct VideoPopularRequest{
}

struct VideoPopularResponse{
    1:Code1 code
    2:string msg
}


service VideoService {
   FeedServiceResponse FeedService(1:FeedServiceRequest req)(api.get="/v1/feed")
   VideoFeedListResponse VideoFeedList(1:VideoFeedListRequest req)(api.get="/v1/video/list")
   VideoSearchResponse  VideoSearch(1: VideoSearchRequest req)(api.post="/v1/video/search")
   VideoPopularResponse VideoPopular(1:VideoPopularRequest req)(api.get="/v1/video/popular")
}
