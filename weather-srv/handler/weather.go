package handler

import (
	"context"
	proto "ms-weather/weather-srv/proto"
	"ms-weather/weather-srv/service"
)

type Weather struct{}

func (w *Weather) Now(ctx context.Context, req *proto.DataReq, rspN *proto.NowData) error {
	nd := service.NewSetupData()
	nd.NowData(req, rspN)
	return nil
}

func (w *Weather) Forecast(ctx context.Context, req *proto.DataReq, rspF *proto.ForecastData) error {
	fd := service.NewSetupData()
	fd.ForecastData(req, rspF)
	return nil
}
