package controllers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/revel/revel"
	"github.com/revel/revel/cache"

	"github.com/website-m-backend/app/models"
)

// API basic controller
type API struct {
	*revel.Controller
}

type regionsData struct {
	Type string           `json:"type"`
	Data []models.Regions `json:"data"`
}

// Index index page
func (c API) Index() revel.Result {
	data := make(map[string]interface{})
	data["code"] = "1024"
	data["msg"] = "please, don't..."
	return c.RenderJson(data)
}

// CityData 城市数据
func (c API) CityData() revel.Result {
	var r []models.Regions
	hotCities := []int64{121, 3, 267, 2316, 1337, 852}
	data := make(map[string]interface{})
	tmp := make(map[string][]models.Regions)
	query := make(map[string]string)
	regions := []regionsData{}

	query["is_esf"] = "0"
	ttl := c.Params.Get("ttl")

	errCache := cache.Get("cityData", &regions)

	if errCache == nil && ttl != "0" && ttl == "" {
		data["code"] = SUCCESS
		data["msg"] = CSUCCESS
		data["result"] = regions
		return c.RenderJson(data)
	}

	errDb := db.Where("pid = ?", 0).Find(&r)
	if errDb != nil {
		data["code"] = "99999"
		data["msg"] = "error"
		return c.RenderJson(data)
	}

	for _, v := range r {
		for _, h := range hotCities {
			if h == v.Id {
				tmp["热门城市"] = append(tmp["热门城市"], v)
			}
		}
		tmp[v.Scope] = append(tmp[v.Scope], v)
	}

	for k, v := range tmp {
		regions = append(regions, regionsData{k, v})
	}

	cache.Set("cityData", regions, 86400*time.Second)

	data["code"] = SUCCESS
	data["msg"] = CSUCCESS
	data["result"] = regions
	return c.RenderJson(data)
}

// NhData 城市数据
func (c API) NhData() revel.Result {
	data := make(map[string]interface{})
	var query = make(map[string]string)
	query["agentId"] = "3010115"

	dataRaw, _ := rq.Get("webdata", "/nh/agent/basicinfo", query)
	nhData, _ := json.Unmarshal(dataRaw)

	fmt.Println(nhData)

	data["code"] = SUCCESS
	data["msg"] = CSUCCESS
	data["result"] = nhData
	return c.RenderJson(data)
}
