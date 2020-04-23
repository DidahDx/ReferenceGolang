package main

import ("fmt")

func main(){
	//creating a struct
	type user struct{
		ID int
		FirstName string 
		secondName string
	}

//assigning it to a variable
	var u user
	fmt.Println(u)

	u.ID=1
	u.FirstName="kjhgf"
	u.secondName="asdf"
    fmt.Println(u)

//end with a comma or the compiler will add a semicolon
u2:=user{
	ID:1,
	FirstName:"John",
	secondName:"Test",
   }

fmt.Println(u2)


}