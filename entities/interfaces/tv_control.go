package interfaces

type TVControl interface {
	control
	IncreaseVolume()
	DecreaseVolume()
	ChangeChannel(channel int)
}
