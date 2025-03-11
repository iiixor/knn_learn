package main

import (
	"math"
)

func Normalize(d []DataPoint) []DataPoint {
	for vi := 1; vi < len(d[0].Values); vi++ {
		normolized := make([]float64, len(d))
		for ei := range d {
			normolized[ei] = d[ei].Values[vi]
		}
		// fmt.Println(normolized)
		normilized := NormalizeSlice(normolized)
		for ei := range d {
			d[ei].Values[vi] = normilized[ei]
			// fmt.Println(d[ei].Values[vi])
		}
	}
	return d
}

func NormalizeSlice(s []float64) []float64 {
	// Проверка на пустой слайс
	if len(s) == 0 {
		return nil
	}

	// Найти минимальное и максимальное значения
	minVal, maxVal := math.Inf(1), math.Inf(-1)
	for _, v := range s {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}

	// Если все значения одинаковы, возвращаем нули (или единицы)
	if minVal == maxVal {
		return make([]float64, len(s))
	}

	// Нормализация значений
	normalized := make([]float64, len(s))
	for i, v := range s {
		normalized[i] = (v - minVal) / (maxVal - minVal)
	}

	return normalized
}
