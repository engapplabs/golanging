package main 

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)

	http.ListenAndServe(":8080", nil)

}

func aboutHandler(response http.ResponseWriter,
					request *http.Request) {
	fmt.Fprintf(response, "<h1>hah</h1>")
}

func indexHandler(response http.ResponseWriter, 
					resquest *http.Request) {
	fmt.Fprintf(response, "GO eh top papai")
}

