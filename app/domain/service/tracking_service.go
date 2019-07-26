package service

import (
	"fmt"
	"go-event-source/app/domain/event"
	"go-event-source/app/domain/model"
)

func recordArrival(ship *model.Ship, port string) {
	ship.Location = port
	arrivalShipping := event.ArrivalShipping{Occured: "12-12-21", Recorded: "Today", Ship: ship, Port: port}
	fmt.Println(arrivalShipping)
}
