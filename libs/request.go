package libs

import (
  "fmt"
  "net/http"
  "strings"
  "time"
  "io/ioutil"

  "github.com/revel/revel"
	"github.com/robfig/config"
)

var httpClient http.Client

func InitClient() {
  transport := &http.Transport{
    MaxIdleConns: 10,
    IdleConnTimeout: 30 * time.Second,
    DisableCompression: true,
  }

  httpClient = http.Client{
    Timeout: time.Second * 2,
    Transport: transport,
  }
}


func RequestGet(api string, route string, params map[string]string) ([]byte, error) {
  c, _ := config.ReadDefault(revel.BasePath + "/conf/api.conf")
  host, _ := c.String(revel.RunMode, fmt.Sprintf("%s.host", api))
  port, _ := c.String(revel.RunMode, fmt.Sprintf("%s.port", api))
  query := []string {}
  var queryString string = ""

  for k, v := range params {
    query = append(query, k + "=" + v)
  }

  if len(query) > 0 {
    queryString = "?" + strings.Join(query, "&")
  }

  url := fmt.Sprintf("http://%s:%s%s%s", host, port, route, queryString)
  resp, err := httpClient.Get(url)

  if err != nil {
    return []byte{}, err
  }
  body, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    return []byte{}, err
  }

  defer resp.Body.Close()
  return body, err
}
