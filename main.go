package main

import (
	todo "go-todo/app/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	todo.Test2()
	r.Group("/api")
	{

	}
}
