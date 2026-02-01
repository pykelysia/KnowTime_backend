<div align="center">
    <h1>知时后端</h1>
    <img alt="Backend" src="https://img.shields.io/badge/Go-1.24.5-blue?logo=go">
</div>

## 项目介绍

### 项目背景

在当前的互联网社会，手机等电子产品逐渐成为年轻人的刚需。在深度使用各大APP方便日常生活的同时，也为使用者带来了相应的后果。诸如视力衰退、沉迷网游、拖延症等各种手机“病症”也随之而来。因此急需要一款能够有效统计和管控手机等电子产品的使用时长、细分使用内容的APP来帮助使用者更有效地利用时间，形成合理的电子设备使用习惯并提出有操作性的使用建议。

## 项目内容

针对这一社会现象，我们决定开发“KnowTime”知时APP。它能通过多种方式收集用户的手机及其他电子产品的使用情况，通过分析各类APP打开次数和使用时长，推测用户使用内容和使用行为。根据用户终端性能进行端侧或云端AI分析，针对用户的电子产品使用情况，给出合理的建议和分析报告，帮助用户养成良好的电子产品使用习惯。

## 部署方法

推荐使用Docker Compose一键部署，首先获得Compose文件：
```bash
curl -o docker-compose.yml https://raw.githubusercontent.com/pykelysia/KnowTime_Backend/main/docker-compose.yml
```

然后启动：
```bash
docker compose up -d
```

## API响应格式统一化说明

### 响应结构
```json
{
  "errcode": 0,
  "message": "Success",
  "data": {
    // 响应数据
  }
}
```

### 字段说明
- `errcode` (int): 错误码，0表示成功，非0表示失败
- `message` (string): 错误消息或成功提示
- `data` (any, optional): 响应数据，成功时包含实际返回数据，失败时可能为null

## 错误码定义

### 成功
- `0` - Success

### 客户端错误 (1xxx)
- `1001` - Invalid request (无效的请求参数)
- `1002` - Invalid request body (无效的请求体)
- `1003` - User not authenticated (未授权)
- `1004` - User not found (用户不存在)
- `1005` - Password error (密码错误)

### 服务器错误 (2xxx)
- `2001` - Internal server error (服务器内部错误)
- `2002` - Database error (数据库错误)
- `2003` - Failed to create time event (创建时间事件失败)
- `2004` - Failed to update time event (更新时间事件失败)
- `2005` - Failed to generate token (生成Token失败)
- `2006` - Failed to build ReAct agent (构建Agent失败)
- `2007` - Failed to call Agent (调用Agent失败)
- `2008` - Failed to create user (创建用户失败)
- `2009` - Database connection failed (数据库连接失败)

## 响应示例

### 成功响应
```json
{
  "errcode": 0,
  "message": "Success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIs...",
    "u_id": 123
  }
}
```

### 错误响应
```json
{
  "errcode": 1005,
  "message": "Password error",
  "data": null
}
```
