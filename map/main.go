package main

import("fmt")

func main(){
//maps example
//key is string, value is int
	m:=map[string]int{"one":12}
	fmt.Println(m)
	fmt.Println(m["one"])
    m["one"]=1223
    fmt.Println(m["one"],m)

    delete(m,"one")
    fmt.Println(m)

}