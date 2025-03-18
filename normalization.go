package main

import (
	"fmt"
	"math"
)

func Normalize(norm_type string, d []DataPoint) []DataPoint {
	var nd []DataPoint
	switch norm_type {
	case "MinMax":
		nd = NormalizeMinMax(d)
	case "Z":
		nd = NormalizeZ(d)
	default:
		norm_type = "MinMax"
		nd = NormalizeMinMax(d)
	}

	fmt.Printf("---\nNormalization Fill: %s\n---\n", norm_type)
	for _, sample := range nd {
		for _, val := range sample.Values {
			fmt.Printf("| %.2f ", val)
		}
		fmt.Printf("| %.f \n", sample.Status)
	}

	return nd

}

func NormalizeMinMax(d []DataPoint) []DataPoint {
	for vi := 1; vi < len(d[0].Values); vi++ {
		normolized := make([]float64, len(d))
		for ei := range d {
			normolized[ei] = d[ei].Values[vi]
		}
		normilized := NormalizeMinMaxSlice(normolized)
		for ei := range d {
			d[ei].Values[vi] = normilized[ei]
			// fmt.Println(d[ei].Values[vi])
		}
	}
	return d
}

func NormalizeMinMaxSlice(s []float64) []float64 {
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

func NormalizeZ(d []DataPoint) []DataPoint {
	for vi := 1; vi < len(d[0].Values); vi++ {
		normolized := make([]float64, len(d))
		for ei := range d {
			normolized[ei] = d[ei].Values[vi]
		}
		normilized := NormalizeZSlice(normolized)
		for ei := range d {
			d[ei].Values[vi] = normilized[ei]
		}
	}
	return d

}

func NormalizeZSlice(s []float64) []float64 {
	// Проверка на пустой слайс
	if len(s) == 0 {
		return nil
	}

	// Проверка на слайс из одного элемента
	if len(s) == 1 {
		return []float64{0} // для одного элемента Z-оценка всегда 0
	}

	// Вычисление среднего значения
	avg := 0.0
	for _, v := range s {
		avg += v
	}
	avg /= float64(len(s))

	// Вычисление выборочного стандартного отклонения
	// Используем n-1 в знаменателе (формула для выборки)
	stdDev := 0.0
	for _, v := range s {
		stdDev += math.Pow(v-avg, 2)
	}
	stdDev /= float64(len(s) - 1)
	stdDev = math.Sqrt(stdDev)

	// Проверка на нулевое стандартное отклонение
	if stdDev == 0 {
		// Все значения одинаковые, возвращаем слайс с нулями
		return make([]float64, len(s))
	}

	// Нормализация значений (Z-преобразование)
	normalized := make([]float64, len(s))
	for i, v := range s {
		normalized[i] = (v - avg) / stdDev
	}

	return normalized
}
