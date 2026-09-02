package main

import (
	"net/http"

	"github.com/swaggo/swag/example/basic/api"
)

func main() {
	http.HandleFunc("/testapi/get-string-by-int/", api.GetStringByInt)
	http.HandleFunc("//testapi/get-struct-array-by-string/", api.GetStructArrayByString)
	http.HandleFunc("/testapi/upload", api.Upload)
	http.ListenAndServe(":8080", nil)
}
