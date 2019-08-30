package controllers

import (
	"github.com/gin-gonic/gin"
	"../services"
	"net/http"
)

const(
	paramSiteID = "siteId"
)

func GetSiteFromAPI(c * gin.Context){
	siteId := c.Param(paramSiteID)
	response, err := services.GetSites(siteId)

	if err != nil{
		c.JSON(err.Status,err)
		return
	}
	c.JSON(http.StatusOK, response)
}