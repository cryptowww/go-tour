package main

import (
	"fmt"
	"os"
	"path/filepath"
	"io"
	"net/http"
	"bufio"
)

func main() {
//	osOp()
	ioOp()
	bufioOp()
}

func osOp() {
	fmt.Println("this is OS")
	// 获取当前目录
	cdir, err := os.Getwd()
	if err != nil {
	//	goto errlabel
	}

	fmt.Println("current Dir:",cdir)

	// 读取文件
	fmt.Println("------readfile-----")
	data,err1 := os.ReadFile(filepath.Join(cdir,"f1.txt"))
	if err1 != nil {
		fmt.Println(err1)
		//goto errlabel
	}
	
	fmt.Println("------writefile-----")
	// 写入文件
	err = os.WriteFile("f2.txt",data,0666)
	fmt.Println(string(data))


	fmt.Println("------readDir-----")
	// 读取目录
	files, err2 := os.ReadDir("F:\\gospace\\")
	//files, err2 := os.ReadDir(".")
	if err2 != nil {
		fmt.Println(err2)
	}

	for _, f := range files {
		fmt.Println(f.IsDir(), f.Name())
	}

	fmt.Println("------mkDirTemp-----")
	// 新建临时目录
	tdir, err3 := os.MkdirTemp("","example")
	if err3 != nil {
		fmt.Println(err3)
	}
	// clean临时目录，程序结束后，临时目录和下边的文件全部清理
	defer os.RemoveAll(tdir)// clean up
	fmt.Println(tdir)

	f := filepath.Join(tdir,"tempfile.txt")
	temp := []byte("this is a tempfile 很好")
	err3 = os.WriteFile(f,temp,0444)
	if err3 != nil {
		fmt.Println(err3)
	}

	fmt.Println("------createTemp-----")
	// 新建临时文件
	f1, err4 := os.CreateTemp(tdir,"example")
	if err4 != nil {
		fmt.Println(err)
	}
	// 清理临时文件
	defer f1.Close()
	defer os.Remove(f1.Name())
	_, err = f1.Write(temp)
	// or
	//err = os.WriteFile(f1.Name(),temp,0444)
	if err != nil {
		fmt.Println(err)
	}

	// 目录状态
	sdir, err5 := os.Stat(tdir)
	if err5 != nil {
		fmt.Println(err)
	}
	fmt.Println(sdir)
	fmt.Println(sdir.Size())
	// 创建文件
	fmt.Println("------Create-----")
	
	f1,err4 = os.Create("newf.txt")
	if err4 != nil {
		fmt.Println(err)
	}
	defer f1.Close()

	finfo, _ := f1.Readdir(-1)
	for _,f := range finfo {
		fmt.Println(f.Name())
	}
	f1.WriteString("hello,世界")

	sns,_ := f1.Readdirnames(-1)
	
	for _,s := range sns {
		fmt.Println(s)
	}
	// 创建一个新的文件句柄
	f11 := os.NewFile(f1.Fd(),"a-newf.txt")
	defer f11.Close()
	fmt.Println(f11)
}

func ioOp() {
	fmt.Println("this is IO")
	// copy流复制，占用内存少，属于指针复制，适用于大文件下载
	url := "http://archive.ics.uci.edu/ml/machine-learning-databases/00607/synchronous%20machine.csv"
	resp, err1 := http.Get(url)
	if err1 != nil {
		fmt.Println(err1)
	}
	
	defer resp.Body.Close()
	
	out, err2 := os.Create("synchronous_machine1.csv")
	if err2 != nil {
		fmt.Println(err2)
	}
	
	wf := bufio.NewWriter(out)
	
	defer out.Close()

	// ----CopyBuffer
	//buf := make([]byte, 1024)
	//n, err3 := io.CopyBuffer(wf,resp.Body,buf)

	// --CopyN
	n, err3 := io.CopyN(wf,resp.Body,10)
	// --Copy
	//n, err3 := io.Copy(wf,resp.Body)
	fmt.Printf("write %d.\n", n)

	if err3 != nil {
		panic(err3)
	}
	wf.Flush()
}

func bufioOp() {
	fmt.Println("this is BUFIO")
}
