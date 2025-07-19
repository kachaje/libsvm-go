package libSvm_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	libSvm "github.com/kachaje/libsvm-go/v3"
)

func TestProblemJSON(t *testing.T) {
	param := libSvm.NewParameter()

	problem, err := libSvm.NewProblem(filepath.Join(".", "cmds", "fixtures", "sample"), param)
	if err != nil {
		t.Fatal(err)
	}

	result, err := problem.ToJSON()
	if err != nil {
		t.Fatal(err)
	}

	content, err := os.ReadFile(filepath.Join(".", "cmds", "fixtures", "sample.json"))
	if err != nil {
		t.Fatal(err)
	}

	refData := string(content)

	if *result != refData {
		t.Fatalf("Test failed. Expected: '%s'; Actual: '%s'\n", refData, *result)
	}
}

func TestFit(t *testing.T) {
	y := []float64{1, -1, 1, -1, -1}
	X := [][]float64{
		{0.708333, 1, 1, -0.320755, -0.105023, -1, 1, -0.419847, -1, -0.225806, 1, -1},
		{0.583333, -1, 0.333333, -0.603774, 1, -1, 1, 0.358779, -1, -0.483871, -1, 1},
		{0.166667, 1, -0.333333, -0.433962, -0.383562, -1, -1, 0.0687023, -1, -0.903226, -1, -1, 1},
		{0.458333, 1, 1, -0.358491, -0.374429, -1, -1, -0.480916, 1, -0.935484, -0.333333, 1},
		{0.875, -1, -0.333333, -0.509434, -0.347032, -1, 1, -0.236641, 1, -0.935484, -1, -0.333333, -1},
	}

	problem, err := libSvm.Fit(X, y, nil)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", problem)
}
