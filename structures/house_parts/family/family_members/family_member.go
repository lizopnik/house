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
		fmt.Print("Семейное положение: Женат\n")
	} else {
		fmt.Print("Семейное положение: Не женат\n")
	}
}

func (f FamilyMembers) WhatStatusFemale() {
	if f.FamilyStatus {
		fmt.Print("Семейное положение: Замужем\n")
	} else {
		fmt.Print("Семейное положение: Не замужем\n")
	}
}

func (f FamilyMembers) WhatSex() {
	if f.Sex {
		fmt.Print("Пол: Мужской\n")
		f.WhatStatusMale()
	} else {
		fmt.Print("Пол: Женский\n")
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
	fmt.Print("Имя: ", f.Name, "\nВозраст: ", f.Age, "\n")
	f.WhatSex()
	f.WhatChildren()
}
