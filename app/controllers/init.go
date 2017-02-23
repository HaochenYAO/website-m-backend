package controllers

import (
  "github.com/revel/revel"

  "github.com/website-m-backend/libs"
)


func init() {
	revel.OnAppStart(Init)
}

func Init() {
	db = libs.DbEngine
  rs = libs.RedisClient
}
