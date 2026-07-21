package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/examples/go-module-support/api"
)

func main() {
	r := gin.New()
	r.GET("/testapi/get-string-by-int/:some_id", api.GetStringByInt)
	r.GET("//testapi/get-struct-array-by-string/:some_id", api.GetStructArrayByString)
	r.Run()

}
