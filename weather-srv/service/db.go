package service

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	proto "ms-weather/weather-srv/proto"
	"ms-weather/weather-srv/units"
	"time"
)

func getApiConfig(sourcePath, configName string) *gjson.Json {
	sourceFile := sourcePath + "/" + configName
	sc, err := gjson.Load(sourceFile)
	if err != nil {
		glog.Error("加载配置文件出错！", err)
		return nil
	}

	glog.Info("接口源名称：", sc.GetString("sourceName"))
	sourceApi := sc.GetString("sourceApi")
	if sourceApi == "" {
		return nil
	}
	//配置文件中的参数
	paramData := sc.GetString("param")
	j := gjson.New(paramData)
	paramDataMap := j.ToMap()

	//处理特殊的参数
	if paramDataMap["autoDate"] != "" {
		curTime := time.Now()                                                             // 获取当前时间
		paramDataMap["autoDate"] = curTime.Format(gconv.String(paramDataMap["autoDate"])) // 2020-05-19 10:32:07.185
	}

	//通过文字模板的处理，进行参数替换配置
	ApiStr := units.StringLiteralTemplate(sourceApi, paramDataMap)
	sc.Set("sourceApi", ApiStr)
	return sc
}

// setNowData 从json数据中获取天气情况数据
func setNowData(apiData string, sc *gjson.Json, wd *proto.NowData) {
	apiJson := gjson.New(apiData)
	wd.City = apiJson.GetString(sc.GetString("now.City"))
	wd.CityCode = apiJson.GetString(sc.GetString("now.CityCode"))
	wd.UpdateTime = apiJson.GetString(sc.GetString("now.UpdateTime"))
	wDataInfo := proto.WeatherData{}
	if wd.Data == nil {
		wd.Data = &wDataInfo
	}
	if wd.Data.City == "" {
		wd.Data.City = units.NormFormat(apiJson.GetString(sc.GetString("now.City")), sc.GetString("filter.City"))
	}
	if wd.Data.CityCode == "" {
		wd.Data.CityCode = apiJson.GetString(sc.GetString("now.CityCode"))
	}
	if wd.Data.Tem == "" {
		wd.Data.Tem = units.NormFormat(apiJson.GetString(sc.GetString("now.Tem")), sc.GetString("filter.Tem"))
	}
	if wd.Data.TemHigh == "" {
		wd.Data.TemHigh = units.NormFormat(apiJson.GetString(sc.GetString("now.TemHigh")), sc.GetString("filter.TemHigh"))
	}
	if wd.Data.TemLow == "" {
		wd.Data.TemLow = units.NormFormat(apiJson.GetString(sc.GetString("now.TemLow")), sc.GetString("filter.TemLow"))

	}
	if wd.Data.Aqi == "" {
		wd.Data.Aqi = units.NormFormat(apiJson.GetString(sc.GetString("now.Aqi")), sc.GetString("filter.Aqi"))

	}
	if wd.Data.Pm10 == "" {
		wd.Data.Pm10 = units.NormFormat(apiJson.GetString(sc.GetString("now.Pm10")), sc.GetString("filter.Pm10"))

	}
	if wd.Data.Pm25 == "" {
		wd.Data.Pm25 = units.NormFormat(apiJson.GetString(sc.GetString("now.Pm25")), sc.GetString("filter.Pm25"))

	}
	if wd.Data.Date == "" {
		wd.Data.Date = units.NormFormat(apiJson.GetString(sc.GetString("now.Date")), sc.GetString("filter.Date"))

	}
	if wd.Data.UpdateTime == "" {
		wd.Data.UpdateTime = units.NormFormat(apiJson.GetString(sc.GetString("now.UpdateTime")), sc.GetString("filter.UpdateTime"))

	}
	if wd.Data.CondCodeDay == "" {
		wd.Data.CondCodeDay = units.NormFormat(apiJson.GetString(sc.GetString("now.CondCodeDay")), sc.GetString("filter.CondCodeDay"))

	}
	if wd.Data.CondTxtDay == "" {
		wd.Data.CondTxtDay = units.NormFormat(apiJson.GetString(sc.GetString("now.CondTxtDay")), sc.GetString("filter.CondTxtDay"))

	}
	if wd.Data.CondCodeNight == "" {
		wd.Data.CondCodeNight = units.NormFormat(apiJson.GetString(sc.GetString("now.CondCodeNight")), sc.GetString("filter.CondCodeNight"))

	}
	if wd.Data.CondTxtNight == "" {
		wd.Data.CondTxtNight = units.NormFormat(apiJson.GetString(sc.GetString("now.CondTxtNight")), sc.GetString("filter.CondTxtNight"))

	}
	if wd.Data.WindDeg == "" {
		wd.Data.WindDeg = units.NormFormat(apiJson.GetString(sc.GetString("now.WindDeg")), sc.GetString("filter.WindDeg"))

	}
	if wd.Data.WindDir == "" {
		wd.Data.WindDir = units.NormFormat(apiJson.GetString(sc.GetString("now.WindDir")), sc.GetString("filter.WindDir"))

	}
	if wd.Data.WindSc == "" {
		wd.Data.WindSc = units.NormFormat(apiJson.GetString(sc.GetString("now.WindSc")), sc.GetString("filter.WindSc"))

	}
	if wd.Data.WindSpd == "" {
		wd.Data.WindSpd = units.NormFormat(apiJson.GetString(sc.GetString("now.WindSpd")), sc.GetString("filter.WindSpd"))

	}
	if wd.Data.Pres == "" {
		wd.Data.Pres = units.NormFormat(apiJson.GetString(sc.GetString("now.Pres")), sc.GetString("filter.Pres"))

	}
	if wd.Data.Vis == "" {
		wd.Data.Vis = units.NormFormat(apiJson.GetString(sc.GetString("now.Vis")), sc.GetString("filter.Vis"))

	}
	if wd.Data.Type == "" {
		wd.Data.Type = units.NormFormat(apiJson.GetString(sc.GetString("now.Type")), sc.GetString("filter.Type"))

	}
	if wd.Data.Cold == "" {
		wd.Data.Cold = units.NormFormat(apiJson.GetString(sc.GetString("now.Cold")), sc.GetString("filter.Cold"))

	}
	if wd.Data.Notice == "" {
		wd.Data.Notice = units.NormFormat(apiJson.GetString(sc.GetString("now.Notice")), sc.GetString("filter.Notice"))

	}
	if wd.Data.Source == "" {
		wd.Data.Source = sc.GetString("sourceName")
	} else {
		wd.Data.Source = wd.Data.Source + "|" + sc.GetString("sourceName")

	}

}

