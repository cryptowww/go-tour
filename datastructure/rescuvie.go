package main

import "fmt"

// 递归求n的阶乘
func rescuvie(n int) int{
    if n == 0 {
        return 1
    }

    return n * rescuvie(n-1)
}

// 尾递归求n的阶乘
func rescuvieTail(n int, a int) int {
    if n == 1 {
        return a
    }

    return rescuvieTail(n-1, a*n)
}

func main(){
    fmt.Println(rescuvie(5))
    fmt.Println(rescuvieTail(5,1))
}
