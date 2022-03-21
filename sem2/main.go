package main

import "fmt"

var a = b + c
var b = f()
var c = 1

func f() int { return c + 1 }

func main() {
	fmt.Println(a, b, c)
}
