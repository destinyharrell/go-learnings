package main

//import "fmt" //Go's format library
//
//func main() {
//	fmt.Println("Hello, in english only")
//}

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.Now().UnixNano())  // You must seed your ruandom numbers in Go
	fmt.Println("My fav number: ", rand.Intn(10))
}