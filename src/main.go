package main

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
)

var history []Operation

func main() {
	fmt.Println("Running at port 8000")

	router := gin.Default()
	router.GET("/calc/sum/:num1/:num2", handlerSum)
	router.GET("/calc/sub/:num1/:num2", handlerSub)
	router.GET("/calc/mult/:num1/:num2", handlerMult)
	router.GET("/calc/div/:num1/:num2", handlerDiv)
	router.GET("/calc/history", handlerHistory)

	err := router.Run("localhost:8000")
	if err != nil {
		log.Println(err.Error())
	}
}