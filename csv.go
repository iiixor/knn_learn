package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

type DataPoint struct {
	Features []float64
	Label    string
	Dx       float64
}

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
		if i == 0 { // Пропускаем заголовок
			continue
		}

		features := make([]float64, len(record)-1)
		for j := 0; j < len(record)-1; j++ {
			val, err := strconv.ParseFloat(record[j], 64)
			if err != nil {
				log.Fatalf("Ошибка преобразования значения %q в строке %d: %v", record[j], i+1, err)
			}
			features[j] = val
		}

		dataset = append(dataset, DataPoint{
			Features: features,
			Label:    record[len(record)-1],
		})
	}
	return dataset, nil
}
