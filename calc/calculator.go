package calc

import "math"

type Calculator struct{}

func New() *Calculator {
	return &Calculator{}
}

func (*Calculator) Add(x, y int) int {
	return x + y
}

func (*Calculator) Sub(x, y int) int {
	return x - y
}

func (*Calculator) Mul(x, y int) int {
	return x * y
}

func (*Calculator) Div(x, y int) int {
	return x / y
}

func (*Calculator) Mod(x, y int) int {
	return x % y
}

func (*Calculator) Pow(x, y int) int {
	return x ^ y
}

func (*Calculator) Sqrt(x int) int {
	return int(math.Sqrt(float64(x)))
}

func (*Calculator) Log(x int) int {
	return int(math.Log(float64(x)))
}
