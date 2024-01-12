package furniture

import (
	"fmt"
	"hause/structures/house_parts/types"
)

type Furniture struct {
	Type          string
	Length        types.Centimeter
	Width         types.Centimeter
	Color         string
	ComfortRating int
}

func (f Furniture) Print() {
	fmt.Print("Тип мебели: ", f.Type, "\nДлина: ", f.Length, "\nШирина: ", f.Width, "\nЦвет: ", f.Color, "\nРейтинг комфорта: ", f.ComfortRating, "\n")
}

func MakeFurniture() []Furniture {
	var furniture []Furniture
	table := Furniture{
		Type:          "Стол",
		Length:        120,
		Width:         60,
		Color:         "White",
		ComfortRating: 3,
	}
	chair := Furniture{
		Type:          "Стул",
		Length:        40,
		Width:         40,
		Color:         "Зеленый",
		ComfortRating: 1,
	}
	bed := Furniture{
		Type:          "Кровать",
		Length:        200,
		Width:         120,
		Color:         "Белый",
		ComfortRating: 4,
	}
	wardrobe := Furniture{
		Type:          "Шкаф",
		Length:        80,
		Width:         50,
		Color:         "Белый",
		ComfortRating: 5,
	}
	shelf := Furniture{
		Type:          "Полка",
		Length:        30,
		Width:         10,
		Color:         "Белый",
		ComfortRating: 1,
	}
	furniture = append(furniture, table, chair, bed, wardrobe, shelf)
	return furniture
}
