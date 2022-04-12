package main

import "fmt"


//单循环链表
func sigCycleLink(){
    type linkNode struct {
        next *linkNode
        data int64
    }

    node1 := new(linkNode)
    node1.data = 1

    node2 := new(linkNode)
    node2.data = 2
    node1.next = node2

    node3 := new(linkNode)
    node3.data = 3
    node2.next = node3
    // 循环链表
    node3.next=node1

    // print node in a link
    nowNode := node1
    p := node1
    for {
        if nowNode != nil {
            fmt.Printf("node.data=%d,nextNode=%v.\n",nowNode.data,nowNode.next)
            nowNode = nowNode.next
            //fmt.Println(nowNode==p)
            //--begin 支持打印循环链表
            if nowNode == p {
                break
            }
            //--end
            continue
        }
    }
}

// 双循环链表


func main(){
    // sigle link
    sigCycleLink()
}
