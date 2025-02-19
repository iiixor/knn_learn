package main

import (
	"fmt"
	"log/slog"
)

func main() {
	slog.SetDefault(logger)

	// Пример логирования информационного события с контекстом
	data, err := ReadCSV("./knn_go.csv")
	if err != nil {
		slog.Error("Failed to import CSV file", slog.Any("Error", err))
	}
	slog.Info("CSV filed imported successfully")
	for _, sample := range data {
		fmt.Println(sample)
	}
}
