package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/testapi/update-product", UpdateProduct)
	http.ListenAndServe(":8080", nil)
}
