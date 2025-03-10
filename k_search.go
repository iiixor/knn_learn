package main

import (
	"fmt"
	"sort"
)

type MatrixFill struct {
	Number float64
	Status float64
	Value  float64
}

type KResult struct {
	K    int
	Good int
	All  int
	Acc  float64
}

func MakeMatrix(data DataSettings) ([][]MatrixFill, error) {
	size := len(data.Data) - 1
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
				Number: float64(k),
				Status: data.Data[k].Status,
				Value:  dx,
			}
			// lg.Info(dx)
		}
	}
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

	return sum / float64(len(modes))
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
			for j := 1; j < ki; j++ {
				values[j-1] = sortedLine[j].Status
			}

			// Определяем моду
			var modeValue float64
			if ki == 1 {
				// Особый случай для k=1
				if len(sortedLine) > 1 {
					modeValue = sortedLine[1].Status
				} else {
					continue
				}
			} else {
				modeValue = mode(values)
			}

			// Обновляем статистику
			kResults[i].All++
			if factValue == modeValue {
				kResults[i].Good++
			}
		}
	}

	// Вычисляем точность для каждого k с помощью индексов,
	// чтобы изменения затрагивали оригинальные объекты
	for i := range kResults {
		if kResults[i].All > 0 {
			kResults[i].Acc = float64(kResults[i].Good) / float64(kResults[i].All)
			fmt.Println(kResults[i])
		}
	}

	return kResults
}
