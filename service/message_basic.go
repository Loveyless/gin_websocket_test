package service

import "context"

type MessageBasic struct {
	UserIdentity string `json:"user_identity" bson:"user_identity"`
	RoomIdentity string `json:"room_identity" bson:"room_identity"`
	Data         string `json:"data" bson:"data"`
	CratedAt     int64  `json:"crated_at" bson:"crated_at"`
	UpdatedAt    int64  `json:"updated_at" bson:"updated_t"`
}

func (MessageBasic) CollectionName() string {
	return "message_basic"
}

//添加一条聊天记录
func InsertOneMessageBasic(messageBasic *MessageBasic) error {
	_, err := Mongo.Collection(MessageBasic{}.CollectionName()).
		InsertOne(context.Background(), messageBasic)
	return err
}
