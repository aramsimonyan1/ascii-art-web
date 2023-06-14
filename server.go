package main

import (
	"fmt"
	"net/http"
)

func main() {

	//create a file server handler for serving static files
	fs := http.FileServer(http.Dir("./static"))

	//register the file server hadler for the "/static/" route
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
