# 使用官方Golang镜像作为基础镜像
FROM golang:1.24.5-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制go.mod和go.sum文件（如果存在）并下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 设置构建参数
ARG GIT_COMMIT=unknown
ARG BUILD_TIME=unknown
ARG VERSION=1.0.0

# 构建应用程序并注入版本信息
RUN go build -ldflags "-X 'knowtime/internal.GitCommit=${GIT_COMMIT}' -X 'knowtime/internal.BuildTime=${BUILD_TIME}' -X 'knowtime/internal.Version=${VERSION}'" -o knowtime .

# 使用轻量级Alpine镜像作为最终镜像
FROM alpine:latest

# 安装必要的软件包
RUN apk --no-cache add ca-certificates tzdata

# 设置时区（可选）
ENV TZ=Asia/Shanghai

# 创建非root用户
RUN addgroup -g 65532 nonroot && \
    adduser -D -u 65532 -G nonroot nonroot

# 设置工作目录
WORKDIR /app

# 从builder阶段复制二进制文件
COPY --from=builder /app/knowtime .

# 切换到非root用户
USER nonroot

# 暴露应用端口
EXPOSE 8080

# 启动命令
CMD ["./knowtime"]