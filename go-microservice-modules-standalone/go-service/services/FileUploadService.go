package service

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/gin-gonic/gin"
	"dbmodel/models"
	"bufio"
	"fmt"
	"io"
	//"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

//var db *gorm.DB
//var dberr error
//var constr string

func Fdownload(c *gin.Context) {
	// 新建一个数据库连接
	constr=fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "root", "localhost", 3306, "gomicro")
	db, dberr = gorm.Open(mysql.Open(constr),&gorm.Config{})
	if dberr != nil {
		panic("connect error!")
	}

	// 模拟新增一个人员信息
	p := models.PersonModel{Name:"Jack", Age:19, Gender:"M"}
	result := db.Create(&p)
	if result.Error != nil {
		fmt.Println("新增人员失败: ", result.Error)
		return
	}
	fmt.Printf("userId:%d, rowsAffected:%d /n",p.ID,result.RowsAffected)


	// 从Reader读取数据
	url := "https://images.cnblogs.com/cnblogs_com/wupeixuan/1186798/o_wallhaven-4d38m0.jpg"
	urls := strings.Split(url, "/")
	filename := urls[len(urls)-1]

	response, err := http.Get(url)
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}

	reader := response.Body
	defer reader.Close()
	// -----以下是把文件保存到服务器
	out, err1 := os.Create(filepath.Join("upload/", filename))
	if err1 != nil {
		fmt.Println("write err:", err)
	}

	defer out.Close()
	wf := bufio.NewWriter(out)

	_, err = io.Copy(wf, reader)

	if err != nil {
		fmt.Println("write err:", err)
	}
	wf.Flush()
}
