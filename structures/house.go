package house

import (
	"fmt"
	"hause/structures/house_parts/devices"
	"hause/structures/house_parts/family"
	"hause/structures/house_parts/family/family_members"
	"hause/structures/house_parts/furniture"
	"hause/structures/house_parts/types"
)

type House struct {
	Family        []family.Family
	Devices       []devices.Devices
	Furniture     []furniture.Furniture
	Rooms         int
	Square        types.SquareMeter
	CeilingHeight types.Centimeter
}

func (h House) Print() {
	fmt.Print("Комнаты: ", h.Rooms, "\nПлощадь: ", h.Square, "\nВысота потолков: ", h.CeilingHeight, "\n")
	for _, families := range h.Family {
		families.Print()
	}
	for _, tmpDevices := range h.Devices {
		tmpDevices.Print()
	}
	for _, tmpFurniture := range h.Furniture {
		tmpFurniture.Print()
	}
}
func Make() House {
	table := furniture.Furniture{
		Type:          "Стол",
		Length:        120,
		Width:         60,
		Color:         "White",
		ComfortRating: 3,
	}
	chair := furniture.Furniture{
		Type:          "Стул",
		Length:        40,
		Width:         40,
		Color:         "Зеленый",
		ComfortRating: 1,
	}
	bed := furniture.Furniture{
		Type:          "Кровать",
		Length:        200,
		Width:         120,
		Color:         "Белый",
		ComfortRating: 4,
	}
	wardrobe := furniture.Furniture{
		Type:          "Шкаф",
		Length:        80,
		Width:         50,
		Color:         "Белый",
		ComfortRating: 5,
	}
	shelf := furniture.Furniture{
		Type:          "Полка",
		Length:        30,
		Width:         10,
		Color:         "Белый",
		ComfortRating: 1,
	}

	neighbourSiba := family_members.FamilyMembers{
		Sex:          false,
		Age:          27,
		Name:         "Сиба",
		FamilyStatus: false,
		Children:     0,
	}
	neighbourDanae := family_members.FamilyMembers{
		Sex:          false,
		Age:          20,
		Name:         "Данае",
		FamilyStatus: false,
		Children:     0,
	}
	cockroach := family_members.FamilyMembers{
		Sex:          true,
		Age:          0,
		Name:         "Володя",
		FamilyStatus: true,
		Children:     324,
	}
	spider := family_members.FamilyMembers{
		Sex:          false,
		Age:          1,
		Name:         "Гульчатай",
		FamilyStatus: false,
		Children:     3,
	}
	batonCat := family_members.FamilyMembers{
		Sex:          true,
		Age:          1,
		Name:         "Платошка",
		FamilyStatus: false,
		Children:     0,
	}

	laptop := devices.Devices{
		Type:         "Ноутбук",
		Length:       35,
		Width:        20,
		Color:        "Черный",
		VoiceControl: false,
	}
	phone := devices.Devices{
		Type:         "Телефон",
		Length:       15,
		Width:        7,
		Color:        "Черный",
		VoiceControl: true,
	}
	pelletMachine := devices.Devices{
		Type:         "Машинка для катышков",
		Length:       20,
		Width:        4,
		Color:        "Черный",
		VoiceControl: false,
	}
	watch := devices.Devices{
		Type:         "Часы",
		Length:       4,
		Width:        1,
		Color:        "Черный",
		VoiceControl: false,
	}
	epilator := devices.Devices{
		Type:         "Эпилятор",
		Length:       10,
		Width:        6,
		Color:        "Белый",
		VoiceControl: false,
	}
	house := House{
		Family: []family.Family{family.Family{
			FamilyMembers: []family_members.FamilyMembers{neighbourDanae, neighbourSiba, cockroach, spider, batonCat},
			Surname:       "Общага",
		},
		},
		Devices:       []devices.Devices{laptop, phone, pelletMachine, watch, epilator},
		Furniture:     []furniture.Furniture{table, chair, wardrobe, bed, shelf},
		Rooms:         1,
		Square:        20,
		CeilingHeight: 250,
	}
	return house
}
