package main

//creating a struct
type User struct {
	ID         int
	FirstName  string
	secondName string
}

func main() {

	u1 := User{
		ID:         1,
		FirstName:  "John",
		secondName: "Test",
	}

	u2 := User{
		ID:         2,
		FirstName:  "John",
		secondName: "Test",
	}

	if u1.ID == u2.ID {
		println("they are same")
	} else {
		println("not the same")
	}

}
