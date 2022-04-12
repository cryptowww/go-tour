# dbmodel

## 初始化

> go mod init dbmodel


## 下载包

> go mod tidy

## main.go

引入models包的方式：
> import "dbmodel/models"


## 实现对db初始化和表结构变更的管理

包结构：

+-dbmodel
+---main
-----main.go
+---models
-----models.go

models实现对表结构的映射管理
main实现数据库迁移

外部工程使用时，使用方式:

```go
// 引用
import "dbmodel/models"

// 代码
&models.PersonModel

```

## 表结构迁移

通过main.go来实现，本模块可以作为一个单独的模块来管理数据层
