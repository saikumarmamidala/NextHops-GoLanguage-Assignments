package main

import (
    "fmt"
    "log"
	"time"
)

func main() {
	c := make(chan int, 10)
	if c==nil{
		log.Print("can't allocate channel");
	}
	go producer(c)
	go consumer(c)
	time.Sleep(1 * time.Second)
}
func consumer(nums chan int){
		for v:= range nums{
			fmt.Printf("Consumer got: %v\n", v)
		}
	}
func producer(nums chan int){
		n:=10
		for i:=0;i<n;i++{
			nums<-i
			fmt.Printf("Producer :%v\n",nums)
		}
		close(nums)
	}