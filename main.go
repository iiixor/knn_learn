package main

import (
	"fmt"
)

func main() {

	defer lg.Sync()

	// Пример логирования информационного события с контекстом
	var dataSettings = DataSettings{}

	dataSettings.Data, err = ReadCSV("./samples/task13.csv")
	if err != nil {
		lg.Errorf("Failed to import CSV file: %v", err)
	}
	lg.Info("CSV filed imported successfully")
	for _, sample := range dataSettings.Data {
		fmt.Println(sample)
	}

	//Добавление весов и p
	dataSettings.Weights = []float64{0.33477068, 0.31833993, 0.34688940}
	dataSettings.P = 2.5
	k := ChoiseKs(dataSettings.Data)

	dataSettings.Data = Normalize("Z", dataSettings.Data)

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

	KRes := EvaluateK(k, matrix)
	bestK := FindBestK(KRes)
	fmt.Printf("Best K : %d --- Acc : %.2f %%\n", bestK.K, bestK.Acc*100)
}
