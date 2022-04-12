package main

import "fmt"


// 递归方式,打印前n项
func fibonacci(n int){
    for i:=0; i <n; i++{
        // 递归
        //fmt.Println(fib(i))
        // 尾递归
        fmt.Println(fibTail(i,1,1))
    }
}
// 递归
func fib(n int) int {
    if n == 0  || n == 1{
        return n
    }

    return fib(n-1) + fib(n-2)
}
// 尾递归
func fibTail(n int, pre1 int, pre2 int) int{
    fmt.Println(n,pre1,pre2)
    if n == 0 {
        return n
    }else if n == 1 {
        return pre1
    }
    
    return fibTail(n-1,pre2,pre1+pre2)
}


func main(){
    fibonacci(10)
}

