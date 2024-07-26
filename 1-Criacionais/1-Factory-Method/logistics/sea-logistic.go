package logistics

import (
	"errors"
	"fmt"
)

type ShipTransport struct{}

func (t *ShipTransport) deliver() {
	fmt.Println("Processo de entrega pelo o meio de transporte Ship")
}

type BoatTransport struct{}

func (t *BoatTransport) deliver() {
	fmt.Println("Processo de entrega pelo o meio de transporte Boat")
}

type SeaLogistics struct{}

func (rl *SeaLogistics) PlanDelivery(transport Transport) error {
	// Realiza algumas operações de transporte por sea
	transport.deliver()
	return nil
}

func (rl *SeaLogistics) CreateTransport(transportType TransportType) (Transport, error) {
	switch transportType {
	case Ship:
		return &ShipTransport{}, nil
	case Boat:
		return &BoatTransport{}, nil
	default:
		return nil, errors.New("tipo de transporte sea incorreto")
	}
}
