package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/web"
	_ "github.com/micro/go-plugins/registry/consul/v2"
	_ "github.com/micro/go-plugins/registry/eureka/v2"
	weather "ms-weather/weather-srv/proto"
	"ms-weather/weather-srv/units"
)

var (
	cl weather.WeatherService
)

//定义api控制器
type Controller struct{}

//Now 获取天气实况信息的api接口
func (c *Controller) Now(r *ghttp.Request) {
	cityCode := r.GetString("cityCode")
	city := r.GetString("city")
	dataReq := weather.DataReq{AppSecret: "", CityCode: cityCode, City: city}
	nowData, err := cl.Now(r.Context(), &dataReq)

	if err != nil {
		units.Json(r, 1, err.Error(), "")
	}

	units.Json(r, 0, "", nowData.Data)
}

//Forecast 获取天气预报信息的api接口
func (c *Controller) Forecast(r *ghttp.Request) {
	cityCode := r.GetString("cityCode")
	city := r.GetString("city")
	dataReq := weather.DataReq{AppSecret: "", CityCode: cityCode, City: city}
	forecastData, err := cl.Forecast(r.Context(), &dataReq)

	if err != nil {
		units.Json(r, 1, err.Error(), "")
	}

	units.Json(r, 0, "", forecastData.Data)
}

func main() {
	// 创建服务
	service := web.NewService(
		web.Name("go.micro.api.weather"),
	)
	service.Init()
	// 创建客户端链接应用服务
	cl = weather.NewWeatherService("go.micro.srv.weather", client.DefaultClient)

	// 创建Http API服务
	s := g.Server()
	c := new(Controller)
	s.BindObject("/weather", c)

	// 注册 Handler
	service.Handle("/", s)
	// 运行微服务
	if err := service.Run(); err != nil {
		glog.Fatal(err)
	}
}
