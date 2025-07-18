package tests

import (
	"os"
	"path/filepath"
	"testing"

	libSvm "github.com/kachaje/libsvm-go"
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
