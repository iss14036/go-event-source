package event

import "go-event-source/app/domain/model"

type ArrivalShipping struct {
	occured  string
	recorded string
	ship     *model.Ship
	port     string
}
