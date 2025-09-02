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

	station.GET("/:id",func(c *gin.Context) {
		CheckScheduledByStation(c, stationService)
	})
}

func GetAllStation(c *gin.Context, service Service) {
	data, err := service.GetAllStation()

	if err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Success: false,
			Message: "Failed to get all station",
			Data: nil,
		})

		return
	}	

	c.JSON(
		http.StatusOK,
		response.APIResponse{
			Success: true,
			Message: "Successfully get all station",
			Data: data,
		},
	)
}

func CheckScheduledByStation(c *gin.Context, service Service) {
	id := c.Param("id")

	data, err := service.CheckScheduledByStation(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Success: false,
			Data: nil,
			Message: "Failed to get scheduled trains",
		})

		return
	}	

	c.JSON(
		http.StatusOK,
		response.APIResponse{
			Success: true,
			Message: "Successfully get scheduled trains ",
			Data: data,
		},
	)
}