package main

import "github.com/ESSantana/poo-golang-style/entities"

func main() {
	// Testing two different types of controls
	var control1 = entities.NewTVControl("tv1", false, 50, 12)
	var control2 = entities.NewACControl("ac1", "cool", false, 20)

	// Turning both controls
	control1.TurnOn()
	control2.TurnOn()

	// Using different methods
	control1.IncreaseVolume()
	control2.IncreaseTemperature()
}
