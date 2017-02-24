package tests

import (
  "fmt"
  "encoding/json"
  "github.com/revel/revel/testing"
)

type ApiTest struct {
	testing.TestSuite
}

func (t *ApiTest) Before() {
	println("Set up")
}

type regionsData struct {
  Code string `json:"code"`
  Msg string `json:"msg"`
}

func (t *ApiTest) TestApiGetCities() {
  var s regionsData

  t.Get("/Api/JsonData")

  json.Unmarshal([]byte(t.ResponseBody), &s)
  t.AssertEqual("00000", s.Code)
  fmt.Println(s.Code)
	t.AssertOk()
	// t.AssertContentType("text/html; charset=utf-8")
}

func (t *ApiTest) After() {
	println("Tear down")
}
