package main

func main() {
	/*
	   Loop examples
	*/

	var i int
	for i < 5 {
		print(i)
		i++
		if i == 3 {
			break
		}
	}

	println(" ")
	for p := 1; p < 9; p++ {
		print(p)
	}

	println("")
	var z int
	for ; z < 9; z++ {
		print(z)
	}

	println("")

	//infinite loop
	var q int
	for {
		if q == 9 {
			break
		}
		print(q)
		q++

	}

	println("")

	// looping through a collection
	slice := []int{1, 2, 3}

	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}

	for i, v := range slice {
		println(i, v)
	}

}
