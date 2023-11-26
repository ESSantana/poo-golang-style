package interfaces

type ACControl interface {
	control
	IncreaseTemperature()
	DecreaseTemperature()
	ChangeMode(mode string)
}
