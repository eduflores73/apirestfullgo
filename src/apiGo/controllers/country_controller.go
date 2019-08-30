package controllers

import (
	"github.com/gin-gonic/gin"
	"../services"
	"net/http"
)

const(
	paramCountryId = "countryId"
)

func GetCountryFromAPI(c * gin.Context){
	countryId := c.Param(paramCountryId)
	response, err := services.GetCountries(countryId)

	if err != nil{
		c.JSON(err.Status,err)
		return
	}
	c.JSON(http.StatusOK, response)
}