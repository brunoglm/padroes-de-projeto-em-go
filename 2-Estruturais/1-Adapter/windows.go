package main

import "fmt"

type Windows struct{}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

type WindowsAdapter struct {
	windowMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.insertIntoUSBPort()
}
