# MS-WEATHER 

**天气预报第三方数据桥接服务**

配置不同的第三方天气数据，统一的数据格式输出，确保天气数据的有效性，并进行数据本地化存储。数据从远程获取支持本地化缓存的功能，减少远程的api调用量。
配置多个不同的api数据源，字段可以自动合并，输出最大的结果集。

系统采用微服务模式开发。

## 特点

* 微服务架构，支持多种服务注册中心（mdns、consul、eureka），默认为mdns

* API数据源数据元可以通过文件配置

* 数据可本地化存储，方便历史数据调用 

## 技术栈

基于go-micro、GoFrame开发的天气预报桥接服务

微服务框架：[go-micro](https://github.com/micro/go-micro) 【 [中文文档](https://learnku.com/docs/go-micro/2.x) 】 、 [micro](https://github.com/micro/micro) 【 [文档](https://micro.mu/docs/) 】

web路由：[GoFrame](https://github.com/gogf/gf) 【 [中文文档](https://goframe.org/index) 】

服务管理：[consul](https://www.consul.io/) 【 [用户手册](https://kingfree.gitbook.io/consul/) 】

数据库：SQLite   【 [中文文档](https://doc.yonyoucloud.com/doc/wiki/project/sqlite/sqlite-intro.html) 】      *GO驱动使用  github.com/mattn/go-sqlite3* 【 [接口文档](https://godoc.org/github.com/mattn/go-sqlite3) 】

目录结构说明：

```
.
├── api-srv       weather对外暴露api接口服务，内部与RPC服务通讯。
├── bin           统一编译后生成的执行文件，可以拷走直接使用。包括了必要的运行文件
├── config        开发环境配置文件目录
├── db            开发环境数据库目录
├── resources     相关资源
├── tools         自行编译的Micro API网关服务
└── weather-srv   weather RPC服务
    ├── handler
    ├── proto
    ├── service
    └── units

```

## 运行


**运行 Micro API网关服务**

micro api 即可启动api一个网关,默认的端口是8080
可以通过--address=0.0.0.0:8080flag或者设置环境MICRO_API_ADDRESS=0.0.0.0:8080来修改

`$ micro api --handler=proxy`

```bash

micro --registry=consul --registry_address=127.0.0.1:8500 api

```


运行 Weather Service

`$ go run main.go`

运行 Weather API

```
$ go run ./api-srv/main.go
Listening on [::]:64738

```

## 服务发现

支持多种服务发现系统。默认情况下，服务发现基于组播DNS(mDNS)机制
如果要使用consul，请按下面方式执行。

**运行consul服务**

`consul agent -dev`

**srv、api两个服务启动加参数**
```bash
 --registry=mdns

```

Curl API
测试Now 天气实况
```
curl http://127.0.0.1:8080/weather/now
{
  "message": "Hi, this is the Now API"
}
```
测试forecast 天气预报未来三天的天气情况
```
curl http://127.0.0.1:8080/weather/forecast
{
  "msg": "测试forecast"
}
```



## 天气数据源

一、sojson.com 
说明：https://www.sojson.com/blog/305.html

二、天气API
https://www.tianqiapi.com/api?version=v6&appid=93247964&appsecret=MRlcpO5o&cityid=101020100

三、国家气象信息中心
http://data.cma.cn/market/index.html



## 其它资料

天气质数说明相关的

http://aqicn.org/city/beijing/cn/


可参考的天气预报接口说明：
https://open.caiyunapp.com/%E5%BD%A9%E4%BA%91%E5%A4%A9%E6%B0%94_API_%E4%B8%80%E8%A7%88%E8%A1%A8





## API数据源配置文件说明
配置文件的修改将即时生效，无所重启服务。

now 为天气实况信息

forecast 为天气预报信息，未来1--5天

filter 为进行数据过滤的设置。如：源数据中的温度值为：高温 35度，如果想只保留数字，就需要设置为 \\\D

filter的值为正则表达。系统会自动过滤掉通过正则表达式选取的内容。

*常用正则：*


`\\D`只保留数字

`\s.* `
以空格开始的所有字符

`[u4e00-\u9fa5]`
选择所有汉字

`[^\u4e00-\u9fa5]^[-,.?:;'\"!']`
选择所有非汉字，但是不包括-,.?:;'"!'这些标点符号



`^((?!abc).)*admin((?!abc).)*$`
包含admin且不包含abc。
