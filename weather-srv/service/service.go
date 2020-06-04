package service

import (
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
type SetupData struct{}

// GetDataSource 获取接口源的列表
func (s *SetupData) GetDataSource() ([]string, error) {
	sourcePath := g.Cfg().GetString("system.sourceConfigPath")
	glog.Info("源所在位置：", sourcePath)
	files, err := gfile.DirNames(sourcePath)
	if err != nil {
		return nil, err
	}
	return files, nil
}

// 跟据api源的配置文件，并按配置文件中约定的方式获取数据
func (s *SetupData) ApiData(req *proto.DataReq, wd *proto.NowData) {
	//定义api源配置信息文件列表
	var sourceFileList []string

	//从配置文件读取使用的源配置文件
	sourceFiles := g.Config().GetArray("system.sourceFiles")
	sourceFileList = gconv.Strings(sourceFiles)

	//如果配置文件未配置，将使用目录内全部文件
	if len(sourceFiles) == 0 {
		//获取到源文件列表
		sourceConfFileList, _ := s.GetDataSource()
		sourceFileList = gconv.Strings(sourceConfFileList)
	}

	if len(sourceFileList) == 0 {
		return
	}

	//进行缓存处理，减少远程api调用次数
	serverName := g.Config().GetString("system.serverName")
	if gcache.Get(serverName) == nil {
		for i := 0; i < len(sourceFileList); i++ {
			//加载接口源的配置信息文件
			sc := getApiConfig(sourceFileList[i])
			if sc == nil {
				return
			}
			glog.Info("数据来源：", sc.GetString("sourceApi"))
			//获取远程api数据
			apiBodyData, err := collect.GetAPIDataBody(sc.GetString("sourceApi"))
			if err != nil {
				return
			}
			//跟据配置文件内的字段配置信息获取数据
			setData(apiBodyData, sc, wd)
		}

		//存入缓存，跟据配置文件里的设置进行缓存
		var dataCacheTime int64 = g.Config().GetInt64("system.dataCacheTime")
		gcache.SetIfNotExist(serverName, wd.Data, time.Duration(dataCacheTime)*time.Minute)

		//判断是否启用了数据库保存
		if g.Config().GetBool("system.isToDb") {
			//存入数据库
			if wd.City != "" && wd.Data.Tem != "" {
				go saveData(wd.Data)
			}
		}

	} else {
		glog.Info("数据来源：服务器缓存")
		cwd := gcache.Get(serverName)
		wd.Data = (cwd).(*proto.WeatherData)
	}

}
