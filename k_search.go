package main

import (
	"fmt"
	"sort"
)

func MakeMatrix(data DataSettings) ([][]MatrixFill, error) {
	num := 0
	for _, d := range data.Data {
		if d.Status == -1.00 {
			num += 1
		}
	}
	size := len(data.Data) - num
	matrix := make([][]MatrixFill, size)
	dataSize := len(data.Data[0].Values)

	for i := range matrix {
		matrix[i] = make([]MatrixFill, size)

		for k := range matrix[i] {
			if len(data.Data[i].Values) < dataSize {
				return nil, fmt.Errorf("insufficient values at index %d", i)
			}

			dx, err := WeightedMinkowski(
				data.Data[i].Values[1:dataSize],
				data.Data[k].Values[1:dataSize],
				data.Weights,
				data.P,
			)
			if err != nil {
				return nil, err
			}

			matrix[i][k] = MatrixFill{
				Number: float64(k) + 1,
				Status: data.Data[k].Status,
				Value:  dx,
			}
			// lg.Info(dx)
		}
	}
	fmt.Printf("---\nMatrix Fill:\n---\n")
	for _, mat := range matrix {
		for _, matt := range mat {
			fmt.Printf("|  %.2f  ", matt.Value)
		}
		fmt.Println()
	}
	fmt.Printf("---\n")
	return matrix, nil
}

func SortMatrixFill(matrix []MatrixFill) []MatrixFill {
	sort.Slice(matrix, func(i, j int) bool {
		return matrix[i].Value < matrix[j].Value
	})
	return matrix
}

func mode(slice []float64) float64 {
	// Проверка на пустой слайс
	if len(slice) == 0 {
		return 0
	}

	// Создаем карту для подсчета частоты каждого значения
	frequency := make(map[float64]int)
	for _, value := range slice {
		frequency[value]++
	}

	// Находим максимальную частоту
	maxFrequency := 0
	for _, count := range frequency {
		if count > maxFrequency {
			maxFrequency = count
		}
	}

	// Собираем все значения с максимальной частотой
	var modes []float64
	for value, count := range frequency {
		if count == maxFrequency {
			modes = append(modes, value)
		}
	}

	// Если только одна мода, возвращаем её
	if len(modes) == 1 {
		return modes[0]
	}

	// Если несколько мод, вычисляем их среднее
	var sum float64
	for _, value := range modes {
		sum += value
	}

	return -1
}

// EvaluateK оценивает точность предсказаний для различных значений k
func EvaluateK(k []int, matrix [][]MatrixFill) []KResult {
	// Инициализируем слайс результатов с нужным размером
	kResults := make([]KResult, len(k))

	// Заполняем значения K в результатах заранее
	for i, ki := range k {
		kResults[i].K = ki
	}

	// Обрабатываем каждую строку матрицы
	for _, line := range matrix {
		// Сортируем строку только один раз для всех k
		sortedLine := SortMatrixFill(line)

		// for _, value := range sortedLine {
		// 	fmt.Println(value)
		// }
		// Получаем фактическое значение вне цикла по k
		factValue := sortedLine[0].Status

		// Анализируем каждое значение k
		for i, ki := range k {
			// Проверяем, что у нас достаточно элементов для этого k
			if ki > len(sortedLine) {
				continue
			}

			// Создаем слайс правильного размера (ki-1, а не ki)
			values := make([]float64, ki)

			// Заполняем слайс значениями

			for j := 1; j <= ki; j++ {
				values[j-1] = sortedLine[j].Status
			}
			// Определяем моду
			var modeValue float64
			if ki == 1 {
				// Особый случай для k=1
				values = []float64{sortedLine[1].Status}
			}
			modeValue = mode(values)
			// Обновляем статистику
			kResults[i].All++
			fmt.Printf("Attempt: %d --- K: %d\n", kResults[i].All, kResults[i].K)
			fmt.Printf("Fact: %.2f --- Mode: %.2f = %v\n", factValue, modeValue, values)
			if factValue == modeValue {
				kResults[i].Good++
			}
			// fmt.Println(kResults[i])
		}
	}

	// Вычисляем точность для каждого k с помощью индексов,
	// чтобы изменения затрагивали оригинальные объекты
	for i := range kResults {
		if kResults[i].All > 0 {
			kResults[i].Acc = float64(kResults[i].Good) / float64(kResults[i].All)
			fmt.Printf("K : %d --- Acc : %.2f %%\n", kResults[i].K, kResults[i].Acc*100)
		}
	}

	return kResults
}

func FindBestK(ks []KResult) KResult {
	var bestK KResult
	bestK.Acc = ks[0].Acc
	for _, k := range ks[1:] {
		if k.Acc >= bestK.Acc {
			bestK = k
		}
	}
	return bestK
}

func ChoiseKs(DataPoints []DataPoint) []int {
	var k []int
	for i := 1; i <= len(DataPoints); i += 2 {
		k = append(k, i)
	}
	return k
}
