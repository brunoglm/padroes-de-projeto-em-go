package logistics

import (
	"errors"
)

type TransportMode int

const (
	Road TransportMode = iota
	Sea
)

type TransportType int

const (
	Truck TransportType = iota
	Car
	Ship
	Boat
)

type Transport interface {
	deliver()
}

type Logistic interface {
	PlanDelivery(transport Transport) error
	CreateTransport(transportType TransportType) (Transport, error)
}

func New(modeOfTransport TransportMode) (Logistic, error) {
	switch modeOfTransport {
	case Road:
		return &RoadLogistics{}, nil
	case Sea:
		return &SeaLogistics{}, nil
	default:
		return nil, errors.New("modeOfTransport incorreto")
	}
}
