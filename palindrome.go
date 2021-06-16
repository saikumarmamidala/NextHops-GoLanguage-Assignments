package main

import "fmt"

func main() {
	s := "saias"
	revs :=""
	for i := len(s)-1; i >=0; i-- {
		revs=revs+string(s[i])
	}
	if s==revs{
		fmt.Println("Palindrome")
	}else{
		fmt.Println("Not Palindrome")
	}
	
}
