package main

import (
	"fmt"
	"log/slog"
)

func main() {

	defer lg.Sync()

	// Пример логирования информационного события с контекстом
	var dataSettings = DataSettings{}

	dataSettings.Data, err = ReadCSV("./samples/knn_go.csv")
	if err != nil {
		lg.Error("Failed to import CSV file", slog.Any("Error", err))
	}
	lg.Info("CSV filed imported successfully")
	for _, sample := range dataSettings.Data {
		fmt.Println(sample)
	}

	//Добавление весов и p
	dataSettings.Weights = []float64{1.00, 1.00, 1.00}
	dataSettings.P = 2

	dx, _ := WeightedMinkowski(dataSettings.Data[1].Values[1:4], dataSettings.Data[2].Values[1:4], dataSettings.Weights, dataSettings.P)

	lg.Infof("DX {1;2} {p=2} = %f", dx)
	lg.Info(dataSettings.Data[1].Values[1:4])
}
