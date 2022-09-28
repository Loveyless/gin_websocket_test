package service

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

func GetMessageListByRoomIdentity(roomIdentity string, limit int64, offset int64) ([]*MessageBasic, error) {

	messageList := make([]*MessageBasic, 0)

	cur, err := Mongo.Collection(MessageBasic{}.CollectionName()).
		Find(context.Background(), bson.D{{Key: "room_identity", Value: roomIdentity}}, &options.FindOptions{
			Limit: &limit,
			Skip:  &offset,
			Sort:  bson.D{{Key: "crated_at", Value: -1}}, //倒序 降序
		})
	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		item := new(MessageBasic)
		err := cur.Decode(item)
		if err != nil {
			return nil, err
		}
		messageList = append(messageList, item)
	}

	return messageList, err
}
