package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func initDb() {
	// 客户端设置
	co := options.Client().ApplyURI("mongo://localhost:27017")

	// 连接MongoDB
	var err error
	client, err = mongo.Connect(context.TODO(), co)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err2 := client.Ping(context.TODO(), nil)
	if err2 != nil {
		log.Fatal(err2)
	} else {
		fmt.Println("连接成功")
	}

}

func main() {
	initDb()
}
