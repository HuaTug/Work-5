package db

import (
	"Hertz_refactored/biz/model/chat"
)

func CreateMessage(message *chat.Message) error {
	return Db.Model(&chat.Message{}).Create(message).Error
}

func GetMessage(id, toId int64) ([]*chat.Message, error) {
	var msgs []*chat.Message
	if err := Db.Model(&chat.Message{}).Where("receiver_id=? And sender_id=? And state=?", id, toId, 0).
		Find(&msgs).Error; err != nil {
		return msgs, err
	}
	return msgs, nil
}
func GetAllMessage() ([]*chat.Message, error) {
	var msgs []*chat.Message
	if err := Db.Model(&chat.Message{}).Find(&msgs).Error; err != nil {
		return msgs, err
	}
	return msgs, nil
}
