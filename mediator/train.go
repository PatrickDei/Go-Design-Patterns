package mediator

import "fmt"

type Train interface {
	arrive()
	depart()
	permitArrival()
}

type PassengerTrain struct {
	mediator Mediator
}

func (pt *PassengerTrain) arrive() {
	if !pt.mediator.canArrive(pt) {
		fmt.Println("PassengerTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("PassengerTrain: Arrived")
}

func (pt *PassengerTrain) depart() {
	fmt.Println("PassengerTrain: Leaving")
	pt.mediator.notifyAboutDeparture()
}

func (pt *PassengerTrain) permitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving")
	pt.arrive()
}

type CargoTrain struct {
	mediator Mediator
}

func (ct *CargoTrain) arrive() {
	if !ct.mediator.canArrive(ct) {
		fmt.Println("CargoTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("CargoTrain: Arrived")
}

func (ct *CargoTrain) depart() {
	fmt.Println("CargoTrain: Leaving")
	ct.mediator.notifyAboutDeparture()
}

func (ct *CargoTrain) permitArrival() {
	fmt.Println("CargoTrain: Arrival permitted, arriving")
	ct.arrive()
}
