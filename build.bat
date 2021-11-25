@echo off

rem 编译
set GOOS=linux
set GOARCH=amd64
go build -o screen main.go

rem 打包
tar -cvf screen.tar.gz ./screen ./etc/* ./external/cert/*

rem 发送
scp ./screen.tar.gz root@10.170.139.10:/home/web/back/screen/