package main

import (

    "fmt"
    "time"
)

var strChan = make(chan string,3)

func main(){
    syncChan1 := make(chan struct{},1)
    syncChan2 := make(chan struct{},2)

    go func(){
        <- syncChan1
        fmt.Println("Received and wait a second...")
        time.Sleep(time.Second)
        for {
            if elem, ok := <-strChan ; ok {
                fmt.Println("Received: ", elem)
            } else {
                break
            }
        }
        fmt.Println("stopped")
        syncChan2 <- struct{}{}
    }()

    go func(){
        for _, elem := range []string{"a","b","c","d"} {
            strChan <- elem
            fmt.Println("sent: ",elem)
            if elem == "c"{
                syncChan1 <- struct{}{}
                fmt.Println("Sent a Sync signal")
            }
        }
        fmt.Println("wait 2 seconds")
        time.Sleep(time.Second * 2)
        close(strChan)
        syncChan2 <- struct{}{}
    }()
    <- syncChan2
    <- syncChan2
}
