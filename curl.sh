#!/bin/bash

logfile=log/app.log

nohup ./micro --api_address=127.0.0.1:8088 api --handler=proxy  &> $logfile &
nohup ./weather-srv  &> $logfile &
nohup ./api-srv  &> $logfile &