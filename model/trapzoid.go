package model

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

var TrapzoidNumber *uint //Число трапеций

// Подпрограмма нахождения интеграла методом трапеций
func IntrgrateTrapzoid(formula string, a float64, b float64) (float64, error) {
	var s float64 = 0
	var fx0 float64 = 0
	var fxh float64 = 0
	var err error
	h := math.Abs(b-a) / float64(*TrapzoidNumber)
	x0 := a
	vars := map[string]interface{}{
		"x":  a,
		"pi": 3.14159265359,
		"E":  2.71828182845,
	}
	fx0, err = EvaluateExpression(formula, vars)
	fmt.Printf("fx0 =%f\n", fx0)
	if err != nil {
		return 0, err
	}
	s += fx0
	vars = map[string]interface{}{
		"x": b,
	}
	fx0, err = EvaluateExpression(formula, vars)
	fmt.Printf("fxn =%f\n", fx0)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	s += fx0
	s /= 2
	var i uint
	for i = 1; i < *TrapzoidNumber; i++ {
		xh := x0 + h
		fmt.Printf("x%d =%f\n", i, x0)
		vars = map[string]interface{}{
			"x": xh,
		}
		fxh, err = EvaluateExpression(formula, vars)
		fmt.Printf("fx%d =%f\n", i, fxh)
		if err != nil {
			s = 0
			log.Println(err)
			break
		}
		s = s + fxh
		x0 = xh
		fx0 = fxh
	}
	s = s * h
	return s, err
}

// Подпрограмма приведения ответа в строчный формат
func IntrgrateTrapzoidString(formula string, a float64, b float64) string {
	var s string = "null"
	f, err := IntrgrateTrapzoid(formula, a, b)
	if err != nil {
		s = err.Error()
	} else {
		s = strconv.FormatFloat(f, 'f', 6, 32)
	}
	return s
}
