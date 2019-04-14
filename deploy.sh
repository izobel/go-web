#! /bin/sh

kill -9 $(pgrep webserver)
cd /opt
git pull https://github.com/izobel/go-web.git
cd go-web/webserver
chmod -R 777 /opt/go-web
./webserver &   # & 表示程序继续运行