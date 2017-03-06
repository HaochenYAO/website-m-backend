package controllers

import (
	"github.com/go-xorm/xorm"
	"github.com/revel/revel"
	redis "gopkg.in/redis.v5"

	"github.com/website-m-backend/libs"
)

const (
	// SUCCESS 成功msg
	SUCCESS = "success"
	// CSUCCESS 成功code
	CSUCCESS = "00000"
)

var (
	db *xorm.Engine
	rs *redis.Client
	rq libs.HTTPHandle
)

func init() {
	revel.OnAppStart(Init)
}

// Init 初始化
func Init() {
	db = libs.DbEngine
	rs = libs.RedisClient
	rq = libs.CreateHTTPClient()
}
