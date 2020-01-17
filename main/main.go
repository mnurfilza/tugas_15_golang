package main

import (
	"belajargolang/tugas15"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/index", tugas15.Tugas15)
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
