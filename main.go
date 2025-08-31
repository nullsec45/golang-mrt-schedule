package main 

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"github.com/nullsec45/golang-mrt-schedule/modules/station"
)

func InitiateRouter(){
	var router = gin.Default()
	var api = router.Group("/api/v1")

	station.Initiate(api)
	router.Run(":9000")
}

func main(){
	InitiateRouter()
}

