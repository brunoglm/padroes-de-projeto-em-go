package logistics

import (
	"errors"
	"fmt"
)

type TruckTransport struct{}

func (t *TruckTransport) deliver() {
	fmt.Println("Processo de entrega pelo o meio de transporte Truck")
}

type CarTransport struct{}

func (t *CarTransport) deliver() {
	fmt.Println("Processo de entrega pelo o meio de transporte Car")
}

type RoadLogistics struct{}

func (rl *RoadLogistics) PlanDelivery(transport Transport) error {
	// Realiza algumas operações de transporte por road
	transport.deliver()
	return nil
}

func (rl *RoadLogistics) CreateTransport(transportType TransportType) (Transport, error) {
	switch transportType {
	case Truck:
		return &TruckTransport{}, nil
	case Car:
		return &CarTransport{}, nil
	default:
		return nil, errors.New("tipo de transporte road incorreto")
	}
}
