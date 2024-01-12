package devices

import (
	"fmt"
	"hause/structures/house_parts/types"
)

type Device struct {
	Type         string
	Length       types.Centimeter
	Width        types.Centimeter
	Color        string
	VoiceControl bool
}

func (d Device) Voice() {
	if d.VoiceControl {
		fmt.Print("Voice control: Да\n")
	} else {
		fmt.Print("Voice control: Нет\n")
	}
}

func (d Device) Print() {
	fmt.Print("Тип: ", d.Type, ", Длина: ", d.Length, ", Ширина: ", d.Width, ", Цвет: ", d.Color, ", ")
	d.Voice()
}

func MakeDevices() []Device {
	var devices []Device
	laptop := Device{
		Type:         "Ноутбук",
		Length:       35,
		Width:        20,
		Color:        "Черный",
		VoiceControl: false,
	}
	phone := Device{
		Type:         "Телефон",
		Length:       15,
		Width:        7,
		Color:        "Черный",
		VoiceControl: true,
	}
	pelletMachine := Device{
		Type:         "Машинка для катышков",
		Length:       20,
		Width:        4,
		Color:        "Черный",
		VoiceControl: false,
	}
	watch := Device{
		Type:         "Часы",
		Length:       4,
		Width:        1,
		Color:        "Черный",
		VoiceControl: false,
	}
	epilator := Device{
		Type:         "Эпилятор",
		Length:       10,
		Width:        6,
		Color:        "Белый",
		VoiceControl: false,
	}
	devices = append(devices, laptop, phone, pelletMachine, watch, epilator)
	return devices
}
