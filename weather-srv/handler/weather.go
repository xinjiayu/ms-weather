package handler

import (
	"context"
	"github.com/golang/glog"
	proto "ms-weather/weather-srv/proto"
	"ms-weather/weather-srv/service"
)

type Weather struct{}

func (w *Weather) Now(ctx context.Context, req *proto.DataReq, rsp *proto.NowData) error {
	commonsOpts := []service.Option{
		service.SetDataReq(req),
	}
	nd := service.NewSetupData(commonsOpts...)
	nd.NowData(req, rsp)
	glog.Info(rsp)
	return nil
}

func (w *Weather) Forecast(ctx context.Context, req *proto.DataReq, rsp *proto.ForecastData) error {
	//设定相关参数
	commonsOpts := []service.Option{
		service.SetDataReq(req),
	}
	fd := service.NewSetupData(commonsOpts...)
	fd.ForecastData(req, rsp)
	glog.Info(rsp)
	return nil
}

func (w *Weather) Seas(ctx context.Context, req *proto.DataReq, rsp *proto.SeasData) error {
	//设定相关参数
	commonsOpts := []service.Option{
		service.SetDataType(service.DATA_SEAS),
		service.SetDataReq(req),
	}
	fd := service.NewSetupData(commonsOpts...)
	fd.SeasData(req, rsp)
	glog.Info(rsp)
	return nil
}
