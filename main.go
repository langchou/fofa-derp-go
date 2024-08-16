package main

import (
	"fofa-derp/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	handlers.SetupRoutes(r)
	r.Run(":8080") // Listen and serve on 0.0.0.0:8080
}
