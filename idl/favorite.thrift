namespace go favorite

enum Code4 {
     Success         = 1
     ParamInvalid    = 2
     DBErr           = 3
}
struct User {
    1: i64 user_id
    2: string user_name
    3: i64 follow_count
    4: i64 follower_count
    5: i64 favorite_count
    6: string password
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

struct Favorite{
    1:i64 FavoriteId
    2:i64 UserId
    3:i64 VideoId
}

struct FavoriteRequest{
    1: i64    ActionType    (api.body="action_type", api.form="action_type",api.vd="($==0 ||$==1)")
    2: i64    VideoId      (api.body="video_id", api.form="video_id",api.vd="$>0")
}

struct FavoriteResponse{
   1: Code4 code
   2: string msg
}

struct ListFavoriteRequest{
}

struct ListFavoriteResponse{
    1: Code4 code
    2: string msg
    3: list<Favorite> favs
    4:list <User> users
}

service FavoriteService{
    FavoriteResponse FavoriteService (1:FavoriteRequest req)(api.post="/like/action")
    ListFavoriteResponse ListFavorite(1:ListFavoriteRequest req)(api.get="/like/list")
}