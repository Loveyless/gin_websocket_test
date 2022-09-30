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
	RoomType  int   `json:"room_type" bson:"room_type"`
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

//判断俩人是否具有单聊房间 (是不是好友)
func JudgeUserIsFriend(userIdentity1, userIdentity2 string) (bool, error) {
	// i, err := Mongo.Collection(UserRoom{}.CollectionName()).
	// 	CountDocuments(context.Background(), bson.D{{Key: "user_identity", Value: userIdentity1}, {Key: "room_type", Value: 1}, {Key: "room_identity", Value: userIdentity2}})

	//查出id1的所有单聊房间数据
	fmt.Println("userIdentity1", userIdentity1, "userIdentity2", userIdentity2)
	cur, err := Mongo.Collection(UserRoom{}.CollectionName()).Find(context.Background(), bson.D{{Key: "user_identity", Value: userIdentity1}, {Key: "room_type", Value: 1}})
	if err != nil {
		log.Println(userIdentity1, "查找id数据出错", err.Error())
	}
	//将id1的所有房间id存入切片
	curList := make([]string, 0)
	for cur.Next(context.Background()) {
		item := new(UserRoom)
		err := cur.Decode(item)
		if err != nil {
			log.Println(err.Error())
		}
		curList = append(curList, item.RoomIdentity)
	}
	fmt.Println(curList)

	//获取关联id2的单聊房间个数 获取一下id2的所有单聊房间有没有id1的房间
	i, err := Mongo.Collection(UserRoom{}.CollectionName()).CountDocuments(context.Background(), bson.M{"user_identity": userIdentity2, "room_type": 1, "room_identity": bson.M{"$in": curList}})
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(i)

	return i > 0, err
}
