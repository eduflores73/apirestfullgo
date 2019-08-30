package controllers

import (
	"../utils"
	"github.com/gin-gonic/gin"
	"../services"
	"net/http"
	"strconv"
)


func GetResultCHFromAPI(c * gin.Context){
	userId := c.Param(paramUserID)
	id, err := strconv.Atoi(userId)
	if err != nil{
		apiError := utils.ApiError{
			Messege:err.Error(),
			Status:http.StatusBadRequest,
		}
		c.JSON(apiError.Status, apiError)
		return
	}
	response, apiErr := services.GetResultCH(id)
	if apiErr != nil{
		c.JSON(apiErr.Status,apiErr)
		return
	}
	c.JSON(http.StatusOK, response)
}