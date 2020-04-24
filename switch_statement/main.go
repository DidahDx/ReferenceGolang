package main

type HTTPRequest struct {
	Method string
}

func main() {
	r := HTTPRequest{Method: "GET"}

	//switch case example
	switch r.Method {
	case "GET":
		println("Get request")

	case "POST":
		println("POST request")
	case "DELETE":
		println("delete request")
	case "PUT":
		println("put request")
	default:
		println("unhandled method")
	}
}
