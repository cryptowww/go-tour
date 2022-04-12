# GO Microservice

## 创建数据库

> CREATE DATABASE `gomicro` CHARACTER SET utf8 COLLATE utf8_general_ci;

## 创建工程

1. 新建工程主目录

> mkdir go-microservice

2. 新建子目录，管理数据库model

> mkdir dbmodel

3. 初始化包,包名go_microservice

> go mod init go_microservice

4. 新建main.go

5. 子包的引用方式

> "go_microservice/dbmodel"

6. 下载引用包

> go mod tidy


## 以上方式是，所有工程都在一个目录的方式
