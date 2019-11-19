package main

import (
	"fmt"
	"time"
)

func main() {
	go func() { fmt.Println("dasdasdasdasdasd") }()
	time.Sleep(1000)
	fmt.Println("Hello, World!")
}
