package floatcalc

import "math"

type Calculator struct{}

func New() *Calculator {
	return &Calculator{}
}

func (*Calculator) Add(x, y float64) float64 {
	return x + y
}

func (*Calculator) Sub(x, y float64) float64 {
	return x - y
}

func (*Calculator) Mul(x, y float64) float64 {
	return x * y
}

func (*Calculator) Div(x, y float64) float64 {
	return x / y
}

func (*Calculator) Mod(float64, float64) float64 {
	panic("not implemented")
}

func (*Calculator) Pow(float64, float64) float64 {
	panic("not implemented")
}

func (*Calculator) Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func (*Calculator) Log(x float64) float64 {
	return math.Log(x)
}
