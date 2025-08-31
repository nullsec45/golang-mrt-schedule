package station

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/nullsec45/golang-mrt-schedule/common/response"
)

func Initiate(router *gin.RouterGroup){
	stationService := NewService()

	station := router.Group("/stations")
	station.GET("/", func (c *gin.Context) {
		GetAllStation(c, stationService)
	})
}

func GetAllStation(c *gin.Context, service Service) {
	data, err := service.GetAllStation()

	if err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Success: false,
			Message: "Failed to fetch stations",
			Data: nil,
		})

		return
	}	

	c.JSON(
		http.StatusOK,
		response.APIResponse{
			Success: true,
			Message: "Stations fetched successfully",
			Data: data,
		},
	)
}