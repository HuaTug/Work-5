
## 完成情况

完整项目目录: /Tree.md

`docker`
```

```

`主要业务逻辑`
```

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
│  │       constant.go       
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

```
`3.14`
```
完成了将原来的Gin框架并且使用自动迁移模式进行数据表的创建 改为使用Hertz框架进行重构，
并且在此过程中引入了Hertz认证的JWT进行Token鉴权验证
```

`3.18`
```
学着使用自己构建SQL脚本去构建数据表，并且在对数据进行修改的过程中引入了触发器的概念，
对点赞关注等操作进行一个自动关联，同时也是有了外键关联，建立起数据表之间的联系
```

`3.20`
```
遇到了websocket的无法连接问题
----------------------------
下午解决了 在项目中看注释
----------------------------
3.20号晚
完成了websocket两人聊天的demo
```

### 需要优化的地方：

`Redis的引入`
```

```
`RabiitMq 消息队列的引入`
```

```
## Bonus:



## 下一阶段

