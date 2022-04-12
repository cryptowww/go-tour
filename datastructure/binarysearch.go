package main

import "fmt"

// 二分法查找
// array 待查找数组
// target 查找目标
// lside、rside 左右边界
func binarySearch(array []int, target int, lside,rside int) int {
    // 出界
    if lside > rside {
        return -1
    }
    // 找中点
    mid := (lside + rside) / 2
    midNum := array[mid]

    if midNum == target {
        return mid
    }else if midNum > target {
        return binarySearch(array, target, 0, mid - 1)
    }else {
        return binarySearch(array, target, mid + 1, rside)
    }
}

func main(){
    array := []int{1, 5, 9, 15, 81, 89, 123, 189, 333}
    target := 500
    result := binarySearch(array, target, 0, len(array)-1)
    fmt.Printf("target=%d,index=%d \n",target,result)

    target = 189
    result = binarySearch(array, target, 0, len(array)-1)
    fmt.Printf("target=%d,index=%d \n",target,result)
}

