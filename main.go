// main.go
package main

import (
	"WiratKhamphan/SHOPLEK/views"

	"github.com/gin-gonic/gin"

	"net/http"
)

func main() {
	router := gin.Default()
	views.WD(router)
	http.ListenAndServe(":8080", router)
}
