package main

import "fmt"
import "errors"

func greeting(name string)string{
	return "Hello "+name
}

func getSum(num1,num2 int)int{
	return num1+num2
}

func getProduct(num1 int,num2 int)int{
	return num1*num2
}

//returning multiple variables
func getDivide(num1 ,num2 int)(int ,error){
return num1/num2,errors.New("Something went wrong")
}

func main(){
fmt.Println("Hello world")

fmt.Println(greeting("Dave"))

fmt.Println(getSum(23,32))

fmt.Println(getProduct(12,23))

fmt.Println(getDivide(12,4))

}