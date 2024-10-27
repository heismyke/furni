package main

import (
	"fmt"
	"net/http"

	"github.com/heismyke/furni/controllers"
)

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello", controllers.HandleHello)
	server.HandleFunc("/", controllers.HandleTemplate)

    
	err := http.ListenAndServe(":8080", server)

	if err != nil {
		fmt.Println("error while opening server")
	}

}
