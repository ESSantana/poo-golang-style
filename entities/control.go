package entities

import "fmt"

type control struct {
	id           string
	target       string
	turnOnStatus bool
}

func (ctrl *control) TurnOn() {
	fmt.Println("Turning control on")
	ctrl.turnOnStatus = true
}

func (ctrl *control) TurnOff() {
	fmt.Println("Turning control off")
	ctrl.turnOnStatus = false
}

func (ctrl *control) GetTurnOnStatus() bool {
	return ctrl.turnOnStatus
}
