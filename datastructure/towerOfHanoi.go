package main

import "fmt"

var total = 0

func main(){
    // 盘子数量
    n := 4
    a,b,c := "a","b","c"

    tower(n, a, b, c)

    fmt.Println(total)
}
// 从a借助b，挪到c
func tower(n int, a,b,c string){
    if n==1{
        total = total +1
        fmt.Println(a,"=>",c)
        return
    }

    tower(n-1,a,c,b)

    total = total +1
    fmt.Println(a,"=>",c)
    tower(n-1,b,a,c)
}
