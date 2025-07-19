package libSvm

import "fmt"

func Fit(X [][]float64, y []int, param *Parameter) (*Problem, error) {
	problem := &Problem{}

	problem.y = nil
	problem.x = nil
	problem.xSpace = nil

	fmt.Println(len(problem.xSpace), len(problem.x), len(problem.y))

	return problem, nil
}
