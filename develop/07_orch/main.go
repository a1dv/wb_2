package main

import (
    "fmt"
    "time"
    "sync"
)

func main() {
    sig := func(after time.Duration) <- chan interface{} {
	   c := make(chan interface{})
	   go func() {
	       defer close(c)
		   time.Sleep(after)
           fmt.Println(after, "gone")
        }()
        return c
    }

    start := time.Now()
    <-or (
    	sig(1*time.Second),
        sig(2*time.Second),
        sig(3*time.Second),
        sig(4*time.Second),
        sig(5*time.Second),
    )
    fmt.Printf("fone after %v", time.Since(start))

}

func or (channels ...<- chan interface{}) <-chan interface{} {
    var wg sync.WaitGroup
    res := make(chan interface{})
    for _, v := range channels {
        wg.Add(1)
        go func(ch <-chan interface{}) {
            for val := range ch {
                res <- val
            }
            wg.Done()
        }(v)
    }
    wg.Wait()
    close(res)
    return res
}
