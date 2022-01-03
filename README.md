# vientiane

## 项目架构设定与说明

http: api 接口层，接受及返回 http 请求，调用 pub 进行业务逻辑处理

pub：grpc client，暴露 grpc 接口，具体逻辑实现在 server

server：接受对应请求进行处理，调用 db 等操作

## 接口文档

接口返回的数据标准格式

```json
{
  "code": 200,
  "message": "ok",
  "data": {
    "res": [],
    "offset": 20,
    "count": 100
  }
}
```

### 账号

列出账号 todo

获取账号详情 todo

更新账号信息 todo

