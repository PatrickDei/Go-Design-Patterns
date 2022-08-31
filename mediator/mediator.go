package mediator

import "fmt"

func ShowcaseMediator() {
	fmt.Println("\nMediator pattern")

	stationManager := newStationManager()

	passengerTrain := &PassengerTrain{
		mediator: stationManager,
	}
	freightTrain := &CargoTrain{
		mediator: stationManager,
	}

	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()
}
