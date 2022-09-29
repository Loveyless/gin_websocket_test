package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis/v9"
)

var ctx = context.Background()

func TestGet(t *testing.T) {
	//redis句柄
	rdb := redis.NewClient(&redis.Options{
		Addr:     "101.200.243.101:6379",
		Password: "at123123", // no password set
		DB:       10,         // use default DB
	})

	err := rdb.Set(ctx, "key", "value", time.Second*30).Err()
	if err != nil {
		panic(err)
	}
}

func TestSet(t *testing.T) {
	//redis句柄
	rdb := redis.NewClient(&redis.Options{
		Addr:     "101.200.243.101:6379",
		Password: "at123123", // no password set
		DB:       10,         // use default DB
	})
	str, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(str)
}
