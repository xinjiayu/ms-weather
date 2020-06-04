package service

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	proto "ms-weather/weather-srv/proto"
	"ms-weather/weather-srv/units"
)

func getApiConfig(configName string) *gjson.Json {
	//加载接口源的配置信息文件
	sourcePath := g.Config().GetString("system.sourceConfigPath")
	sourceFile := sourcePath + "/" + configName
	sc, err := gjson.Load(sourceFile)
	if err != nil {
		glog.Error(err)
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
	//通过文字模板的处理，进行参数替换配置
	ApiStr := units.StringLiteralTemplate(sourceApi, paramDataMap)
	sc.Set("sourceApi", ApiStr)
	return sc
}

func setData(apiData string, sc *gjson.Json, wd *proto.NowData) {
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
		wd.Data.WindDeg = apiJson.GetString(sc.GetString("now.WindDeg"))
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
		wd.Data.Source = units.NormFormat(apiJson.GetString(sc.GetString("now.Source")), sc.GetString("filter.Source"))
	} else {
		wd.Data.Source = wd.Data.Source + "|" + units.NormFormat(apiJson.GetString(sc.GetString("now.Source")), sc.GetString("filter.Source"))

	}

}

func saveData(wDataInfo *proto.WeatherData) {
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
