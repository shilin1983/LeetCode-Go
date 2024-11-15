#!/bin/bash

echo "测试开始..."

go test -v -coverprofile=coverage.out -covermode=atomic ./src...

go tool cover -html=coverage.out

echo "测试结束！"
