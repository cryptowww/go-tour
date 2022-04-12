package main

import (
	"go_service_1/services"
	"github.com/mattn/go-colorable"
	"github.com/gin-gonic/gin"
	"fmt"
	"os"
	"io"
	"log"
)


// 通过postman调用post请求 http://localhost:9999/api/v1/fdownload
// 可以看到数据库新增了一条记录
// 同时，磁盘下载了一个图片文件
func main() {
	// 启用gin的日志输出带颜色
	gin.ForceConsoleColor()
	// 替换默认Writer（关键步骤）,解决日志乱码问题
	//gin.DefaultWriter = colorable.NewColorableStdout()
	// 写日志到文件和控制台
	logfile, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(logfile,colorable.NewColorableStdout())
	// 定义日志格式
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	r := gin.Default()
	// 规划路由
	v1 := r.Group("api/v1/person") 
	{
		// 新增
		// curl -d "name=mini&age=22&gender=F" http://localhost:9999/api/v1/person/additem
		// curl -d "name=GGG" -d "age=22" -d "gender=M" http://localhost:9999/api/v1/person/additem

		v1.POST("/additem",service.AddItem2)
		// 查询所有
		// curl -s -X GET http://localhost:9999/api/v1/person/queryall
		v1.GET("/queryall",service.QueryAllItem)
		// 按照id查询
		// curl -s -X GET http://localhost:9999/api/v1/person/get/4
		v1.GET("/get/:id",service.QueryById)
		// 按照id更新
		// curl -d "name=SLLLL" -X PUT http://localhost:9999/api/v1/person/update/2
		// curl -d "name=gold" -d "gender=F" -X PUT http://localhost:9999/api/v1/person/update/7
		v1.PUT("/update/:id",service.UpdateById)
		// 按照id删除
		// curl -X DELETE http://localhost:9999/api/v1/person/delete/8
		v1.DELETE("/delete/:id",service.DeleteById)
	}

	fmt.Println("Server Started!")
	r.Run(":9999")
}

