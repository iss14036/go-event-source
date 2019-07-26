package event

import "go-event-source/app/domain/model"

type ArrivalShipping struct {
	Occured  string
	Recorded string
	Ship     *model.Ship
	Port     string
}
