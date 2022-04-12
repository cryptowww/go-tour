package service

import (
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	//"github.com/stretchr/testify/assert"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"encoding/json"
	"bytes"
)

// 定义路由
var router *gin.Engine

// 初始化
func init() {
	router = gin.Default()
	gin.ForceConsoleColor()
	colorable.NewColorableStdout()
	// 规划路由
	v1 := router.Group("api/v1/person")
	{
		v1.POST("/additem", AddItem)
		v1.GET("/queryall", QueryAllItem)
		v1.GET("/get/:id", QueryById)
		v1.PUT("/update/:id", UpdateById)
		v1.DELETE("/delete/:id", DeleteById)
	}

}

// ParseToStr 将map中的键值对输出成querystring形式
func ParseToStr(mp map[string]string) string {
	values := ""
	for key, val := range mp {
		values += "&" + key + "=" + val
	}
	temp := values[1:]
	values = "?" + temp
	return values
}

// 公共GET请求方法
func Get(uri string, router *gin.Engine) []byte {
	req := httptest.NewRequest("GET", uri, nil)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	result := w.Result()

	defer result.Body.Close()

	body, _ := ioutil.ReadAll(result.Body)
	return body
}

// 公共POST请求方法
// 构造POST请求，表单数据以 querystring 的形式加在uri之后
// 这种方式，POST 请求获取参数是时要调用 c.Query("users")，而不是c.PostFprm("users")，更不是c.Param("users)
func PostForm(uri string, param map[string]string, router *gin.Engine) []byte {
	req := httptest.NewRequest("POST", uri+ParseToStr(param), nil)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	result := w.Result()

	defer result.Body.Close()

	body, _ := ioutil.ReadAll(result.Body)
	return body
}

// PostJson 根据特定请求uri和参数param，以Json形式传递参数，发起post请求返回响应
func PostJson(uri string, param map[string]interface{}, router *gin.Engine) []byte {
	jsonByte, _ := json.Marshal(param)
	req := httptest.NewRequest("POST", uri, bytes.NewReader(jsonByte))

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	result := w.Result()

	defer result.Body.Close()

	body, _ := ioutil.ReadAll(result.Body)
	return body
}

// 测试QueryAllItem方法
func TestQueryAllItem(t *testing.T) {
	// 请求地址
	uri := "/api/v1/person/queryall"

	// send Get请求
	body := Get(uri, router)

	fmt.Printf("response: %v\n", string(body))
}


func TestAddItem(t *testing.T) {
	uri := "/api/v1/person/additem"

	param := make(map[string]interface{})
	param["name"] = "Json"
	param["age"] = "10"
	param["gender"] = "M"

	body := PostJson(uri, param, router)

	fmt.Printf("response: %v\n", string(body))
}
