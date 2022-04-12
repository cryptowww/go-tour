package main

import (
    "fmt"
    "time"
)

type Counter struct {
    count int
}

func (counter *Counter) String() string{
    return fmt.Sprintf("{count: %d}", counter.count)
}

var mapChan = make(chan map[string]*Counter, 1)

func main(){
    syncChan := make(chan struct{},2)

    // receive op
    go func(){
        for {
            if elem, ok := <- mapChan; ok {
                counter := elem["count"]
                counter.count ++
                fmt.Println(counter.String())
            } else {
                break
            }
        }
        
        fmt.Println("Stopped. [reciver]")
        syncChan <- struct{}{}
    }()

    // sender op
    go func(){
        countMap := map[string]*Counter{
            "count": &Counter{},
        }

        for i := 0; i < 5; i ++ {
            mapChan <- countMap
            time.Sleep(time.Millisecond)
            fmt.Printf("The count map: %v. [sender] \n", countMap)
        }
        close(mapChan)
        syncChan <- struct{}{}
    }()

    <- syncChan
    <- syncChan
}
