package controllers

import (
  "fmt"
  "time"
  "github.com/revel/revel"
	"github.com/revel/revel/cache"
  "encoding/json"

  "github.com/go-xorm/xorm"
  "gopkg.in/redis.v5"
  "github.com/website-m-backend/libs"
  "github.com/website-m-backend/app/models"
)

var (
  db *xorm.Engine
  rs *redis.Client
)

type Api struct {
	*revel.Controller
}

type results struct {
  Code string `json:"code"`
  Data []models.Regions `json:"data"`
}

func (c Api) Index() revel.Result {
  greeting := "Hello World"
  return c.Render(greeting)
}

func (c Api) JsonData() revel.Result {
    var s results
    hotCities := []int64 {121, 3, 267, 2316, 1337, 852}
    data := make(map[string]interface{})
    regions := make(map[string][]models.Regions)

    ttl :=  c.Params.Get("ttl")

    err := cache.Get("cityData", &regions)
    // fmt.Println(ttl)
    if err == nil && ttl != "0" && ttl == "" {
      data["code"] = "00000"
      data["msg"] = "success"
      data["result"] = regions
      return c.RenderJson(data)
    }


    body, err := libs.RequestGet("webdata", "/esf/web/getCityList")
    if err != nil {
      fmt.Println(err)
      data["code"] = "99999"
      data["msg"] = "server error"
      return c.RenderJson(data)
    }

    json.Unmarshal([]byte(body), &s)

    for _, v := range s.Data {
      for _, h := range hotCities {
        if h == v.Id {
          regions["热门城市"] = append(regions["热门城市"], v)
        }
      }
      regions[v.Scope] = append(regions[v.Scope], v)
    }

    // regionsString, _ := json.Marshal(regions)
    cache.Set("cityData", regions, 86400 * time.Second)

    data["code"] = "00000"
    data["msg"] = "success"
    data["result"] = regions
    return c.RenderJson(data)
}
