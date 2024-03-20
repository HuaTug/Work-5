namespace go user


enum Code {
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

struct CreateUserRequest{
    1: string name      (api.body="name", api.form="name",api.vd="(len($) > 0 && len($) < 100)")
    2: string password (api.body="password",api.form="password",api.vd="len($)>5 &&len($)<12")
}

struct CreateUserResponse{
   1: Code code
   2: string msg
}

struct QueryUserRequest{
   1: optional string Keyword (api.body="keyword", api.form="keyword",api.query="keyword")
   2: i64 page (api.body="page", api.form="page",api.query="page",api.vd="$ > 0")
   3: i64 page_size (api.body="page_size", api.form="page_size",api.query="page_size",api.vd="($ > 0 || $ <= 100)")
}

struct QueryUserResponse{
   1: Code code
   2: string msg
   3: list<User> users
   4: i64 totoal
}

struct DeleteUserRequest{

}

struct DeleteUserResponse{
   1: Code code
   2: string msg
}

struct UpdateUserRequest{
    1: string name      (api.body="name", api.form="name",api.vd="(len($) > 0 && len($) < 100)")
    3: string password (api.body="password",api.form="password",api.vd="(len($)>5 &&len($)<12)")
}

struct UpdateUserResponse{
   1: Code code
   2: string msg
}

struct LoginUserResquest{
    1:string Username   (api.body="user_name",api.form="user_name",api.vd="(len($)>0&&len($)<100)")
    2:string Password   (api.body="password",api.form="password",api.vd="(len($)>0&&len($)<100)")
}

struct LoginUserResponse{
    1:Code code
    2:string msg
    3:string token
}

service UserService {
   UpdateUserResponse UpdateUser(1:UpdateUserRequest req)(api.post="/v1/user/update")
   DeleteUserResponse DeleteUser(1:DeleteUserRequest req)(api.delete="/v1/user/delete")
   QueryUserResponse  QueryUser(1: QueryUserRequest req)(api.post="/v1/user/query/")
   CreateUserResponse CreateUser(1:CreateUserRequest req)(api.post="/v1/user/create/")
   LoginUserResponse  LoginUser(1:LoginUserResquest req)(api.post="/v1/user/login")
}