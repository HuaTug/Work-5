namespace go chat

struct BaseResp{
    1:i64 code,
    2:string msg,
}
struct Message{
    1:i64 MessageID
    2:i64 SenderID
    3:i64 ReceiverID
    4:string MessageText
    5:string SendTime
    6:i64 State
}
struct HistoryMsg{
    1:i64 From
    2:string content

}
struct SendMsg{
    1:string Conetent
    2:i64 id
    3:i64 ToUid
    4:string sendTime
}
struct ReplyMsg{
    1:string From
    2:i64 Code
    3:string Content
}
struct MessageChatRequest{
    1:i64 to_id
}

struct MessageChatResp{
    1:BaseResp base
}

service ChatService{
    MessageChatResp Chat (1:MessageChatRequest req)(api.get="/chats/ws")
}