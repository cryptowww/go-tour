// POST方式请求 c.PostForm
// 路径参数/:id c.Param
// get参数?name=a&age=18 c.Query
package service

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/gin-gonic/gin"
	"dbmodel/models"
	"fmt"
	//"log"
	"strconv"
	"net/http"
)

var db *gorm.DB
var dberr error
var constr string

func init() {
	// 新建一个数据库连接
	constr=fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "root", "localhost", 3306, "gomicro")
	db, dberr = gorm.Open(mysql.Open(constr),&gorm.Config{})
	if dberr != nil {
		panic("connect error!")
	}
}

// 新增一个人员
func AddItem(c *gin.Context) {
	// 获取人员信息
	name	:= c.PostForm("name")
	age		:= c.PostForm("age")
	gender	:= c.PostForm("gender")

	fmt.Printf("name:%s, age:%s, gender:%s \n",name,age,gender)
	uintAge, err := strconv.Atoi(age)

	if err != nil {
		fmt.Println("年龄有误",err)
		return
	}

	// 模拟新增一个人员信息
	p := models.PersonModel{Name:name, Age:uintAge, Gender:gender}
	
	result := db.Create(&p)

	if result.Error != nil {
		fmt.Println("新增人员失败: ", result.Error)
		//return
		c.JSON(http.StatusNoContent,gin.H{
			"message": "请检查参数，新增人员失败",
		})
	}

	msg := fmt.Sprintf("新增人员>>userId:%d, rowsAffected:%d",p.ID,result.RowsAffected)

	fmt.Println(msg)

	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}


// 新增一个人员
// 通过变量绑定的方式
func AddItem2(c *gin.Context) {
	// 获取人员信息
	var ps models.PersonModel

	// 使用绑定，要注意大小写，传值和struct的大小写要保持一致
	if err := c.ShouldBind(&ps); err != nil {
		fmt.Println("获取参数失败")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "获取参数失败",
		})
		return
	}


	
	fmt.Printf("name:%s, age:%s, gender:%s\n",ps.Name,ps.Age,ps.Gender)

	// 模拟新增一个人员信息
	//p := models.PersonModel{Name:name, Age:uintAge, Gender:gender}
	
	result := db.Create(&ps)

	if result.Error != nil {
		fmt.Println("新增人员失败: ", result.Error)
		//return
		c.JSON(http.StatusNoContent,gin.H{
			"message": "请检查参数，新增人员失败",
		})
	}

	msg := fmt.Sprintf("新增人员>>userId:%d, rowsAffected:%d",ps.ID,result.RowsAffected)

	fmt.Println(msg)

	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}

// 查询所有人员
func QueryAllItem(c *gin.Context) {
	// 查询变量
	var allPerson []models.PersonModel
	// 查询所有
	db.Find(&allPerson)

	if len(allPerson) <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"message":"没有数据",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data": allPerson,
	})
}

// 按照id查询
// 路径参数的查询方式：c.Param()
func QueryById(c *gin.Context){
	// 查询变量
	var person models.PersonModel
	// 获取参数id
	uid := c.Param("id")

	// 使用变量绑定的方式
//	if err := c.ShouldBind(&person); err != nil {
//		fmt.Println("获取参数失败")
//		c.JSON(http.StatusBadRequest, gin.H{
//			"message": "获取参数失败",
//		})
//		return
//	}

	fmt.Printf("查询id：%d\n",uid)

	// 查询
	db.First(&person, uid)

	if person.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message":"数据不存在",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data": person,
	})
}

// 按照id更新
func UpdateById(c *gin.Context) {
	// 查询变量
	var person models.PersonModel
	// 获取参数id
	uid := c.Param("id")

	// 查询
	db.First(&person, uid)

	if person.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message":"数据不存在",
		})

		return
	}
	// 更新单个字段
	// db.Model(&person).Update("name", c.PostForm("name"))
	//var toUpdateMap  map[string]interface
	toUpdateMap := make(map[string]interface{}, 3)

	age := c.PostForm("age")
	if age != "" {
		uintage, _ := strconv.Atoi(age)
		toUpdateMap["age"] = uint8(uintage)
	}
	toUpdateMap["name"] = c.PostForm("name")
	toUpdateMap["gender"] = c.PostForm("gender")

	
	// 更新多列
	//db.Model(&person).Updates(models.PersonModel{Name: c.PostForm("name"), Age: uint8(c.PostForm("age")), Gender: c.PostForm("gender")})
	db.Model(&person).Updates(toUpdateMap)

	c.JSON(http.StatusOK, gin.H{
		"message": "update ok",
		"data": person,
	})

}

// 按照id删除
func DeleteById(c *gin.Context) {
	// 查询变量
	var person models.PersonModel
	// 获取参数id
	uid := c.Param("id")

	// 查询
	db.First(&person, uid)

	if person.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message":"数据不存在",
		})
		fmt.Println("数据不存在")
		return
	}

	db.Delete(&person)

	c.JSON(http.StatusOK, gin.H{
		"message": "delete ok",
		"data": person,
	})

}
