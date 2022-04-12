package main


import (
    "fmt"
    "github.com/tealeg/xlsx"
    //"strconv"
)


func main() {
    // var name :="2021年7月内部考核版直接费用追踪.xlsx"
    var name string

    fmt.Printf("Please enter your excel name: ")
    fmt.Scanf("%s", &name)
    fmt.Println("您要拆分的文件是：", name)

    // 读取文件

    xlfile, err := xlsx.OpenFile(name)

    if err != nil {
        fmt.Printf("Open file failed :  \n", err)
        return
    }

    for _, sheet := range xlfile.Sheets {
        fmt.Println("sheet name is : ", sheet.Name)

        var newFile *xlsx.File
        newFile = xlsx.NewFile()

        if err != nil {
            fmt.Println("文件拆分失败")
        }
        //var stt = *sheet
        //var st = &stt
        _, err := newFile.AppendSheet(*sheet,sheet.Name)

        if err != nil {
            fmt.Println("文件拆分失败1")
        }

        err = newFile.Save(sheet.Name + ".xlsx")

        if err != nil {
            fmt.Println("文件保存失败")
        }

    }
}
