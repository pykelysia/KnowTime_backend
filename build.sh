#!/bin/bash

# 获取当前Git Commit ID
GIT_COMMIT=$(git rev-parse HEAD)
GIT_COMMIT_SHORT=$(git rev-parse --short HEAD)

# 获取当前时间作为编译时间
BUILD_TIME=$(date -u +%Y-%m-%dT%H:%M:%SZ)

# 获取版本号 (如果有的话)
VERSION=${1:-"1.0.0"}

echo "Building KnowTime Backend..."
echo "Git Commit: $GIT_COMMIT_SHORT"
echo "Build Time: $BUILD_TIME"
echo "Version: $VERSION"

# 编译时注入版本信息
go build -ldflags "-X 'knowtime/internal.GitCommit=$GIT_COMMIT_SHORT' -X 'knowtime/internal.BuildTime=$BUILD_TIME' -X 'knowtime/internal.Version=$VERSION'" -o knowtime .

if [ $? -eq 0 ]; then
    echo "Build successful!"
    echo "Binary: ./knowtime"
else
    echo "Build failed!"
    exit 1
fi