package station 

import (
	"encoding/json"
	"time"
	"net/http"
	"github.com/nullsec45/golang-mrt-schedule/common/client"
)

type Service interface {
	GetAllStation() (response []StationResponse, err error)
}

type service struct {
	client *http.Client
}

func NewService() Service {
	return &service{
		client:&http.Client{
			Timeout: 10 * time.Second,
		},
	}
}


func (s *service) GetAllStation() (response []StationResponse, err error){
	url := "https://jakartamrt.co.id/id/val/stasiuns"

	byteResponse, err := client.DoRequest(s.client,url)											

	if err != nil {
		return 
	}

	var stations []Station
	err =json.Unmarshal(byteResponse, &stations)

	for _, item := range stations {
		response = append(response, StationResponse{
			Id: item.Id,
			Name: item.Name,
		})
	}

	// Implement the logic to fetch and return all stations
	return response, nil
}