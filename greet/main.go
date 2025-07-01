package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}

func displayGreet(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Elodie")
}

func main() {
	Greet(os.Stdout, "Elodie")
	http.ListenAndServe(":8080", http.HandlerFunc(displayGreet))
}
