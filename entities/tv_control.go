package entities

import (
	"fmt"

	"github.com/ESSantana/poo-golang-style/entities/interfaces"
)

type TVControl struct {
	control
	volume  int
	channel int
}

func NewTVControl(id string, turnOnStatus bool, volume, channel int) interfaces.TVControl {
	return &TVControl{
		control: control{
			id:           id,
			target:       "TV",
			turnOnStatus: turnOnStatus,
		},
		volume:  volume,
		channel: channel,
	}
}

func (ctrl *TVControl) IncreaseVolume() {
	ctrl.volume += 1
}

func (ctrl *TVControl) DecreaseVolume() {
	ctrl.volume -= 1
}

func (ctrl *TVControl) ChangeChannel(channel int) {
	ctrl.channel = channel
}

func (ctrl *TVControl) TurnOn() {
	fmt.Println("Turning control on from tv control struct")
	ctrl.turnOnStatus = false
}
