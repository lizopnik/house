package family_members

import "fmt"

type FamilyMembers struct {
	Sex          bool
	Age          int
	Name         string
	FamilyStatus bool
	Children     int
}

func (f FamilyMembers) WhatStatusMale() {
	if f.FamilyStatus {
		fmt.Print("Семейное положение: Женат, ")
	} else {
		fmt.Print("Семейное положение: Не женат, ")
	}
}

func (f FamilyMembers) WhatStatusFemale() {
	if f.FamilyStatus {
		fmt.Print("Семейное положение: Замужем, ")
	} else {
		fmt.Print("Семейное положение: Не замужем, ")
	}
}

func (f FamilyMembers) WhatSex() {
	if f.Sex {
		fmt.Print("Пол: Мужской, ")
		f.WhatStatusMale()
	} else {
		fmt.Print("Пол: Женский, ")
		f.WhatStatusFemale()
	}
}

func (f FamilyMembers) WhatChildren() {
	if f.Children == 0 {
		fmt.Print("Дети: Нет\n")
	} else {
		fmt.Print("Дети: ", f.Children, "\n")
	}
}

func (f FamilyMembers) Print() {
	fmt.Print("Имя: ", f.Name, ", Возраст: ", f.Age, ", ")
	f.WhatSex()
	f.WhatChildren()
}
