package numbers

import "errors"

// GetVars Devuelve tres tipos de enteros diferentes
func GetVars() (int, int32, int64) {
	var (
		a = 1
		b = 223412341
		c = 612342134123412344
	)
	return a, int32(b), int64(c)
}

// GetFloat devuelve dos flotantes
func GetFloat() (float32, float64) {
	return 0.1, float64(float32(0.1))
}

// Sum suma dos enteros de diferentes tipos
func Sum(a int, b int32) int32 {
	return int32(a) + b
}

//SumInterface dos int
func SumInterface(a interface{}, b interface{}) (int, error) {
	switch a.(type) {
	case string:
		return 0, errors.New("El valor A es un string")
	}
	switch b.(type) {
	case string:
		return 0, errors.New("El valor B es un string")
	}
	return a.(int) + b.(int), nil

}
