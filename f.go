package main

import "fmt"

func main() {
	var message string
	message = "Hello bro"

	var name string
	var age int
	var weight float32
	var isAdult bool
	name = "Dauren"
	weight = 75.5
	age = 34
	const goldenRatio = 1.618

	fmt.Println(message)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(weight)
	fmt.Println(isAdult)
	personInfo(name, age, weight, isAdult)
}

func personInfo(name string, age int, weight float32, isAdult bool) {
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("My name is %s", name)
}
