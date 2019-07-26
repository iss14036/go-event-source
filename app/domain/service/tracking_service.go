package service

import (
	"fmt"
	"go-event-source/app/domain/event"
	"go-event-source/app/domain/model"
)

type TrackingService struct{}

func (self *TrackingService) RecordArrival(ship *model.Ship, port string) {
	ship.Location = port
	arrivalShipping := event.ArrivalShipping{Occured: "12-12-21", Recorded: "Today", Ship: ship, Port: port}
	fmt.Println(arrivalShipping)
}

func (self *TrackingService) RecordDeparture(ship *model.Ship, port string) {
	ship.Location = port
	departureShipping := event.DepartureShipping{Occured: "12-12-21", Recorded: "Today", Ship: ship, Port: port}
	fmt.Println(departureShipping)
}
