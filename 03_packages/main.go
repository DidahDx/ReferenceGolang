package main

//importing multiple packages
import (
	"fmt"
    "math"
    "github.com/DidahDx/go_demo_app/03_packages/strutil"
)

func main(){

//used to round down
fmt.Println(math.Floor(2.74))

fmt.Println(math.Ceil(2.74))

fmt.Println(math.Sqrt(2.74))

//using another package
fmt.Println(strutil.Reverse("qwertyuiop"))
}