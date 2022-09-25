package service

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mongo = InitMongo()

func InitMongo() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://go_test:test@101.200.243.101:27017/go_test"))
	if err != nil {
		panic("连接mongo数据库失败" + err.Error())
	}
	//检查连接
	err2 := client.Ping(context.TODO(), nil)
	if err2 != nil {
		panic("连接mongo数据库失败" + err2.Error())
	} else {
		fmt.Println("mongo连接成功")
	}

	return client.Database("go_test")
}
