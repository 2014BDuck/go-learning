package main

import (
	"fmt"
	"context"
	"time"
)

// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
// 	defer cancel()
// 	res_ch := make(chan int)
// 	go real_processing(res_ch)
// 	go handle(ctx, res_ch)
// 	// status_ch<-1
// 	select {
// 	case <-ctx.Done():
// 		fmt.Println("main", ctx.Err())
// 	}
// }

// func handle(ctx context.Context, res_ch <-chan int) {
// 	select {
// 	case <-ctx.Done():
// 		fmt.Println("handle", ctx.Err())
// 	case v := <-res_ch:
// 		fmt.Println("Calculation done, result: ",v)
// 	}
// }

// func real_processing(res_ch chan<- int){
// 	sum := 0
// 	for i:= 0; i < 1000000000; i++{
// 		sum += i
// 	}
// 	res_ch<-sum
// 	return
// }


func main(){
	ctxa, cancel := context.WithCancel(context.Background())

	go work(ctxa, "work1")

	time.Sleep(1 * time.Second)
	cancel()
}




func work(ctx context.Context, name string){
	for {
		select{
		case <-ctx.Done():
			fmt.Println("Work Expired")
		default:
			fmt.Println(name)
			time.Sleep(100000 * time.Microsecond)
		}
	}
}





