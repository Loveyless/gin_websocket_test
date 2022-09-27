package service

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

type UserRoom struct {
	UserIdentity string `json:"user_identity" bson:"user_identity"`
	RoomIdentity string `json:"room_identity" bson:"room_identity"`
	// MessageIdentity string `json:"message_identity" bson:"message_identity"`
	CratedAt  int64 `json:"crated_at" bson:"crated_at"`
	UpdatedAt int64 `json:"updated_at" bson:"updated_t"`
}

func (UserRoom) CollectionName() string {
	return "user_room"
}

//获取用户房间 通过用户id和房间id
func GetUserRoomByUserIdentityRoomIdentity(userIdentity, roomIdentity string) (*UserRoom, error) {
	ur := new(UserRoom)
	err := Mongo.Collection(UserRoom{}.CollectionName()).
		FindOne(context.Background(), bson.D{{Key: "user_identity", Value: userIdentity}, {Key: "room_identity", Value: roomIdentity}}).
		Decode(ur)
	return ur, err
}

//查询房间的所有用户 通过房间identity
func GetUserRoomByRoomIdentity(roomIdentity string) ([]*UserRoom, error) {

	cur, err := Mongo.Collection(UserRoom{}.CollectionName()).
		Find(context.Background(), bson.D{{Key: "room_identity", Value: roomIdentity}})

	//放入切片
	userRoomList := make([]*UserRoom, 0)
	for cur.Next(context.Background()) {
		item := new(UserRoom)
		err := cur.Decode(item)
		if err != nil {
			log.Println(err.Error())
		}
		fmt.Println(item)
		userRoomList = append(userRoomList, item)
	}
	fmt.Println(err)
	return userRoomList, err

}
