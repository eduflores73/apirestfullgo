package main

import (
	"../controllers"
	"github.com/gin-gonic/gin"
	"log"
)

const (
	port = ":8080"
)

var (
	router = gin.Default()
)

func main(){

	router.GET("/users/:userId", controllers.GetUserFromAPI)
	router.GET("/countries/:countryId", controllers.GetCountryFromAPI)
	router.GET("/sites/:siteId", controllers.GetSiteFromAPI)
	router.GET("/results/:userId", controllers.GetResultFromAPI)
	router.GET("/resultsGW/:userId", controllers.GetResultWGFromAPI)
	router.GET("/resultsCH/:userId", controllers.GetResultCHFromAPI)

	log.Fatal(router.Run(port))



}
