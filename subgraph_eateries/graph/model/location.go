package model

type Location struct {
	ID                  string    `json:"id"`
	Latitude            *float64  `json:"latitude"`
	Longitude           *float64  `json:"longitude"`
	EateriesForLocation []*Eatery `json:"eateriesForLocation"`
}

func (Location) IsEntity() {}
