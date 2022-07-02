package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	fmt.Println(http.ListenAndServe(":8000", nil))
}
func HelloServer(w http.ResponseWriter, r *http.Request) {
	name := "Hello world1"
	fmt.Fprintf(w, name)

}
