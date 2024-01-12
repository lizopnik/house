package house

import (
	"fmt"
	"hause/structures/house_parts/devices"
	"hause/structures/house_parts/family"
	"hause/structures/house_parts/furniture"
	"hause/structures/house_parts/types"
)

type House struct {
	Family        family.Family
	Devices       []devices.Device
	Furniture     []furniture.Furniture
	Rooms         int
	Square        types.SquareMeter
	CeilingHeight types.Centimeter
}

func (h House) Print() {
	fmt.Print("Комнаты: ", h.Rooms, "\nПлощадь: ", h.Square, "\nВысота потолков: ", h.CeilingHeight, "\n")
	fmt.Print("Соседи:\n")
	h.Family.Print()
	fmt.Print("Техника:\n")
	for _, tmpDevices := range h.Devices {
		tmpDevices.Print()
	}
	fmt.Print("Мебель:\n")
	for _, tmpFurniture := range h.Furniture {
		tmpFurniture.Print()
	}
}

func Make() House {
	allFurniture := furniture.MakeFurniture()
	allFamilies := family.MakeFamily()
	allDevices := devices.MakeDevices()

	house := House{
		Family:        allFamilies,
		Devices:       allDevices,
		Furniture:     allFurniture,
		Rooms:         1,
		Square:        20,
		CeilingHeight: 250,
	}
	return house
}
