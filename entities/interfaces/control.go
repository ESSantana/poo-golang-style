package interfaces

type control interface {
	TurnOn()
	TurnOff()
	GetTurnOnStatus() bool
}
