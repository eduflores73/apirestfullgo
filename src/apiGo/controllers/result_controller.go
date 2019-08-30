package controllers

import (
	"../utils"
	"github.com/gin-gonic/gin"
	"../services"
	"net/http"
	"strconv"
	"fmt"
)


func GetResultFromAPI(c * gin.Context){
	userId := c.Param(paramUserID)
	cont := 0
	id, err := strconv.Atoi(userId)
	if err != nil{
		apiError := utils.ApiError{
			Messege:err.Error(),
			Status:http.StatusBadRequest,
		}
		c.JSON(apiError.Status, apiError)
		return
	}
	cb := utils.CircuitBreaker{}

	if cb.State() == utils.Open {
		fmt.Println("LLegaste al maximo de errores, intente nuevamente mas tarde")
		//response, apiErr := services.GetResult(id)
		/*if err != nil{
			c.JSON(apiErr.Status,err)
			return
		}
		c.JSON(http.StatusOK, response)*/
	}
	if cb.State() == utils.Close {
		response, apiErr := services.GetResult(id)
		if err != nil{
			cont++
			if utils.ContadorError(cont){
				c.JSON(apiErr.Status,err)
				return
			}
		}
		c.JSON(http.StatusOK, response)
	}


}


