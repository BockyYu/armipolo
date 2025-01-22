package main

import (
	routes "armipolo/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.InitRouter(r)
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
