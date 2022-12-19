package main

import (
	"fmt"
	"time"
)

func main () {
    // Fungsi time.Sleep()
    // fmt.Println("start")
    // time.Sleep(time.Second * 4)
    // fmt.Println("after 4 seconds")

    // Scheduler Menggunakan time.Sleep()
    // for true {
    //     fmt.Println("Hello !!")
    //     time.Sleep(1 * time.Second)
    // }

    // Fungsi time.NewTimer()
    // var timer = time.NewTimer(4 * time.Second)
    // fmt.Println("start")
    // <-timer.C
    // fmt.Println("finish")


    // Fungsi time.AfterFunc()
    var ch = make(chan bool)

    time.AfterFunc(4*time.Second, func() {
        fmt.Println("expired")
        ch <- true
    })
    
    fmt.Println("start")
    <-ch
    fmt.Println("finish")
}