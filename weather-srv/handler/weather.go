package handler

import (
	"context"
	"github.com/gogf/gf/os/glog"
	"github.com/micro/go-micro/v2/errors"
	proto "ms-weather/weather-srv/proto"
	"ms-weather/weather-srv/service"
)

type Weather struct{}

func (w *Weather) Now(ctx context.Context, req *proto.DataReq, rsp *proto.NowData) error {
	sd := new(service.SetupData)
	sd.ApiData(req, rsp)
	return nil
}

func (w *Weather) Forecast(ctx context.Context, req *proto.DataReq, rsp *proto.ForecastData) error {
	glog.Info("收到获取天气预报的请求请求")
	if len(req.CityCode) == 0 {
		return errors.BadRequest("go.micro.weather", "no content")
	}
	rsp.City = "Hello " + req.GetCity()
	return nil
}
