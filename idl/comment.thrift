namespace go comment

enum Code5 {
     Success         = 1
     ParamInvalid    = 2
     DBErr           = 3
}
struct Comment{
    1: i64 VideoId
    2: i64 CommentId
    3: string Comment
    4: string Time
    5: i64 UserId
}

struct CreateCommentRequest{
    1: i64 VideoId (api.body="video_id",api.form="video_id",api.vd="$>0")
    2: string Comment (api.body="comment",api.form="comment",api.vd="$>0")
}

struct CreateCommentResponse{
   1: Code5 code
   2: string msg
}

struct ListCommentRequest{
    1:i64 PageNum        (api.body="page_num",api.form="page_num")
    2:i64 PageSize       (api.body="page_size",api.form="page_size")
}

struct ListCommentResponse{
    1: Code5 code
    2: string msg
    3: list <Comment> comments
    4: i64  total
}

struct CommentDeleteRequest{
    1: i64    VideoId      (api.body="video_id", api.form="video_id",api.vd="$>0")
    2: string CommentId    (api.body="comment_id", api.form="comment_id",api.vd="(len($) > 0 && len($) < 1000)")
}
struct CommentDeleteResponse{
    1:Code5 code
    2:string msg
}
service CommentService {
   CreateCommentResponse CreateComment(1:CreateCommentRequest req)(api.post="/v1/comment/publish")
   ListCommentResponse ListComment(1:ListCommentRequest req)(api.get="/v1/comment/list")
   CommentDeleteResponse DeleteComment(1:CommentDeleteRequest req)(api.delete="/v1/comment/delete")
}