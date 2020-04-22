package main

import "fmt"

func main(){
//using var 
var name string="golang demo"

//variable types are infered
//variables decalred must be used
var names="golang demo"

//shorthand
// age :=239876543

//assigning mulitple variables at once
age,dob:="12","123"

//constant 
const isConnect=true

//float
money:=123.33


fmt.Println(name,age,names,isConnect,money,dob)

//finding types of the variables
fmt.Printf("%T",money)


}