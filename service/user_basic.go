package service

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type UserBasic struct {
	Identity  string `json:"identity" bson:"identity"`
	Username  string `json:"username" bson:"username"`
	Password  string `json:"password" bson:"password"`
	Nicname   string `json:"nicname" bson:"nicname"`
	Sex       int    `json:"sex" bson:"sex"`
	Email     string `json:"email" bson:"email"`
	Avatar    string `json:"avatar" bson:"avatar"`
	CreatedAt int64  `json:"crated_at" bson:"created_at"`
	UpdatedAt int64  `json:"updated_at" bson:"updated_at"`
}

func (UserBasic) CollectionName() string {
	return "user_basic"
}

//验证查询用户通过账号密码
func GetUserBasicByUsernamePassword(username, password string) (*UserBasic, error) {
	ub := new(UserBasic)
	err := Mongo.Collection(UserBasic{}.CollectionName()).
		FindOne(context.TODO(),
			bson.D{{Key: "username", Value: username}, {Key: "password", Value: password}}).Decode(ub)
	return ub, err
}

//用户详情通过identity
func GetUserBasicByIdentity(identity string) (*UserBasic, error) {
	ub := new(UserBasic)
	err := Mongo.Collection(UserBasic{}.CollectionName()).
		FindOne(context.TODO(),
			bson.D{{Key: "identity", Value: identity}}).Decode(ub)
	return ub, err
}

//用户详情通过username (非本人查询)
func GetUserBasicByUsernameInfo(username string) (*UserBasic, error) {
	ub := new(UserBasic)
	err := Mongo.Collection(UserBasic{}.CollectionName()).
		FindOne(context.TODO(),
			bson.D{{Key: "username", Value: username}}).Decode(ub)
	return ub, err
}

//查询用户通过邮箱 (用于校验数据库是否有这个邮箱)
func GetUserBasicByEmail(emailString string) bool {
	sum, err := Mongo.Collection(UserBasic{}.CollectionName()).
		CountDocuments(context.TODO(), bson.D{{Key: "email", Value: emailString}})
	return (err == nil && sum > 0)
}

//查询用户通过用户名 (用于校验数据库是否有这个用户名)
func GetUserBasicByUsername(username string) bool {
	sum, err := Mongo.Collection(UserBasic{}.CollectionName()).
		CountDocuments(context.TODO(), bson.D{{Key: "username", Value: username}})
	return (err == nil && sum > 0)
}

//添加用户
func InsertUserBasic(userBasic *UserBasic) error {
	_, err := Mongo.Collection(UserBasic{}.CollectionName()).InsertOne(context.TODO(), userBasic)
	return err
}
