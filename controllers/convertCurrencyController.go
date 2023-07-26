package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"AsiaYo_BackendEngineer/service"
)

func ConvertCurrency(c *gin.Context) {
	var params struct {
		Source string `form:"source" binding:"required"`
		Target string `form:"target" binding:"required"`
		Amount string `form:"amount" binding:"required"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	params.Amount = strings.TrimPrefix(params.Amount, "$")
	params.Amount = strings.ReplaceAll(params.Amount, ",", "")
	convertedAmount, err := strconv.ParseFloat(params.Amount, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	amount, err := service.Convert(params.Source, params.Target, convertedAmount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":    "success",
		"amount": amount,
	})
}
