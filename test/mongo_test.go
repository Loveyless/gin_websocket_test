package test

import (
	"context"
	"fmt"
	"gin_websocket_test/service"
	"log"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//mongo读多条数据
func TestFind(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://go_test:test@101.200.243.101:27017/go_test"))

	db := client.Database("go_test")

	cur, _ := db.Collection("user_basic").Find(context.TODO(), bson.D{})
	for cur.Next(context.TODO()) {
		list := new(service.UserBasic)
		err := cur.Decode(list)
		if err != nil {
			log.Println(err.Error())
		}
		fmt.Println(list)
	}

}
