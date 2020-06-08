package handler

import (
	"context"
	proto "ms-weather/weather-srv/proto"
	"ms-weather/weather-srv/service"
)

type Weather struct{}

func (w *Weather) Now(ctx context.Context, req *proto.DataReq, rsp *proto.NowData) error {
	nd := service.NewSetupData()
	nd.NowData(req, rsp)
	return nil
}

func (w *Weather) Forecast(ctx context.Context, req *proto.DataReq, rsp *proto.ForecastData) error {
	fd := service.NewSetupData()
	fd.ForecastData(req, rsp)
	return nil
}

func (w *Weather) Seas(ctx context.Context, req *proto.DataReq, rsp *proto.SeasData) error {
	//设定相关参数
	commonsOpts := []service.Option{
		service.GetDataType(service.DATA_SEAS),
	}
	fd := service.NewSetupData(commonsOpts...)
	fd.SeasData(req, rsp)
	return nil
}
