package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/micro/go-micro/v2"
	_ "github.com/micro/go-plugins/registry/consul/v2"
	_ "github.com/micro/go-plugins/registry/eureka/v2"
	"ms-weather/weather-srv/handler"
	proto "ms-weather/weather-srv/proto"
)

func main() {

	logPath := g.Config().GetString("logger.path")
	glog.SetPath(logPath)
	// 创建服务
	service := micro.NewService(
		micro.Name("go.micro.srv.weather"),
		micro.Version("latest"),
	)

	// 初始化服务
	service.Init()

	// 注册Handler
	proto.RegisterWeatherServiceHandler(service.Server(), new(handler.Weather))

	// 运行服务
	if err := service.Run(); err != nil {
		glog.Fatal(err)
	}

}
