# golang-ddd-template
a DDD project template in golang

- Web Framework: [gin](https://github.com/gin-gonic/gin)
    * [Middlewares](github.com/jacexh/goutil/gin-middleware)
- Logging Library: [zap](https://github.com/uber-go/zap)
- Data Access Library
    * [gendry](https://github.com/didi/gendry)
    * [go-sql-driver](https://github.com/go-sql-driver/mysql)
- Configuration Library: [multiconfig](https://github.com/jacexh/multiconfig)

## Change Log

### 0.2.0

基于依赖反转原则以及六边形架构重构整个项目

- 层名称变更：`sevice` -> `application`，以及`repository` -> `infrastructure`
- `Repository`定义在`Domain`层内
- `Application`+`Domain` 使用依赖反转，具备了更好的可测试性
- 更清晰的分支管理：`master`分支为golang项目，`template`分支为模板
- 严格区分了`Entity`、`ValueObject`、`DataTransferObject`、`DomainEvent`等