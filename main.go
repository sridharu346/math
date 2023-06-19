package main

import (
	"fmt"
	"net/http"
)

func main() {

	//hosting a server
	fmt.Println("Local host is Served at port 8808")
	err := http.ListenAndServe(":8808", nil)
	if err != nil {
		fmt.Println("Error in hosting the server", err)
		return
	}
}