//saveNowData 保存天气情况
func saveNowData(wDataInfo *proto.WeatherData) {
	ctime := gtime.Now().Format("U")
	sdata := g.Map{
		"city":          wDataInfo.City,
		"cityCode":      wDataInfo.CityCode,
		"tem":           wDataInfo.Tem,
		"temHigh":       wDataInfo.TemHigh,
		"temLow":        wDataInfo.TemLow,
		"aqi":           wDataInfo.Aqi,
		"pm10":          wDataInfo.Pm10,
		"pm25":          wDataInfo.Pm25,
		"date":          wDataInfo.Date,
		"updateTime":    wDataInfo.UpdateTime,
		"condCodeDay":   wDataInfo.CondCodeDay,
		"condTxtDay":    wDataInfo.CondTxtDay,
		"condCodeNight": wDataInfo.CondCodeNight,
		"condTxtNight":  wDataInfo.CondTxtNight,
		"windDeg":       wDataInfo.WindDeg,
		"windDir":       wDataInfo.WindDir,
		"windSc":        wDataInfo.WindSc,
		"windSpd":       wDataInfo.WindSpd,
		"pres":          wDataInfo.Pres,
		"vis":           wDataInfo.Vis,
		"type":          wDataInfo.Type,
		"cold":          wDataInfo.Cold,
		"notice":        wDataInfo.Notice,
		"source":        wDataInfo.Source,
		"CreateTime":    ctime,
	}
	m := g.DB().Table("now")
	r, err := m.Data(sdata).Insert()
	if err != nil {
		glog.Error(err)
	}
	rid, _ := r.LastInsertId()
	glog.Info("已存入数据库", rid)
}

