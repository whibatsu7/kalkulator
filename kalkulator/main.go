package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Input struct {
	Num1     float32 `json:"num1"`
	Num2     float32 `json:"num2"`
	Operator string  `json:"operator"`
}

type Status struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Result  float32 `json:"result"`
}

func Calculate(c *gin.Context) {
	var input Input
	var status Status

	if err := c.BindJSON(&input); err != nil {
		return
	} else {
		switch input.Operator {
		case "+":
			status.Result = input.Num1 + input.Num2
		case "-":
			status.Result = input.Num1 - input.Num2
		case "/":
			status.Result = input.Num1 / input.Num2
		case "*":
			status.Result = input.Num1 * input.Num2
		default:
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "Gagal hitung",
			})
			return
		}
		status.Code = 200
		status.Message = "OK"
	}
	c.IndentedJSON(http.StatusOK, status)
}

func main() {
	router := gin.Default()
	router.POST("/calculator", Calculate)
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Kalkulator nihh",
		})
	})

	router.Run("localhost:8080")
}
