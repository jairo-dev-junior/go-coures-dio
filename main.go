package main

import "fmt"

const farenheight float64 = 212.0

func main() {
	println("Hello, here we convert the water to celsius")
	var temp float64 = farenheight
	fmt.Scanf("%g", &temp)
	celsius := ((temp - 32) * 5) / 9
	fmt.Printf("The temperature to transform water in air is: %g ", celsius)
}
