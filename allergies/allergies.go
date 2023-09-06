package allergies

var allergens = []string{
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

	allergies = allergies % 256

	for i := 0; i < len(allergens); i++ {
        if allergies & (1 << uint(i)) != 0 {
            res = append(res, allergens[i])
        }
	}

	return res
}

func AllergicTo(allergies uint, allergen string) bool {
    for i, item := range allergens {
        if item == allergen {
            return allergies & (1 << uint(i)) != 0
        }
    }

    return false
}
