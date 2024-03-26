├─biz
│  ├─config
│  │      config.go
│  │      config.yaml
│  │      mysql.sql
│  │      types.go
│  │
│  ├─dal
│  │  ├─cache
│  │  │      comment.go
│  │  │      init.go
│  │  │      user.go
│  │  │      video.go
│  │  │
│  │  └─db
│  │      │  chat.go
│  │      │  comment.go
│  │      │  favorite.go
│  │      │  init.go
│  │      │  relaiton.go
│  │      │  user.go
│  │      │  video.go
│  │      │
│  │      └─mq
│  │          │  consumer.go
│  │          │  producer.go
│  │          │
│  │          ├─script
│  │          │      task.go
│  │          │
│  │          └─task
│  │                  task_sync.go
│  │
│  ├─handler
│  │  │  ping.go
│  │  │  
│  │  ├─chat
│  │  │      chat_handler.go
│  │  │
│  │  ├─comment
│  │  │      comment_handler.go
│  │  │
│  │  ├─favorite
│  │  │      favorite_handler.go
│  │  │
│  │  ├─publish
│  │  │      up_load_video_handler.go
│  │  │
│  │  ├─relation
│  │  │      follow_handler.go
│  │  │
│  │  ├─user
│  │  │      user_handler.go
│  │  │
│  │  └─video
│  │          video_handler.go
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
│  │      sequence.go
│  │      user.go
│  │
│  ├─pkg
│  │  │  code.go
│  │  │  msg.go
│  │  │
│  │  ├─errno
│  │  │      errno.go
│  │  │
│  │  ├─logging
│  │  │      file.go
│  │  │      log.go
│  │  │
│  │  └─util
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
│  ├─service
│  │  ├─chats
│  │  │  │  chat.go
│  │  │  │
│  │  │  └─im
│  │  │          chat.go
│  │  │          init.go
│  │  │          model.go
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
│  └─utils
│          dsn.go
│          md5.go
│          transfer.go
│
├─Err
│      2024-03-25.error
│      2024-03-26.error
│      2024-03-27.error
│
└─idl
chat.thrift
comment.thrift
favorite.thrift
publish.thrift
relation.thrift
user.thrift
video.thrift
