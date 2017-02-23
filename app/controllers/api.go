package controllers

import (
  "fmt"
  "github.com/revel/revel"
  "encoding/json"

  "github.com/go-xorm/xorm"
  "github.com/fdd/website-m-backend/libs"
)

var (
	engine *xorm.Engine
)

type Api struct {
	*revel.Controller
}

type city struct {
  Id int64 `json:"id"`
  Name string `json:"name"`
  Maplat string `json:"maplat"`
  Maplng string `json:"maplng"`
  Scope string `json:"scope"`
  Sort int64 `json:"sort"`
}

type dataRaw struct {
  Code string `json:"code"`
  Data []city `json:"data"`
}

func (c Api) Index() revel.Result {
  greeting := "Aloha World"
  return c.Render(greeting)
}

func (c Api) JsonData() revel.Result {
    var s dataRaw
    hotCities := []int64 {121, 3, 267, 2316, 1337, 852}
    data := make(map[string]interface{})
    regions := make(map[string][]city)


    body, err := libs.RequestGet("webdata", "/esf/web/getCityList")
    if err != nil {
      fmt.Println(err)
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


    data["code"] = "00000"
    data["msg"] = "success"
    data["result"] = regions

    return c.RenderJson(data)
}
