package complexnumbers

import "math"

type Number struct {
	a, b float64
}

func (n Number) Real() float64 {
	return n.a
}

func (n Number) Imaginary() float64 {
	return n.b
}

func (n1 Number) Add(n2 Number) Number {
	return Number{
		a: n1.a + n2.a,
		b: n1.b + n2.b,
	}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{
		a: n1.a - n2.a,
		b: n1.b - n2.b,
	}
}

func (n1 Number) Multiply(n2 Number) Number {
	return Number{
		a: n1.a*n2.a - n1.b*n2.b,
		b: n1.a*n2.b + n1.b*n2.a,
	}
}

func (n Number) Times(factor float64) Number {
	return Number{
		a: n.a * factor,
		b: n.b * factor,
	}
}

func (n1 Number) Divide(n2 Number) Number {
	return Number{
		a: (n1.a*n2.a + n1.b*n2.b) / (n2.a*n2.a + n2.b*n2.b),
        b: (n1.b*n2.a - n1.a*n2.b) / (n2.a*n2.a + n2.b*n2.b),
	}
}

func (n Number) Conjugate() Number {
	return Number{
		a: n.a,
		b: -n.b,
	}
}

func (n Number) Abs() float64 {
	return math.Sqrt(n.a*n.a + n.b*n.b)
}

func (n Number) Exp() Number {
	return Number{
		a: math.Exp(n.a) * math.Cos(n.b),
		b: math.Exp(n.a) * math.Sin(n.b),
	}
}
