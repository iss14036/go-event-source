package event

import "go-event-source/app/domain/model"

type DepartureShipping struct {
	Occured  string
	Recorded string
	Ship     *model.Ship
	Port     string
}
