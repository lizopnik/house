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

func MakeFamily() Family {
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
	family := Family{
		FamilyMembers: []family_members.FamilyMembers{neighbourDanae, neighbourSiba, cockroach, spider, batonCat},
		Surname:       "Общага",
	}
	return family
}
