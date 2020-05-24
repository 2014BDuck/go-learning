package main

import (
	"runtime"
	"time"
)

func sum(){
	sum := 0
	for i := 0; i< 10000; i++{
		sum += i
	}

	println(sum)
	time.Sleep(1 * time.Second)
}

func main(){
	println(runtime.GOMAXPROCS(0))

	go sum()

	time.Sleep(1 * time.Second)
	go func(){
		sum := 0
		for i := 0;i < 10000; i++{
			sum += i
		}
		println(sum)
		time.Sleep(1 * time.Second)
	}()

	println("NumGoroutine=", runtime.NumGoroutine())

	time.Sleep(5 * time.Second)
}