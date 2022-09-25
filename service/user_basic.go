package service

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserBasic struct {
	Id        string `json:"_id" bson:"_id"`
	Username  string `json:"username" bson:"username"`
	Password  string `json:"password" bson:"password"`
	Nicname   string `json:"nicname" bson:"nicname"`
	Sex       int    `json:"sex" bson:"sex"`
	Email     string `json:"email" bson:"email"`
	Avatar    string `json:"avatar" bson:"avatar"`
	CratedAt  int64  `json:"crated_at" bson:"crated_at"`
	UpdatedAt int64  `json:"updated_at" bson:"updated_t"`
}

func (UserBasic) CollectionName() string {
	return "user_basic"
}

//查询用户
func GetUserBasicByUsernamePassword(username, password string) (*UserBasic, error) {
	ub := new(UserBasic)
	err := Mongo.Collection(UserBasic{}.CollectionName()).
		FindOne(context.TODO(),
			bson.D{{Key: "username", Value: username}, {Key: "password", Value: password}}).Decode(ub)
	return ub, err
}

//用户详情
func GetUserBasicByIdentity(identity primitive.ObjectID) (*UserBasic, error) {
	ub := new(UserBasic)
	err := Mongo.Collection(UserBasic{}.CollectionName()).
		FindOne(context.TODO(),
			bson.D{{Key: "_id", Value: identity}}).Decode(ub)
	return ub, err
}
