package main

import (
	"fmt"
	"sort"
)

type MatrixFill struct {
	Status string
	Value  float64
}

func MakeMatrix(data DataSettings) ([][]MatrixFill, error) {
	size := len(data.Data) - 1
	matrix := make([][]MatrixFill, size)

	for i := range matrix {
		matrix[i] = make([]MatrixFill, size)

		for k := range matrix[i] {
			if len(data.Data[i].Values) < 5 {
				return nil, fmt.Errorf("insufficient values at index %d", i)
			}

			dx, err := WeightedMinkowski(
				data.Data[i].Values[1:5],
				data.Data[k].Values[1:5],
				data.Weights,
				data.P,
			)
			if err != nil {
				return nil, err
			}

			matrix[i][k] = MatrixFill{
				Status: data.Data[k].Status,
				Value:  dx,
			}
			lg.Info(dx)
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
