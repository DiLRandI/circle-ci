package calc

type Calculator struct{}

func New() *Calculator {
	return &Calculator{}
}

func (*Calculator) Add(x, y int) int {
	return x + y
}
