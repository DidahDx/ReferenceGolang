package main

import "fmt"

func greeting(name string)string{
	return "Hello "+name
}

func getSum(num1,num2 int)int{
	return num1+num2
}

func getProduct(num1 int,num2 int)int{
	return num1*num2
}

func main(){
fmt.Println("Hello world")

fmt.Println(greeting("Dave"))

fmt.Println(getSum(23,32))

fmt.Println(getProduct(12,23))
}