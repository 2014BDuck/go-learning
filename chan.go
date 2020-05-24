package main

import (
	"runtime"
	"sync"
	"net/http"
)

var wg sync.WaitGroup

var urls = []string{
	"http://www.golang.org/",
	"http://www.google.com/",
	"http://www.qq.com/",
}

func main(){
	for _, url := range urls{
		wg.Add(1)

		go func(url string){
			defer wg.Done()
			resp, err := http.Get(url)
			if err == nil{
				println(resp.Status)
			}
		}(url)
	}

	wg.Wait()



	c := make(chan struct{})
	ci := make(chan int, 100)
	go func(i chan struct{}, j chan int){
		for i := 0; i < 10; i++ {
			ci <- i
		}
		close(ci)
		c <- struct{}{}
	}(c, ci)

	println(runtime.NumGoroutine())

	<-c
	println(runtime.NumGoroutine())


	for v := range(ci){
		println(v)
	}
}