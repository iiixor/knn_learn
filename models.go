package main

type DataPoint struct {
	Values []float64
	Status float64
}

type DataSettings struct {
	Data    []DataPoint
	Weights []float64
	P       float64
}

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
