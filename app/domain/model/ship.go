package model

type Ship struct {
	name     string
	location string
}

func (u *Ship) GetName() string {
	return u.name
}

func (u *Ship) GetLocation() string {
	return u.location
}
