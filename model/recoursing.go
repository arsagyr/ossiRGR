package model

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func TrapzoidMethodStringReturn(formula string, a float64, b float64) string {
	var s string
	vars1 := map[string]interface{}{
		"x":  a,
		"pi": 3.14159265359,
		"E":  2.71828182845,
	}
	vars2 := map[string]interface{}{
		"x":  b,
		"pi": 3.14159265359,
		"E":  2.71828182845,
	}

	fxa, err := EvaluateExpression(formula, vars1)
	if err != nil {
		s = err.Error()
	} else {
		fmt.Printf("func(a) = %f \n", fxa)
		fxb, err := EvaluateExpression(formula, vars2)
		if err != nil {
			s = err.Error()
		} else {
			fmt.Printf("func(b) = %f \n", fxb)
			s = strconv.FormatFloat(math.Abs(b-a)*(math.Abs(fxb)+math.Abs(fxa))/2, 'f', 6, 32)
		}
	}
	return s
}
func TrapzoidMethod(formula string, a float64, b float64) float64 {
	var f float64 = 0
	vars1 := map[string]interface{}{
		"x":  a,
		"pi": 3.14159265359,
		"E":  2.71828182845,
	}
	vars2 := map[string]interface{}{
		"x":  b,
		"pi": 3.14159265359,
		"E":  2.71828182845,
	}

	fxa, err := EvaluateExpression(formula, vars1)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("func(a) = %f \n", fxa)
		fxb, err := EvaluateExpression(formula, vars2)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Printf("func(b) = %f \n", fxb)
			f = math.Abs(b-a) * (math.Abs(fxb) + math.Abs(fxa)) / 2
		}
	}
	return f
}
