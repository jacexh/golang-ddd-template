# golang-ddd-template
a DDD project template in golang

- Web Framework: [gin](https://github.com/gin-gonic/gin)
    * [Middlewares](github.com/jacexh/goutil/gin-middleware)
- Logging Library: [zap](https://github.com/uber-go/zap)
- Data Access Library
    * [gendry](https://github.com/didi/gendry)
    * [go-sql-driver](https://github.com/go-sql-driver/mysql)
- Configuration Library: [multiconfig](https://github.com/jacexh/multiconfig)

## Quick Start

```
# download project generator
go get -u github.com/jacexh/gdp@master

gdp
```

## Change Log

### 0.2.3

- 独立option模块，减少main.go文件代码量

### 0.2.2

- 使用`xorm.io`替换`github.com/didi/gendry`
- 升级gin/zap等依赖版本
- dto <--> entity, do <--> entity 转换层更加显性地表达

### 0.2.1

- 区分`DataObject`以及`Entity`
- 修改目录名称 `infrastructure/repository` -> `infrastructure/persistence`

### 0.2.0

基于依赖反转原则以及六边形架构重构整个项目

- 层名称变更：`sevice` -> `application`，以及`repository` -> `infrastructure`
- `Repository`定义在`Domain`层内
- `Application`+`Domain` 使用依赖反转，具备了更好的可测试性
- 更清晰的分支管理：`master`分支为golang项目，`template`分支为模板
- 严格区分了`Entity`、`ValueObject`、`DataTransferObject`、`DomainEvent`等
