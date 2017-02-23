package libs

import (
  "fmt"
  "net/http"
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


func RequestGet(api string, route string) ([]byte, error) {
  c, _ := config.ReadDefault(revel.BasePath + "/conf/api.conf")
  host, _ := c.String(revel.RunMode, fmt.Sprintf("%s.host", api))
  port, _ := c.String(revel.RunMode, fmt.Sprintf("%s.port", api))

  url := fmt.Sprintf("http://%s:%s%s", host, port, route)
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
