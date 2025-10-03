package model

import (
	"fmt"
	"math"

	expr "github.com/expr-lang/expr"
)

func EvaluateExpression(expression string, vars map[string]interface{}) (float64, error) {
	coef := expr.Env(vars) // Принимаем значения
	//Определяем интегрируемые функции
	sin := expr.Function(
		"sin",
		func(params ...any) (any, error) {
			switch params[0].(type) {
			case float64:
				return math.Sin(params[0].(float64)), nil
			}
			return nil, fmt.Errorf("invalid type")
		},
		new(func(float64) int),
	)
	cos := expr.Function(
		"cos",
		func(params ...any) (any, error) {
			switch params[0].(type) {
			case float64:
				return math.Cos(params[0].(float64)), nil
			}
			return nil, fmt.Errorf("invalid type")
		},
		new(func(float64) int),
	)
	tan := expr.Function(
		"tan",
		func(params ...any) (any, error) {
			switch params[0].(type) {
			case float64:
				return math.Tan(params[0].(float64)), nil
			}
			return nil, fmt.Errorf("invalid type")
		},
		new(func(float64) int),
	)
	asin := expr.Function(
		"asin",
		func(params ...any) (any, error) {
			switch params[0].(type) {
			case float64:
				return math.Asin(params[0].(float64)), nil
			}
			return nil, fmt.Errorf("invalid type")
		},
		new(func(float64) int),
	)
	acos := expr.Function(
		"acos",
		func(params ...any) (any, error) {
			switch params[0].(type) {
			case float64:
				return math.Acos(params[0].(float64)), nil
			}
			return nil, fmt.Errorf("invalid type")
		},
		new(func(float64) int),
	)
	atan := expr.Function(
		"atan",
		func(params ...any) (any, error) {
			switch params[0].(type) {
			case float64:
				return math.Atan(params[0].(float64)), nil
			}
			return nil, fmt.Errorf("invalid type")
		},
		new(func(float64) int),
	)
	ln := expr.Function(
		"ln",
		func(params ...any) (any, error) {
			switch params[0].(type) {
			case float64:
				return math.Log(params[0].(float64)), nil
			}
			return nil, fmt.Errorf("invalid type")
		},
		new(func(float64) int),
	)
	log2 := expr.Function(
		"log2",
		func(params ...any) (any, error) {
			switch params[0].(type) {
			case float64:
				return math.Log2(params[0].(float64)), nil
			}
			return nil, fmt.Errorf("invalid type")
		},
		new(func(float64) int),
	)
	log10 := expr.Function(
		"log10",
		func(params ...any) (any, error) {
			switch params[0].(type) {
			case float64:
				return math.Log10(params[0].(float64)), nil
			}
			return nil, fmt.Errorf("invalid type")
		},
		new(func(float64) int),
	)

	sqrt := expr.Function(
		"sqrt",
		func(params ...any) (any, error) {
			switch params[0].(type) {
			case float64:
				return math.Sqrt(params[0].(float64)), nil
			}
			return nil, fmt.Errorf("invalid type")
		},
		new(func(float64) int),
	)

	program, err := expr.Compile(expression, coef, sin, cos, tan, asin, acos, atan, ln, log2, log10, sqrt)
	if err != nil {
		return 0, fmt.Errorf("ошибка компиляции: %v", err)
	}
	// Выполняем выражение
	output, err := expr.Run(program, vars)
	if err != nil {
		return 0, fmt.Errorf("ошибка выполнения: %v", err)
	}
	// Приводим результат к float64
	switch result := output.(type) {
	case float64:
		return result, nil
	case int:
		return float64(result), nil
	default:
		return 0, fmt.Errorf("неожиданный тип результата: %T", output)
	}
}