// setForecastData 解析json中的forecast数据
func setForecastData(apiData string, sc *gjson.Json, fd *proto.ForecastData) {
	forecastDataJson := gjson.New(apiData)
	fd.City = forecastDataJson.GetString(sc.GetString("forecast.City"))
	fd.CityCode = forecastDataJson.GetString(sc.GetString("forecast.CityCode"))
	fd.UpdateTime = forecastDataJson.GetString(sc.GetString("forecast.UpdateTime"))

	dataNode := sc.GetString("forecast.Config.DataNode")
	startRow := sc.GetInt("forecast.Config.StartRow")
	endRow := sc.GetInt("forecast.Config.EndRow")
	dataList := forecastDataJson.GetMaps(dataNode)

	if endRow > len(dataList) {
		endRow = len(dataList)
	}
	if startRow > endRow {
		return
	}

	wDataList := []*proto.WeatherData{}

	if fd.Data == nil {
		fd.Data = wDataList
	}

	for i := startRow; i < endRow; i++ {
		wDataInfo := proto.WeatherData{}
		glog.Info(dataList[i])
		if wDataInfo.TemHigh == "" {

			wDataInfo.TemHigh = units.NormFormat(gconv.String(dataList[i][sc.GetString("forecast.Data.TemHigh")]), sc.GetString("filter.TemHigh"))
		}
		if wDataInfo.TemLow == "" {
			wDataInfo.TemLow = units.NormFormat(gconv.String(dataList[i][sc.GetString("forecast.Data.TemLow")]), sc.GetString("filter.TemLow"))

		}

		if wDataInfo.Date == "" {
			wDataInfo.Date = units.NormFormat(gconv.String(dataList[i][sc.GetString("forecast.Data.Date")]), sc.GetString("filter.Date"))

		}
		if wDataInfo.UpdateTime == "" {
			wDataInfo.UpdateTime = units.NormFormat(gconv.String(dataList[i][sc.GetString("forecast.Data.UpdateTime")]), sc.GetString("filter.UpdateTime"))

		}
		if wDataInfo.CondCodeDay == "" {
			wDataInfo.CondCodeDay = units.NormFormat(gconv.String(dataList[i][sc.GetString("forecast.Data.CondCodeDay")]), sc.GetString("filter.CondCodeDay"))

		}
		if wDataInfo.CondTxtDay == "" {
			wDataInfo.CondTxtDay = units.NormFormat(gconv.String(dataList[i][sc.GetString("forecast.Data.CondTxtDay")]), sc.GetString("filter.CondTxtDay"))

		}
		if wDataInfo.CondCodeNight == "" {
			wDataInfo.CondCodeNight = units.NormFormat(gconv.String(dataList[i][sc.GetString("forecast.Data.CondCodeNight")]), sc.GetString("filter.CondCodeNight"))

		}
		if wDataInfo.CondTxtNight == "" {
			wDataInfo.CondTxtNight = units.NormFormat(gconv.String(dataList[i][sc.GetString("forecast.Data.CondTxtNight")]), sc.GetString("filter.CondTxtNight"))

		}
		if wDataInfo.WindDeg == "" {
			wDataInfo.WindDeg = units.NormFormat(gconv.String(dataList[i][sc.GetString("forecast.Data.WindDeg")]), sc.GetString("filter.WindDeg"))

		}
		if wDataInfo.WindDir == "" {
			wDataInfo.WindDir = units.NormFormat(gconv.String(dataList[i][sc.GetString("forecast.Data.WindDir")]), sc.GetString("filter.WindDir"))

		}
		if wDataInfo.WindSc == "" {
			wDataInfo.WindSc = units.NormFormat(gconv.String(dataList[i][sc.GetString("forecast.Data.WindSc")]), sc.GetString("filter.WindSc"))

		}
		if wDataInfo.WindSpd == "" {
			wDataInfo.WindSpd = units.NormFormat(gconv.String(dataList[i][sc.GetString("forecast.Data.WindSpd")]), sc.GetString("filter.WindSpd"))

		}
		if wDataInfo.Pres == "" {
			wDataInfo.Pres = units.NormFormat(gconv.String(dataList[i][sc.GetString("forecast.Data.Pres")]), sc.GetString("filter.Pres"))

		}

		if wDataInfo.Type == "" {
			wDataInfo.Type = units.NormFormat(gconv.String(dataList[i][sc.GetString("forecast.Data.Type")]), sc.GetString("filter.Type"))

		}
		if wDataInfo.Cold == "" {
			wDataInfo.Cold = units.NormFormat(gconv.String(dataList[i][sc.GetString("forecast.Data.Cold")]), sc.GetString("filter.Cold"))

		}
		if wDataInfo.Notice == "" {
			wDataInfo.Notice = units.NormFormat(gconv.String(dataList[i][sc.GetString("forecast.Data.Notice")]), sc.GetString("filter.Notice"))

		}
		fd.Data = append(fd.Data, &wDataInfo)
	}

}

// setSeasData 解析近海天气情况
func setSeasData(apiData string, sc *gjson.Json, sd *proto.SeasData) {
	seasDataJson := gjson.New(apiData)
	//处理24小时潮汐数据
	sDataInfo := proto.SeasWeatherData{}
	if sd.Data == nil {
		sd.Data = &sDataInfo
	}
	sd.Data.PortName = seasDataJson.GetString(sc.GetString("seas.portName"))
	sd.Data.SeaLevel = seasDataJson.GetString(sc.GetString("seas.seaLevel"))
	sd.Data.Date = seasDataJson.GetString(sc.GetString("seas.date"))

	dataNode := sc.GetString("seas.tide")
	dataList := seasDataJson.GetMap(dataNode)
	sDataList := gconv.Strings(dataList)
	if sd.Data.Tide == nil {
		sd.Data.Tide = sDataList
	}

	//处理range数据
	rDataInfo := []*proto.SeasRangeData{}
	if sd.Data.Range == nil {
		sd.Data.Range = rDataInfo
	}
	rangeNode := sc.GetString("seas.range")
	rangeList := seasDataJson.GetMaps(rangeNode)
	for i := 0; i < len(rangeList); i++ {
		rDataInfo := proto.SeasRangeData{}
		rDataInfo.Type = units.NormFormat(gconv.String(rangeList[i][sc.GetString("seas.rangeData.type")]), sc.GetString("filter.seas.rangeData.type"))
		rDataInfo.Time = units.NormFormat(gconv.String(rangeList[i][sc.GetString("seas.rangeData.time")]), sc.GetString("filter.seas.rangeData.time"))
		rDataInfo.Height = units.NormFormat(gconv.String(rangeList[i][sc.GetString("seas.rangeData.height")]), sc.GetString("filter.seas.rangeData.height"))
		sd.Data.Range = append(sd.Data.Range, &rDataInfo)
	}

}
