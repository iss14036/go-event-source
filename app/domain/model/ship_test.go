package model

import "testing"

func TestGetName(t *testing.T) {
	ship := Ship{Name: "Mug Glass", Location: "Islandia"}
	want := "Mug Glass"
	if got := ship.GetName(); got != want {
		t.Errorf("ship.GetName() = %q, want %q", got, want)
	}
}

func TestGetLocation(t *testing.T) {
	ship := Ship{Name: "Mug Glass", Location: "Islandia"}
	want := "Islandia"
	if got := ship.GetLocation(); got != want {
		t.Errorf("Ship.GetLocation() = %q, want %q", got, want)
	}
}

func TestString(t *testing.T) {
	ship := Ship{Name: "Mug Glass", Location: "Islandia"}
	want := "(" + ship.Name + " " + ship.Location + ")"
	if got := ship.String(); got != want {
		t.Errorf("Ship.GetLocation() = %q, want %q", got, want)
	}
}
