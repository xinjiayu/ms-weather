package service

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	_ "github.com/mattn/go-sqlite3"
	proto "ms-weather/weather-srv/proto"
	"ms-weather/weather-srv/service/collect"
	"time"
)

//定义数据处理结构体
type SetupData struct {
	opts             options
	sourceConfigInfo []*gjson.Json
}

// DATA_SEAS 天气数据类型常量，默认城市天气，常量为近海天气
const DATA_SEAS string = "data_seas"

//定义参数
type options struct {
	getDataType string
}
type Option func(c *options)

func GetDataType(sf string) Option {
	return func(opts *options) {
		opts.getDataType = sf
	}
}

// 构造函数，new一个对象并赋值
func NewSetupData(opts ...Option) *SetupData {
	setupDataOpts := options{}
	for _, opt := range opts {
		// 函数指针的赋值调用
		opt(&setupDataOpts)
	}

	setupData := new(SetupData)
	setupData.opts = setupDataOpts
	//定义api源配置信息文件列表
	var sourceFileList []string

	//从配置文件读取使用的源配置文件
	sourceFiles := g.Config().GetArray("system.sourceFiles")
	sourcePath := g.Cfg().GetString("system.sourceConfigPath")

	if setupData.opts.getDataType == DATA_SEAS {
		sourceFiles = g.Config().GetArray("system.seasSourceFiles")
		sourcePath = g.Cfg().GetString("system.seasSourceConfigPath")
	}

	sourceFileList = gconv.Strings(sourceFiles)

	//如果配置文件未配置，将使用目录内全部文件
	if len(sourceFiles) == 0 {
		//获取到源文件列表
		sourceConfFileList, _ := setupData.GetDataSource(sourcePath)
		sourceFileList = gconv.Strings(sourceConfFileList)
	}
	for i := 0; i < len(sourceFileList); i++ {
		//加载接口源的配置信息文件
		setupData.sourceConfigInfo = append(setupData.sourceConfigInfo, getApiConfig(sourcePath, sourceFileList[i]))
	}
	return setupData
}

// GetDataSource 获取接口源的列表
func (s *SetupData) GetDataSource(sourcePath string) ([]string, error) {
	glog.Info("源所在位置：", sourcePath)
	files, err := gfile.DirNames(sourcePath)
	if err != nil {
		return nil, err
	}
	return files, nil
}

// 跟据api源的配置文件，并按配置文件中约定的方式获取数据
func (s *SetupData) NowData(req *proto.DataReq, wd *proto.NowData) {

	nowDataCacheName := g.Config().GetString("system.serverName") + "_nowData"

	// 判断是否已经存入缓存
	if gcache.Get(nowDataCacheName) == nil {
		//从远程获取新的数据
		for i := 0; i < len(s.sourceConfigInfo); i++ {
			glog.Info("数据来源：", s.sourceConfigInfo[i].GetString("sourceApi"))

			//获取远程api数据
			apiBodyData, err := collect.GetAPIDataBody(s.sourceConfigInfo[i].GetString("sourceApi"))
			if err != nil {
				return
			}

			//跟据配置文件内的字段配置信息获取数据
			setNowData(apiBodyData, s.sourceConfigInfo[i], wd)
		}

		//存入缓存，跟据配置文件里的设置进行缓存
		var dataCacheTime int64 = g.Config().GetInt64("system.dataCacheTime")
		gcache.SetIfNotExist(nowDataCacheName, wd.Data, time.Duration(dataCacheTime)*time.Minute)

		//判断是否启用了数据库保存
		if g.Config().GetBool("system.isToDb") {
			//存入数据库
			if wd.City != "" && wd.Data.Tem != "" {
				go saveNowData(wd.Data)
			}
		}

	} else {
		//从缓存中提取数据
		glog.Info("数据来源：服务器缓存")
		cwd := gcache.Get(nowDataCacheName)
		wd.Data = (cwd).(*proto.WeatherData)
		glog.Info(wd.Data)
	}

}

//获取未来一周的天气情况
func (s *SetupData) ForecastData(req *proto.DataReq, fd *proto.ForecastData) {

	forecastDataCacheName := g.Config().GetString("system.serverName") + "_forecastData"

	// 判断是否已经存入缓存
	if gcache.Get(forecastDataCacheName) == nil {
		//从远程获取新的数据
		for i := 0; i < len(s.sourceConfigInfo); i++ {
			glog.Info("数据来源：", s.sourceConfigInfo[i].GetString("sourceApi"))

			//获取远程api数据
			apiBodyData, err := collect.GetAPIDataBody(s.sourceConfigInfo[i].GetString("sourceApi"))
			if err != nil {
				return
			}

			//跟据配置文件内的字段配置信息获取数据
			setForecastData(apiBodyData, s.sourceConfigInfo[i], fd)
		}

		//存入缓存，跟据配置文件里的设置进行缓存
		var dataCacheTime int64 = g.Config().GetInt64("system.dataCacheTime")
		gcache.SetIfNotExist(forecastDataCacheName, fd.Data, time.Duration(dataCacheTime)*time.Minute)

	} else {
		//从缓存中提取数据
		glog.Info("数据来源：服务器缓存")
		cfd := gcache.Get(forecastDataCacheName)
		fd.Data = (cfd).([]*proto.WeatherData)
		glog.Info(fd.Data)

	}

}

//获取海洋天气情况
func (s *SetupData) SeasData(req *proto.DataReq, sd *proto.SeasData) {

	seasDataCacheName := g.Config().GetString("system.serverName") + "_seasData"

	// 判断是否已经存入缓存
	if gcache.Get(seasDataCacheName) == nil {
		//从远程获取新的数据
		for i := 0; i < len(s.sourceConfigInfo); i++ {
			glog.Info("数据来源：", s.sourceConfigInfo[i].GetString("sourceApi"))

			//获取远程api数据
			apiBodyData, err := collect.GetAPIDataBody(s.sourceConfigInfo[i].GetString("sourceApi"))
			if err != nil {
				return
			}

			//跟据配置文件内的字段配置信息获取数据
			setSeasData(apiBodyData, s.sourceConfigInfo[i], sd)
		}

		//存入缓存，跟据配置文件里的设置进行缓存
		var dataCacheTime int64 = g.Config().GetInt64("system.dataCacheTime")
		gcache.SetIfNotExist(seasDataCacheName, sd.Data, time.Duration(dataCacheTime)*time.Minute)

	} else {
		//从缓存中提取数据
		glog.Info("数据来源：服务器缓存")
		cfd := gcache.Get(seasDataCacheName)
		sd.Data = (cfd).(*proto.SeasWeatherData)
		glog.Info(sd.Data)

	}
}
