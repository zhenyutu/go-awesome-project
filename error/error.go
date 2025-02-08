package main

import (
	"errors"
	"fmt"
	"math"
	"reflect"
)

func main() {

	err := errors.New("this is an error")
	fmt.Println(err)

	res, e := Sqrt(-1)
	fmt.Println(res, e)
	fmt.Println(reflect.TypeOf(e))

}

func Sqrt(n int) (float64, error) {
	if n < 0 {
		return 0, errors.New("n is negative")
		//return 0, fmt.Errorf("n must be a positive number")
	}
	return math.Sqrt(float64(n)), nil
}
