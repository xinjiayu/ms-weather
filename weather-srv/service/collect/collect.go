package collect

import (
	"github.com/gogf/gf/net/ghttp"
)

//GetAPIDataBody 通过http api获取json格式的数据
func GetAPIDataBody(apiURL string) (string, error) {
	c := ghttp.NewClient()
	c.SetHeader("Accept", "application/json")
	if res, e := c.Get(apiURL); e != nil {
		return "", e
	} else {
		defer res.Close()
		body := res.ReadAllString()
		return body, nil
	}
}
