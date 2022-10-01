package service

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type RoomBasic struct {
	Identity     string `json:"identity" bson:"identity"`           //房间唯一表示
	Number       string `json:"number" bson:"number"`               //房间号
	Name         string `json:"name" bson:"name"`                   //房间名
	Info         string `json:"info" bson:"info"`                   //房间信息
	UserIdentity string `json:"user_identity" bson:"user_identity"` //房主id
	CratedAt     int64  `json:"crated_at" bson:"crated_at"`
	UpdatedAt    int64  `json:"updated_at" bson:"updated_t"`
}

func (RoomBasic) CollectionName() string {
	return "room_basic"
}

//添加一条room_basic数据
func InsertOneRoomBasic(rb *RoomBasic) error {
	_, err := Mongo.Collection(RoomBasic{}.CollectionName()).InsertOne(context.Background(), rb)
	return err
}

//删除room_basic数据
func DeleteRoomBasicByRoomIdentity(userRoomIdentity string) error {
	_, err := Mongo.Collection(RoomBasic{}.CollectionName()).
		DeleteOne(context.Background(), bson.D{{Key: "identity", Value: userRoomIdentity}})
	return err
}
