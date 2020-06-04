#!/bin/sh

servername=weather
function help() {
    echo "$0 linux|windows|mac"
}

function linux(){
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

}
function windows(){
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build

}
function mac(){
    go build

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

mkdir -p bin/config
mkdir -p bin/db

cp -r config/. bin/config/
cp -r db/weather.sqlite bin/db/


if [ "$1" == "windows" ]; then
    cp weather-srv/weather-srv.exe bin/
    cp api-srv/api-srv.exe bin/
    rm -f weather-srv/weather-srv.exe
    rm -f api-srv/api-srv.exe

else
    cp weather-srv/weather-srv bin/
    cp api-srv/api-srv bin/
    rm -f weather-srv/weather-srv
    rm -f api-srv/api-srv
fi