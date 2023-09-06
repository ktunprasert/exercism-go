package allergies

var allergens = [8]string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

func Allergies(allergies uint) []string {
	res := []string{}

	for i := 0; i < len(allergens); i++ {
		if AllergicTo(allergies, allergens[i]) {
			res = append(res, allergens[i])
		}
	}

	return res
}

func AllergicTo(allergies uint, allergen string) bool {
	switch allergen {
	case "eggs":
		return allergies&1 != 0
	case "peanuts":
		return allergies&2 != 0
	case "shellfish":
		return allergies&4 != 0
	case "strawberries":
		return allergies&8 != 0
	case "tomatoes":
		return allergies&16 != 0
	case "chocolate":
		return allergies&32 != 0
	case "pollen":
		return allergies&64 != 0
	case "cats":
		return allergies&128 != 0
	}

	return false
}
