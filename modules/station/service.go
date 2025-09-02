package station 

import (
	"encoding/json"
	"time"
	"net/http"
	"github.com/nullsec45/golang-mrt-schedule/common/client"
	"errors"
	"strings"
	// "fmt"
)

type Service interface {
	GetAllStation() (response []StationResponse, err error)
	CheckScheduledByStation(stationId string) (response []ScheduleResponse, err error)
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

	return response, nil
}

func (s *service) CheckScheduledByStation(stationId string) (response []ScheduleResponse, err error){
	url := "https://jakartamrt.co.id/id/val/stasiuns"
	
	byteResponse, err := client.DoRequest(s.client,url)											

	if err != nil {
		return 
	}

	var schedule []Schedule
	err =json.Unmarshal(byteResponse, &schedule)
	if err != nil {
		return 
	}

	// fmt.Printf("Schedules: %+v\n", schedule)

	// fmt.Println(schedule)
	var scheduledSelected Schedule
	for _, item := range schedule {
		// fmt.Println(item.StationId, stationId)
		if item.StationId == stationId {
			scheduledSelected = item
			break
		}
	}

	// fmt.Printf("Schedule struct: %+v\n", scheduledSelected)


	if scheduledSelected.StationId == "" {
		err = errors.New("station not found")
		return
	}	
	
	response, err = ConvertDataToResponses(scheduledSelected)
	if err != nil {
		return 
	}

	return response, nil
}	

func ConvertDataToResponses(schedule Schedule) (response []ScheduleResponse, err error) {
	// fmt.Println("schedule", schedule)

	// fmt.Printf("Schedule struct: %+v\n", schedule)
	var (
		LebakBulusTripName = "Stasiun Lebak Bulus Grab"
		BundaranHITripName = "Stasiun Bundaran HI Bank DKI" 
	)

	scheduleLebakBulus := schedule.ScheduleLebakBulus
	scheduleBundaranHI := schedule.ScheduleBundaranHI

	// fmt.Println("scheduleLebakBulus", scheduleLebakBulus)

	scheduleLebakBulusParsed, err := ConvertScheduleToTimeFormat(scheduleLebakBulus)
	if err != nil {
		return 
	}

	scheduleBundaranHIParsed, err := ConvertScheduleToTimeFormat(scheduleBundaranHI)
	if err != nil {
		return 
	}

	for _, item := range scheduleLebakBulusParsed {
		if item.Format("15:04") > time.Now().Format("15:04") {
			response=append(response, ScheduleResponse{
				StationName: LebakBulusTripName,
				Time: item.Format("15:04"),
			})
		}
	}

	for _, item := range scheduleBundaranHIParsed {
		if item.Format("15:04") > time.Now().Format("15:04") {
			response=append(response, ScheduleResponse{
				StationName: BundaranHITripName,
				Time: item.Format("15:04"),
			})
		}
	}

	return response, nil
}

func ConvertScheduleToTimeFormat(schedule string) (response []time.Time, err error) {
	// fmt.Println(schedule)

	var (
		parsedTime time.Time
		schedules = strings.Split(schedule, ",")
	)

	for _, item := range schedules {
		trimmedTime := strings.TrimSpace(item)
		if trimmedTime	== "" {
			continue
		}

		parsedTime, err = time.Parse("15:04", trimmedTime)
		if err != nil {
			err = errors.New("Invalid time format "+trimmedTime)
			return 
		}

		response = append(response, parsedTime)
	}

	// fmt.Println(response)

	return response, nil
}