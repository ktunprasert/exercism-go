// Triangle is a package for classifying triangles.
package triangle

type Kind string

const (
	NaT = "NaT" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
)

// KindFromSides returns the kind of triangle given the lengths of its sides.
// It returns NaT if the sides cannot form a triangle.
// It returns Equ if the triangle is equilateral.
// It returns Iso if the triangle is isosceles.
// It returns Sca if the triangle is scalene.
func KindFromSides(a, b, c float64) Kind {
	switch {
	case a <= 0 || b <= 0 || c <= 0, a+b < c || a+c < b || b+c < a:
		return NaT
	case a == b && b == c:
		return Equ
	case a == b || b == c || a == c:
		return Iso
	default:
		return Sca
	}
}
