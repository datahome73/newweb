#! /bin/sh

kill -9 $(pgrep webserver)
cd ~/go/src/newweb/
git pull https://github.com/datahome73/newweb.git
cd webserver/
go build
./webserver &