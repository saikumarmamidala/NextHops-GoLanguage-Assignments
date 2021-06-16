package main

import (
    "fmt"
    "log"
)

func main() {
	c := make(chan int, 100)
	if c==nil{
		log.Print("can't allocate channel");
	}
	producer(c)
	consumer(c)
}
func consumer(nums chan int){
		for v:= range nums{
			fmt.Printf("Consumer got: %v\n", v)
		}
	}
func producer(nums chan int){
		n:=100
		for i:=0;i<n;i++{
			nums<-i
			fmt.Printf("Producer :%v\n",nums)
		}
		close(nums)
	}
	