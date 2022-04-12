package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	//"time"
    "strings"
)

func main() {
	var name string

	for {
		fmt.Println("请输入要拆分的文件名: ")
		fmt.Scanf("%s", &name)
		fmt.Println("您要拆分的文件是：", name)
		getSheetNames(name)
	}
	/*
	   f, err := excelize.OpenFile(name)
	   if err != nil {
	       fmt.Println(err)
	       return
	   }

	   for index, name := range f.GetSheetMap() {
	       //if f.GetSheetVisible(name) {
	           fmt.Println(index, name, f.GetSheetVisible(name))

	           nf := *f
	           newf := &nf
	           for _, name1 := range newf.GetSheetMap(){
	               if name1 != name {
	                   newf.DeleteSheet(name1)
	               }
	           }
	           fmt.Println("新文件：", name+".xlsx")
	           //err = newf.SaveAs(name)
	           if err := newf.SaveAs(name+".xlsx"); err != nil {
	               fmt.Println("拆分文件错误: ", err)
	           }
	       //}
	   }*/

}

func getSheetNames(fpath string) {
	f, err := excelize.OpenFile(fpath)
	if err != nil {
		fmt.Println(err)
		return
	}

	/* var names [len(f.GetSheetMap())]string
	   var id =0
	   for _, name := range f.GetSheetMap() {
	       names[id] = name
	       id ++
	   }*/

	for _, name := range f.GetSheetMap() {
		split1(fpath, name)
		//time.Sleep(10)
	}

}

func split(name string) {

	f, err := excelize.OpenFile(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	for index, name := range f.GetSheetMap() {
		//if f.GetSheetVisible(name) {
		fmt.Println(index, name, f.GetSheetVisible(name))
		/*if !f.GetSheetVisible(name) {
		    fmt.Println(name, "是一个隐藏sheet，不拆分")
		    continue
		}*/
		nf := *f
		newf := &nf
		for _, name1 := range newf.GetSheetMap() {
			if name1 != name {
				newf.DeleteSheet(name1)
			}
		}
		fmt.Println("新文件：", name+".xlsx")
		//err = newf.SaveAs(name)
		if err := newf.SaveAs(name + ".xlsx"); err != nil {
			fmt.Println("拆分文件错误: ", err)
		}
		//}
	}
}

func split1(fpath string, name string) {

	f, err := excelize.OpenFile(fpath)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, name1 := range f.GetSheetMap() {
		//if f.GetSheetVisible(name) {
		//fmt.Println(index, name, f.GetSheetVisible(name))
		/*if !f.GetSheetVisible(name) {
		    fmt.Println(name, "是一个隐藏sheet，不拆分")
		    continue
		}*/
		//nf := *f
		//newf := &nf
		//for _, name1 := range newf.GetSheetMap(){
		if name1 != name {
			f.DeleteSheet(name1)
		}
		//}
		//fmt.Println("新文件：", name+".xlsx")
		//err = newf.SaveAs(name)
		//}
	}
    var fname = fpath[0:strings.Index(fpath,".xlsx")]

    //fmt.Println(fname)
	if err := f.SaveAs(fname + "-" + name + ".xlsx"); err != nil {
		fmt.Println("拆分文件错误: ", err)
	}
}
