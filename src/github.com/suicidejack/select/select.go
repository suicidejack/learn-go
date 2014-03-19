package main

import (
    "fmt"
    "time"
)

func fibonacci(c, quit chan int) {
    x, y := 0, 1
    for {
        // select lets a goroutine wait on multiple communication ops
        // blocks until one of the cases can run then executes that case
        // one is chosen at random if multiple are available
        select {
        case c <- x:
            x, y = y, x+y
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}

func main() {
    c := make(chan int)
    quit := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-c)
        }
        quit <- 0
    }()
    fibonacci(c, quit)

    // default case in select is run if no other case is ready
    tick := time.Tick(100 * time.Millisecond)
    boom := time.After(500 * time.Millisecond)
    for {
        select {
        case <-tick:
            fmt.Println("tick.")
        case <-boom:
            fmt.Println("BOOM")
            return
        default: //recieving blocks this
            fmt.Println("    .")
            time.Sleep(50 * time.Millisecond)
        }
    }
}
