package event

import "go-event-source/app/domain/model"

type DepartureShipping struct {
	occured  string
	recorded string
	ship     *model.Ship
	port     string
}
