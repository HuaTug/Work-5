namespace go publish
enum Code2 {
     Success         = 1
     ParamInvalid    = 2
     DBErr           = 3
}

struct UpLoadVideoRequest{
    1: string ContentType     (api.body="content_type", api.form="content_type",api.vd="(len($) > 0 && len($) < 100)")
    2: string ObjectName    (api.body="object_name", api.form="object_name",api.vd="(len($) > 0 && len($) < 100)")
    3: string BucketName    (api.body="bucket_name", api.form="bucket_name",api.vd="(len($) > 0 && len($) < 100)")
}

struct UpLoadVideoResponse{
   1: Code2 code
   2: string msg
}


service UpLoadVideoService {
   UpLoadVideoResponse UploadVideo(1:UpLoadVideoRequest req)(api.post="/v1/video/upload")
}