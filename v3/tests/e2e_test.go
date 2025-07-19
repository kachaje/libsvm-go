package tests

import (
	"os"
	"path/filepath"
	"testing"

	libSvm "github.com/kachaje/libsvm-go/v3"
)

func TestTrain(t *testing.T) {
	param := libSvm.NewParameter() // Create a parameter object with default values
	param.KernelType = libSvm.POLY // Use the polynomial kernel

	model := libSvm.NewModel(param) // Create a model object from the parameter attributes

	// Create a problem specification from the training data and parameter attributes
	problem, err := libSvm.NewProblem(filepath.Join("fixtures", "a9a.train"), param)
	if err != nil {
		t.Fatal(err)
	}

	model.Train(problem) // Train the model from the problem specification

	model.Dump("a9a.model") // Dump the model into a user-specified file
	defer func() {
		os.Remove("a9a.model")
	}()

	if _, err := os.Stat("a9a.model"); os.IsNotExist(err) {
		t.Fatal("Test failed. Failed to create model")
	}
}

func TestPredict(t *testing.T) {
	model := libSvm.NewModelFromFile(filepath.Join("fixtures", "a9a.model"))

	x := make(map[int]float64)
	// Populate with the test vector

	predictLabel := model.Predict(x) // Predicts a float64 label given the test vector

	if predictLabel != -1 {
		t.Fatalf("Test failed. Expected: -1; Actual: %f", predictLabel)
	}
}

func TestCustom(t *testing.T) {
	param := libSvm.NewParameter()
	param.C = 4

	y := []float64{1, -1, 1, -1, -1}
	X := [][]float64{
		{0.708333, 1, 1, -0.320755, -0.105023, -1, 1, -0.419847, -1, -0.225806, 0, 1, -1},
		{0.583333, -1, 0.333333, -0.603774, 1, -1, 1, 0.358779, -1, -0.483871, 0, -1, 1},
		{0.166667, 1, -0.333333, -0.433962, -0.383562, -1, -1, 0.0687023, -1, -0.903226, -1, -1, 1},
		{0.458333, 1, 1, -0.358491, -0.374429, -1, -1, -0.480916, 1, -0.935484, 0, -0.333333, 1},
		{0.875, -1, -0.333333, -0.509434, -0.347032, -1, 1, -0.236641, 1, -0.935484, -1, -0.333333, -1},
	}

	problem, err := libSvm.Fit(X, y, param)
	if err != nil {
		t.Fatal(err)
	}

	model := libSvm.NewModel(param)

	model.Train(problem)

	x := make(map[int]float64)

	predictLabel := model.Predict(x)

	if predictLabel != -1 {
		t.Fatalf("Test failed. Expected: -1; Actual: %f", predictLabel)
	}
}
