package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

func ReadCSV(path string) ([]DataPoint, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var dataset []DataPoint
	for i, record := range records {
		// Пропускаем заголовки
		if i == 0 { // Пропускаем заголовок
			continue
		}
		values := make([]float64, len(record)-1)
		status := float64(0)
		for j := range record {
			val, err := strconv.ParseFloat(record[j], 64)
			if err != nil {
				lg.Fatalf("Ошибка преобразования значения %q в строке %d: %v", record[j], i+1, err)
			}
			if j == len(record)-1 {
				status = val
				break
			}
			values[j] = val
		}

		dataset = append(dataset, DataPoint{
			Values: values,
			Status: status,
		})
	}
	return dataset, nil
}
