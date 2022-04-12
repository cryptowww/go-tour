# go-service

## 引用工程外的module

如果没有放到github，把他们放在本地同级目录，使用以下方式

## 首先初始化

> go mod init go_service_1

## 使用replace命令替换到本地目录

> go mod edit --replace dbmodel=../dbmodel

这样就可以引用本地的另外module了


## 下载包

> go mod tidy

## 执行测试

> go run main.go
