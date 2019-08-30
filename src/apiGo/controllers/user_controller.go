package controllers

import (
	"../utils"
	"github.com/gin-gonic/gin"
	"../services"
	"net/http"
	"strconv"
	"fmt"
)

const(
	paramUserID = "userId"
)

func GetUserFromAPI(c * gin.Context){
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

	response, apiErr := services.GetUsers(id)
	fmt.Println(response)
	if err != nil{
		c.JSON(apiErr.Status,err)
		return
	}
	c.JSON(http.StatusOK, response)
}