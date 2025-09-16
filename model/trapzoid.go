package model

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

var PointNumbers *uint

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

func TrapzoidMethod(formula string, a float64, b float64) (float64, error) {
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
	fmt.Printf("xn = %f \n", a)
	fmt.Printf("xn+1 = %f \n", b)
	fxa, err := EvaluateExpression(formula, vars1)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("func(xn) = %f \n", fxa)
		fxb, err := EvaluateExpression(formula, vars2)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Printf("func(xn+1) = %f \n", fxb)
			f = math.Abs(b-a) * (math.Abs(fxb) + math.Abs(fxa)) / 2
		}
	}
	return f, err
}

func Intrgrate(formula string, a float64, b float64) (float64, error) {
	var f float64 = 0
	var fx float64 = 0
	var err error
	h := math.Abs(b-a) / float64(*PointNumbers-1)
	x0 := a
	var i uint
	for i = 0; i < *PointNumbers; i++ {
		xh := x0 + h
		fx, err = TrapzoidMethod(formula, x0, xh)
		f = f + fx
		x0 = xh
	}
	return f, err
}

func IntrgrateString(formula string, a float64, b float64) string {
	var f float64 = 0
	var fx float64 = 0
	var s string = "null"
	var err error
	h := math.Abs(b-a) / float64(*PointNumbers-1)
	x0 := a
	var i uint
	for i = 2; i <= *PointNumbers; i++ {
		xh := x0 + h
		fx, err = TrapzoidMethod(formula, x0, xh)
		if err != nil {
			break
		}
		f = f + fx
		x0 = xh
	}
	if err != nil {
		s = err.Error()
	} else {
		s = strconv.FormatFloat(f, 'f', 6, 32)
	}
	return s
}

// Более оптимизированный метод трапеций
func IntrgrateTrapzoid(formula string, a float64, b float64) (float64, error) {
	var s float64 = 0
	var fx0 float64 = 0
	var fxh float64 = 0
	var err error

	x0 := a
	vars := map[string]interface{}{
		"x":  x0,
		"pi": 3.14159265359,
		"E":  2.71828182845,
	}
	fx0, err = EvaluateExpression(formula, vars)

	h := math.Abs(b-a) / float64(*PointNumbers-1)

	if err != nil {
		log.Println(err)
		return 0, err
	}
	var i uint
	for i = 1; i < *PointNumbers; i++ {
		xh := x0 + h
		fmt.Printf("x%d =%f\n", i, x0)
		vars = map[string]interface{}{
			"x": xh,
		}
		fxh, err = EvaluateExpression(formula, vars)
		fmt.Printf("fx%d =%f\n", i, fxh)
		if err != nil {
			log.Println(err)
			break
		}
		s = s + (h * (math.Abs(fxh) + math.Abs(fx0)) / 2)
		x0 = xh
		fx0 = fxh
	}
	return s, err
}

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
