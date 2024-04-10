#!/bin/bash

#构建一个自动化脚本
RUN_NAME=test
mkdir -p output/bin
cp script/* output 2>/dev/null
chmod +x output/start.sh
go build -o output/bin/${RUN_NAME}
./output/start.sh