package main

import "factory-method/logistics"

func main() {
	logistic, err := logistics.New(logistics.Road)
	if err != nil {
		panic(err)
	}
	transport, err := logistic.CreateTransport(logistics.Truck)
	if err != nil {
		panic(err)
	}
	err = logistic.PlanDelivery(transport)
	if err != nil {
		panic(err)
	}
}
