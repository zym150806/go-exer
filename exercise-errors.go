package main

import "fmt"

type ErrNegativeSqrt struct {
	value float64
}

func (e *ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", e.value)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, &ErrNegativeSqrt{
			x,
			}
	}
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
