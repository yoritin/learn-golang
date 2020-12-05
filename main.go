package main

import (
  "fmt"
  "time"
	"math/rand"
	// "./hello"
	// "./functions"
)

func main() {
	randnum()
	// hello.Hello()
	// fmt.Println(functions.Add(34, 678))
}

func randnum() {
  rand.Seed(time.Now().UnixNano())
  fmt.Println("My favorite number is", rand.Intn(10))
}
