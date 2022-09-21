# vientiane

## 项目架构说明

http: api 接口层，接受及返回 http 请求，调用 pub 进行业务逻辑处理

pub：grpc client，暴露 grpc 接口，具体逻辑实现在 server

server：接受对应请求进行处理，调用 db 等操作

- vientiane/http/main.go：http 服务入口
- vientiane/http/router：路由管理
- vientiane/http/handle：调用后端服务处理器
- vientiane/main.go: 后端服务入口
- dao：数据库接口层
- service：业务逻辑处理
- utils：通用工具

## 参考文档

- [Gin Web Framework](https://gin-gonic.com/zh-cn/docs/examples/query-and-post-form/)
- [GORM 中文文档](https://jasperxu.com/Programming/Golang/GORM/)
- [GORM 指南](https://gorm.io/zh_CN/docs/)
- [gRPC 官方文档中文版](http://doc.oschina.net/grpc?t=58009)
- [gRPC](https://www.grpc.io/docs/languages/go/quickstart/)
- [go 语言中文网](https://studygolang.com/)
- [go 官网](https://pkg.go.dev/)
- [go 学习资料](https://www.topgoer.com/)
- [go 语言设计与实现](https://draveness.me/golang/)
- 其他
  - https://github.com/tmrts/go-patterns/blob/master/idiom/functional-options.md
  - https://colobu.com/
  - https://icyfenix.cn/
  - https://www.programming-books.io/essential/go/

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
