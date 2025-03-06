package main

import (
	"errors"
	"math"
)

// WeightedMinkowski вычисляет взвешенное расстояние Минковского
func WeightedMinkowski(x, y, weights []float64, p float64) (float64, error) {
    // Валидация входных данных
    if len(x) != len(y) || len(x) != len(weights) {
        return 0, errors.New("векторы и веса должны иметь одинаковую длину")
    }
    if p <= 0 {
        return 0, errors.New("параметр p должен быть положительным числом")
    }
    if len(x) == 0 {
        return 0, errors.New("векторы не могут быть пустыми")
    }

    var sum float64
    for i := range x {
        diff := math.Abs(x[i] - y[i])
        sum += weights[i] * math.Pow(diff, p)
    }

    return math.Pow(sum, 1/p), nil
}
