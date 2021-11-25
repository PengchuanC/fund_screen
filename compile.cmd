@echo off

echo compiling proto buffer files for Golang and Python now

cd rpc

protoc --go_out=plugins=grpc:../services/ *.proto
python -m grpc_tools.protoc --python_out=../external/python/ --grpc_python_out=../external/python/ -I. *.proto

cd ../