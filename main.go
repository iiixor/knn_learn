package main

import (
	"fmt"
)

func main() {

	defer lg.Sync()

	// Пример логирования информационного события с контекстом
	var dataSettings = DataSettings{}

	dataSettings.Data, err = ReadCSV("./samples/task3.csv")
	if err != nil {
		lg.Errorf("Failed to import CSV file: %v", err)
	}
	lg.Info("CSV filed imported successfully")
	for _, sample := range dataSettings.Data {
		fmt.Println(sample)
	}

	//Добавление весов и p
	dataSettings.Weights = []float64{1, 1, 1}
	dataSettings.P = 2
	k := []int{1, 3, 5}

	matrix, err := MakeMatrix(dataSettings)
	if err != nil {
		lg.Errorf("Failed to make K matrix: %v", err)
	}
	// for _, line := range matrix {
	// 	sorted := SortMatrixFill(line)
	// 	for _, line := range sorted {
	// 		fmt.Printf("%.f | %s | %.2f\n", line.Number, line.Status, line.Value)
	// 	}
	// 	fmt.Println("---")
	// }
	//

	EvaluateK(k, matrix)
}
