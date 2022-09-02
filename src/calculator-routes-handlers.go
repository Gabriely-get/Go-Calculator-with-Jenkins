package main

import (
	"fmt"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
)

func handlerSum(c *gin.Context) {
	floatNum1, err := convertToFloat( c.Param("num1") )
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, returnError("Could not converter value"))
	}

	floatNum2, err := convertToFloat( c.Param("num2") )
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, returnError("Could not converter value"))
	}

	sumResult := sum(floatNum1, floatNum2)

	newOperation := Operation {
		Time: time.Now(),
		Operation: "ADDITION",
		Result: fmt.Sprintf("%.2f + %.2f = %.2f", floatNum1, floatNum2, sumResult),
 	}
	addToHistory(newOperation)

	c.IndentedJSON(http.StatusOK, returnSuccess(sumResult))
}

func handlerSub(c *gin.Context) {
	floatNum1, err := convertToFloat( c.Param("num1") )
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, returnError("Could not convert value"))
	}

	floatNum2, err := convertToFloat( c.Param("num2") )
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, returnError("Could not convert value"))
	}

	subResult := sub(floatNum1, floatNum2)

	newOperation := Operation {
		Time: time.Now(),
		Operation: "SUBTRACTION",
		Result: fmt.Sprintf("%.2f - %.2f = %.2f", floatNum1, floatNum2, subResult),
 	}
	addToHistory(newOperation)

	c.IndentedJSON(http.StatusOK, returnSuccess(subResult))
}

func handlerMult(c *gin.Context) {
	floatNum1, err := convertToFloat( c.Param("num1") )
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, returnError("Could not convert value"))
	}

	floatNum2, err := convertToFloat( c.Param("num2") )
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, returnError("Could not convert value"))
	}

	multResult := mult(floatNum1, floatNum2)

	newOperation := Operation {
		Time: time.Now(),
		Operation: "MULTIPLICATION",
		Result: fmt.Sprintf("%.2f * %.2f = %.2f", floatNum1, floatNum2, multResult),
 	}
	addToHistory(newOperation)

	c.IndentedJSON(http.StatusOK, returnSuccess(multResult))
}

func handlerDiv(c *gin.Context) {
	floatNum1, err := convertToFloat( c.Param("num1") )
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, returnError("Could not convert value"))
	}

	floatNum2, err := convertToFloat( c.Param("num2") )
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, returnError("Could not convert value"))
	}

	if floatNum2 == 0 {
		c.IndentedJSON(http.StatusBadRequest, returnError("An error occurred on division by 0"))
	}

	divResult := div(floatNum1, floatNum2)

	newOperation := Operation {
		Time: time.Now(),
		Operation: "DIVISION",
		Result: fmt.Sprintf("%.2f / %.2f = %.2f", floatNum1, floatNum2, divResult),
 	}
	addToHistory(newOperation)

	c.IndentedJSON(http.StatusOK, returnSuccess(divResult))
}

func handlerHistory(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, history)
}