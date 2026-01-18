<div align="center">
    <h1>KnowTime_Backend</h1>
    <!-- <img alt="GitHub License" src="https://img.shields.io/github/license/pykelysia/KnowTime_Backend"> -->
    <img alt="Backend" src="https://img.shields.io/badge/Go-1.24.5-blue?logo=go">
</div>

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
