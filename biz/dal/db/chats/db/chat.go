package db

import (
	"Hertz_refactored/biz/dal/mysql"
	"Hertz_refactored/biz/model/chat"
)

func CreateMessage(message *chat.Message) error {
	return mysql.Db.Model(&chat.Message{}).Create(message).Error
}

func GetMessage(id, to_id int64) ([]*chat.Message, error) {
	var msgs []*chat.Message
	if err := mysql.Db.Model(&chat.Message{}).Where("receiver_id=? And sender_id=? And state=?", id, to_id, 0).Find(&msgs).Error; err != nil {
		return msgs, err
	}
	return msgs, nil
}
func GetAllMessage() ([]*chat.Message, error) {
	var msgs []*chat.Message
	if err := mysql.Db.Model(&chat.Message{}).Find(&msgs).Error; err != nil {
		return msgs, err
	}
	return msgs, nil
}
