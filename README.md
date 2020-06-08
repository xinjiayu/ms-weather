# MS-WEATHER 

**天气预报第三方数据网关服务**

配置不同的第三方天气数据，统一的数据格式输出，确保天气数据的有效性，并同时进行数据本地化缓存与持久化存储。数据从远程获取通过本地缓存的功能，减少远程的api调用量。通过配置文件配置多个不同的api数据源，字段可以自动合并，输出最大的结果集。新加入数据源只要添加相应的配置文件就可以，无需重新启动服务。

系统采用微服务模式开发。

服务部署后，提供两个接口：

* 1、天气实况数据：/weather/now

* 2、天气预报数据（未来几天的）：/weather/forecast

* 3、近海天气数据 /weather/seas


## 特点

* 微服务架构，支持多种服务注册中心（mDNS、consul、eureka），默认为mdns

* API数据源数据元可以通过文件配置

* 数据可本地化存储，方便历史数据调用 

## 技术栈

基于go-micro、GoFrame开发的天气预报桥接服务

微服务框架：[go-micro](https://github.com/micro/go-micro) V2【 [中文文档](https://learnku.com/docs/go-micro/2.x) 】 、 [micro](https://github.com/micro/micro) V2【 [文档](https://micro.mu/docs/) 】

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

## 源码运行

**1、运行 Micro API网关服务**

micro api 即可启动api一个网关,默认的端口是8080

`$ micro api --handler=proxy`

```bash

$ micro --registry=consul --registry_address=127.0.0.1:8500 api

```

注意：micro2 默认支持的服务注册中心是mdns。micro1默认支持的是consul。

**2、运行micro web :  Weather API**

```
$ go run ./api-srv/main.go
Listening on [::]:64738

```

**3、运行 Weather Service**

进入到weather-srv目录

`$ go run ./weather-srv/main.go`



Curl API 测试Now 天气实况

```
$ curl http://127.0.0.1:8080/weather/now
{
  "message": "Hi, this is the Now API"
}
```

测试forecast 天气预报未来三天的天气情况

```
$ curl http://127.0.0.1:8080/weather/forecast
{
  "msg": "测试forecast"
}
```



## 部署运行

1，编译

需要跟据目标运行环境，编译。已编写了一个简单的脚本，进行交叉。

Linux系统：  `$ ./build.sh linux`

Windows系统：  `$ ./build.sh windows`

Mac系统：  `$ ./build.sh mac`

通过build.sh脚本编译后，工程下会生成一个bin目录。就是编译好的运行文件及相关的配置文件、数据库文件。直接复制到目标位置即可。

2，运行

1）运行micro API   `$ micro api`

2）运行Web Api   `$ ./api-srv`

3）运行micro srv   `$ ./weather-srv`

在浏览器中直接访问接口：http://127.0.0.1:8080/weather/now



## 服务发现

支持多种服务发现系统（mDNS、consul、eureka）。默认支的mDNS。

1、mDns 方式

无需添加注册服务中心相关的启动参数，默认支持。可以通过Micro 的web服务查看服务的注册情况。

```
$ micro web

2020-06-05 08:59:05  level=info service=web HTTP API Listening on [::]:8082

```

2、consul 方式


如果要使用consul，请按下面方式执行。

**运行consul服务**

`$ consul agent -dev`

**micro网关、srv、api三个服务启动加参数**

```bash
 --registry=consul

```

3、eureka 方式


**micro网关、srv、api三个服务启动加参数**

```bash
 --registry=eureka

```




## 天气数据源

一、sojson.com 
说明：https://www.sojson.com/blog/305.html

二、天气API
https://www.tianqiapi.com

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
