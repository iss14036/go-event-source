package main

import (
	"fmt"
	"go-event-source/app/domain/model"
	"go-event-source/app/domain/service"
)

func main() {
	trackingService := service.TrackingService{}
	ship := model.Ship{Name: "Mug Glass", Location: "Start"}
	trackingService.RecordArrival(&ship, "London")
	trackingService.RecordArrival(&ship, "Las Vegas")
	trackingService.RecordArrival(&ship, "Texas")
	trackingService.RecordArrival(&ship, "Kansas")
	fmt.Println(ship.String())
}
