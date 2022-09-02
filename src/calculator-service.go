package main

import (
	"log"
	"strconv"
)

func addToHistory(newOperation Operation) []Operation {
	history = append(history, newOperation)
	return history
}

func convertToFloat(num string) (float64, error) {
	floatNum, err := strconv.ParseFloat(num, 64)

	if err != nil {
		log.Fatalf("error: %v", err)
		return 0,err
	}
	return floatNum,nil
}

func returnError(message string) ResponseError {
	responseError := ResponseError {
		Message: message,
	}

	return responseError
}

func returnSuccess(result float64) ResponseSuccess {
	responseSuccess := ResponseSuccess {
		Result: result,
	}

	return responseSuccess
}

func sub(num1 float64, num2 float64) float64 {
	return num1 - num2;
}

func mult(num1 float64, num2 float64) float64 {
	return num1 * num2;
}

func sum(num1 float64, num2 float64) float64 {
	return num1 + num2;
}

func div(num1 float64, num2 float64) float64 {
	return num1 / num2;
}