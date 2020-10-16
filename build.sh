#!/bin/sh

BuildVersion=`git describe --abbrev=0 --tags`
BuildTime=`date +%FT%T%z`
CommitID=`git rev-parse HEAD`

servername=weather
function help() {
    echo "$0 linux|windows|mac"
}

function linux(){
    echo "编译中..."
    CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=x86_64-linux-musl-gcc CGO_LDFLAGS="-static" go build -a -ldflags "-w -s -X main.BuildVersion=${BuildVersion} -X main.CommitID=${CommitID} -X main.BuildTime=${BuildTime}"
}
function windows(){
    CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -ldflags "-w -s -X main.BuildVersion=${BuildVersion} -X main.CommitID=${CommitID} -X main.BuildTime=${BuildTime}"
}
function mac(){
    go build -ldflags "-w -s -X main.BuildVersion=${BuildVersion} -X main.CommitID=${CommitID} -X main.BuildTime=${BuildTime}"
}
function copyFile() {
    echo "编译完成，正在发布到执行目录..."

    cd ../

    mkdir -p bin/config
    mkdir -p bin/db
    mkdir -p bin/log


    cp -r config/. bin/config/
    cp -r db/weather.sqlite bin/db/

    if [ "$1" == "windows" ]; then
        cp weather-srv/weather-srv.exe bin/
        cp api-srv/api-srv.exe bin/
        cp tools/micro/micro.exe bin/

        rm -f weather-srv/weather-srv.exe
        rm -f api-srv/api-srv.exe
        rm -f tools/micro/micro.exe

    else
        cp weather-srv/weather-srv bin/
        cp api-srv/api-srv bin/
        cp tools/micro/micro bin/

        rm -f weather-srv/weather-srv
        rm -f api-srv/api-srv
        rm -f tools/micro/micro

        cp curl.sh bin/

    fi
}



cd weather-srv

if [ "$1" == "" ]; then
    help
elif [ "$1" == "linux" ];then
    linux
elif [ "$1" == "windows" ];then
    windows
elif [ "$1" == "mac" ];then
    mac
fi

cd ../

cd api-srv

if [ "$1" == "linux" ];then
    linux
elif [ "$1" == "windows" ];then
    windows
elif [ "$1" == "mac" ];then
    mac
fi

cd ../

cd tools/micro

if [ "$1" == "linux" ];then
    linux
elif [ "$1" == "windows" ];then
    windows
elif [ "$1" == "mac" ];then
    mac
fi
cd ../

if  [ "$1" != "" ]; then
    copyFile
    echo "编译结束。"

fi


