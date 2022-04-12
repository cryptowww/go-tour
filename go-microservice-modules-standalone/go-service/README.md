# 开发一个go模块化的微服务go-microservice

## 引用工程外的module（FileUploadService引用dbmodel）

如果没有放到github，把他们放在本地同级目录，使用以下方式

> import "dbmodel/models"

参见services/FileUploadService.go

## 首先mod初始化

> go mod init go_service_1

## 使用replace命令替换到本地目录

> go mod edit --replace dbmodel=../dbmodel

这样就可以引用本地的另外module了

## 引用工程下的module（services）
	
> import "go_service_1/services"

其中go_service_1是主module名字，参见mod初始化时的命名

## 下载包

> go mod tidy

## 关于引用services包

请注意虽然目录名是services（多个s），其实包名我写的是service（没有s），在main函数调用是要写service而不是services

## 执行测试

> go run main.go


用postman调用一个post测试

http://localhost:9999/api/v1/fdownload

可以看到：

1. 数据库中新增了一条记录
2. 下载了一张图片（在工程下upload目录，记得提前新建目录）

---

## dbmodel

这是一个独立的模块，数据层可以通过这种方式分离出来，单独管理。
以此类推，其他公共的服务和模块也可以单独分离出来，单独管理。

