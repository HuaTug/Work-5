├─.idea
│      .gitignore
│      Hertz_refactored.iml
│      modules.xml
│      workspace.xml
│
├─biz
│  ├─dal
│  │  ├─db
│  │  │  ├─chats
│  │  │  │  │  chat.go
│  │  │  │  │
│  │  │  │  └─monitor
│  │  │  │          chat.go
│  │  │  │          init.go
│  │  │  │          model.go
│  │  │  │
│  │  │  ├─comment
│  │  │  │      comment.go
│  │  │  │
│  │  │  ├─favorite
│  │  │  │      favorite.go
│  │  │  │
│  │  │  ├─publish
│  │  │  │      publish.go
│  │  │  │
│  │  │  ├─relation
│  │  │  │      relation.go
│  │  │  │
│  │  │  ├─user
│  │  │  │      user.go
│  │  │  │
│  │  │  └─video
│  │  │          video.go
│  │  │
│  │  ├─mongodb
│  │  │      init.go
│  │  │
│  │  ├─mysql
│  │  │      init.go
│  │  │
│  │  └─redis
│  │          init.go
│  │
│  ├─handler
│  │  │  ping.go
│  │  │
│  │  ├─chat
│  │  │      chat_service.go
│  │  │
│  │  ├─comment
│  │  │      comment_service.go
│  │  │
│  │  ├─favorite
│  │  │      favorite_service.go
│  │  │
│  │  ├─publish
│  │  │      up_load_video_service.go
│  │  │
│  │  ├─relation
│  │  │      follow_service.go
│  │  │
│  │  ├─user
│  │  │      user_service.go
│  │  │
│  │  └─video
│  │          video_service.go
│  │
│  ├─model
│  │  ├─chat
│  │  │      chat.go
│  │  │
│  │  ├─comment
│  │  │      comment.go
│  │  │
│  │  ├─favorite
│  │  │      favorite.go
│  │  │
│  │  ├─publish
│  │  │      publish.go
│  │  │
│  │  ├─relation
│  │  │      relation.go
│  │  │
│  │  ├─user
│  │  │      user.go
│  │  │
│  │  └─video
│  │          video.go
│  │
│  ├─mv
│  │      jwt.go
│  │
│  ├─pack
│  │      user.go
│  │
│  ├─pkg
│  │  │  code.go
│  │  │  msg.go
│  │  │
│  │  ├─configs
│  │  │  ├─minio
│  │  │  │      config
│  │  │  │
│  │  │  ├─redis
│  │  │  │      redis.conf
│  │  │  │
│  │  │  └─sql
│  │  │          init.sql
│  │  │
│  │  ├─constants
│  │  │      constant.go
│  │  │
│  │  ├─data
│  │  │  └─minio
│  │  │      ├─avatar
│  │  │      │  └─test1.jpg
│  │  │      │      │  xl.meta
│  │  │      │      │
│  │  │      │      └─0df91c0b-64d7-4dae-9c69-e06189c23615
│  │  │      │              part.1
│  │  │      │
│  │  │      └─background
│  │  │          └─test1.png
│  │  │              │  xl.meta
│  │  │              │
│  │  │              └─653fd01e-49bc-462f-866d-f552e51eaf61
│  │  │                      part.1
│  │  │
│  │  ├─errno
│  │  │      errno.go
│  │  │
│  │  └─utils
│  │          crypto.go
│  │          resp.go
│  │          time.go
│  │          video.go
│  │
│  ├─router
│  │  │  register.go
│  │  │
│  │  ├─chat
│  │  │      chat.go
│  │  │      middleware.go
│  │  │
│  │  ├─comment
│  │  │      comment.go
│  │  │      middleware.go
│  │  │      
│  │  ├─favorite
│  │  │      favorite.go
│  │  │      middleware.go
│  │  │
│  │  ├─publish
│  │  │      middleware.go
│  │  │      publish.go
│  │  │
│  │  ├─relation
│  │  │      middleware.go
│  │  │      relation.go
│  │  │
│  │  ├─user
│  │  │      middleware.go
│  │  │      user.go
│  │  │
│  │  └─video
│  │          middleware.go
│  │          video.go
│  │
│  └─utils
│          md5.go
│
└─idl
chat.thrift
comment.thrift
favorite.thrift
publish.thrift
relation.thrift
user.thrift
video.thrift

