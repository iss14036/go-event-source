package model

type Ship struct {
	Name     string
	Location string
}

func (u *Ship) GetName() string {
	return u.Name
}

func (u *Ship) GetLocation() string {
	return u.Location
}
