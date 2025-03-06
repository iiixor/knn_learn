package main

import (
	"fmt"
)

func main() {

	defer lg.Sync()

	// Пример логирования информационного события с контекстом
	var dataSettings = DataSettings{}

	dataSettings.Data, err = ReadCSV("./samples/task8.csv")
	if err != nil {
		lg.Errorf("Failed to import CSV file: %v", err)
	}
	lg.Info("CSV filed imported successfully")
	for _, sample := range dataSettings.Data {
		fmt.Println(sample)
	}

	//Добавление весов и p
	dataSettings.Weights = []float64{0.1, 0.4, 0.3, 0.2}
	dataSettings.P = 1.5

	matrix, err := MakeMatrix(dataSettings)
	if err != nil {
		lg.Errorf("Failed to make K matrix: %v", err)
	}
	for _, line := range matrix {
		sorted := SortMatrixFill(line)
		for _, line := range sorted {
			fmt.Printf("%s %.2f\n", line.Status, line.Value)
		}
		fmt.Println("---")
	}
}
