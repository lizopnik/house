package family

import (
	"fmt"
	"hause/structures/house_parts/family/family_members"
)

type Family struct {
	FamilyMembers []family_members.FamilyMembers
	Surname       string
}

func (f Family) Print() {
	fmt.Print("Семья: ", f.Surname, "\n")
	for _, familyMembers := range f.FamilyMembers {
		familyMembers.Print()
	}
}
