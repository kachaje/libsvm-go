package libSvm_test

import (
	"fmt"
	"path/filepath"
	"testing"

	libSvm "github.com/kachaje/libsvm-go/v3"
)

func TestProblemRead(t *testing.T) {
	param := libSvm.NewParameter()

	problem, err := libSvm.NewProblem(filepath.Join(".", "cmds", "fixtures", "sample"), param)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v\n", problem)
}

func TestFit(t *testing.T) {
	y := []int{1, -1, 1, -1, 1, -1}
	X := [][]float64{
		{0.5, 1.2, 0.8},
		{0.6, 1.1, 0.7},
		{0.4, 1.3, 0.9},
		{0.7, 1.0, 0.6},
		{0.5, 1.2, 0.8},
		{0.6, 1.1, 0.7},
	}

	problem, err := libSvm.Fit(X, y, nil)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", problem)
}
