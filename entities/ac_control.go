package entities

import "github.com/ESSantana/poo-golang-style/entities/interfaces"

type ACControl struct {
	control
	temperature int
	mode        string
}

func NewACControl(id, mode string, turnOnStatus bool, temperature int) interfaces.ACControl {
	return &ACControl{
		control: control{
			id:           id,
			target:       "AC",
			turnOnStatus: turnOnStatus,
		},
		mode:        mode,
		temperature: temperature,
	}
}

func (ctrl *ACControl) IncreaseTemperature() {
	ctrl.temperature += 1
}

func (ctrl *ACControl) DecreaseTemperature() {
	ctrl.temperature -= 1
}

func (ctrl *ACControl) ChangeMode(mode string) {
	ctrl.mode = mode
}
