package libs

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/revel/revel"
	"github.com/robfig/config"
)

// HTTPClient client结构体
type HTTPClient struct {
	client http.Client
}

// HTTPHandle http请求处理
type HTTPHandle interface {
	Get(string, string, map[string]string) ([]byte, error)
}

// CreateHTTPClient 创建client对象
func CreateHTTPClient() *HTTPClient {
	transport := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	httpClient := http.Client{
		Timeout:   time.Second * 2,
		Transport: transport,
	}
	return &HTTPClient{httpClient}
}

func (h *HTTPClient) Get(api string, route string, params map[string]string) ([]byte, error) {
	c, _ := config.ReadDefault(revel.BasePath + "/conf/api.conf")
	host, _ := c.String(revel.RunMode, fmt.Sprintf("%s.host", api))
	port, _ := c.String(revel.RunMode, fmt.Sprintf("%s.port", api))
	query := []string{}
	var queryString string

	for k, v := range params {
		query = append(query, k+"="+v)
	}

	if len(query) > 0 {
		queryString = "?" + strings.Join(query, "&")
	}

	url := fmt.Sprintf("http://%s:%s%s%s", host, port, route, queryString)
	resp, err := h.client.Get(url)

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
