package main

import "fmt"

func main(){
fmt.Println("Hello world")

//arrays
var arr[5] int
arr[0]=1
arr[1]=2

fmt.Println(arr)

rest:=[3]int{2,3,4}

fmt.Println(rest)

//slice initialising with array
asd:=arr[:]
fmt.Println(asd)


slice:=[]int{1,2,3}
slice=append(slice,4,5,6)

s1:=slice[1:]
s2:=slice[1:2]
s3:=slice[:2]

fmt.Println(slice,s1,s2,s3)

}