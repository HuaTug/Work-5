namespace go relation

enum Code3 {
     Success         = 1
     ParamInvalid    = 2
     DBErr           = 3
}

struct Relation{
    1:i64 RelationId
    2:i64 FollowerId
    3:i64 FollowId
    4:i64 UserId
}
struct User {
    1: i64 user_id
    2: string user_name
    3: i64 follow_count
    4: i64 follower_count
    5: i64 favorite_count
    6: string password
}

struct RelationServiceRequest{
    1: i64    ActionType    (api.body="action_type", api.form="action_type",api.vd="($==0 ||$==1)")
    2: i64    ToUserId      (api.body="to_user_id", api.form="to_user_id",api.vd="$>0")
}

struct RelationServiceResponse{
   1: Code3 code
   2: string msg
}

struct RelationServicePageRequest{
   1:i64 PageNum (api.body="page_num",api.form="page_num",api.vd="$>0")
   2:i64 PageSize (api.body="page_size",api.form="page_size",api.vd="$>0")
}

struct RelationServicePageResponse{
    1: Code3 codeh
    2: string msg
    3:list<Relation> relaitons
    4:list<User> users
}

service FollowService {
   RelationServiceResponse RelationService (1:RelationServiceRequest req)(api.post="/v1/relation/action")
   RelationServicePageResponse RelationServicePage (1:RelationServicePageRequest req)(api.get="/v1/relation/list")
}