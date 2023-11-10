package devices

import (
	"fmt"
	"hause/structures/house_parts/types"
)

type Devices struct {
	Type         string
	Length       types.Centimeter
	Width        types.Centimeter
	Color        string
	VoiceControl bool
}

func (d Devices) Print() {
	fmt.Print("Тип устройства: ", d.Type, "\nДлина устройства: ", d.Length, "\nШирина устройства: ", d.Width, "\nЦвет: ", d.Color, "\n")
	if d.VoiceControl {
		fmt.Print("Управление голосом: Да\n")
	} else {
		fmt.Print("Управление голосом: Нет\n")
	}
}
