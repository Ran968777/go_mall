#!/bin/bash

# 清理旧的构建文件
rm -rf build
mkdir -p build

# 构建当前平台版本
echo "Building..."
go build -o build/go-mall main.go

echo "Build complete! Output: build/go-mall" 