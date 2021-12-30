package models

import (
	"context"
	"go-gin-seed/pkg/logger"
	"go-gin-seed/pkg/setting"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

//表基类
type Model struct {
	ID        int `gorm:"primary_key" json:"id"`
	GmtCreate int `json:"gmt_create"`
	GmtUpdate int `json:"gmt_update"`
}

func Init() {
	ctx, cancel := context.WithTimeout(context.Background(), setting.DatabaseSetting.Timeout)
	defer cancel()
	o := options.Client().ApplyURI(setting.DatabaseSetting.Host)
	o.SetMaxPoolSize(setting.DatabaseSetting.MaxPoolSize)
	client, err := mongo.Connect(ctx, o)
	if err != nil {
		logger.Fatal("connect to db error:%s", err)
		return
	}
	db = client.Database(setting.DatabaseSetting.Name)
}
