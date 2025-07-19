package libSvm

import "fmt"

func Fit(X [][]float64, y []float64, param *Parameter) (*Problem, error) {
	if len(X) != len(y) {
		return nil, fmt.Errorf("unbalanced input data")
	}

	problem := &Problem{}

	problem.y = y
	problem.x = []int{}
	problem.xSpace = []snode{}

	var max_idx int = 0
	var l int = 0

	for _, row := range X {
		problem.x = append(problem.x, len(problem.xSpace))

		for i, value := range row {
			if value == 0 {
				continue
			}

			index := i + 1

			node := snode{
				index: index,
				value: value,
			}

			problem.xSpace = append(problem.xSpace, node)
			if index > max_idx {
				max_idx = index
			}
		}

		problem.xSpace = append(problem.xSpace, snode{index: -1})
		l++
	}

	problem.l = l

	if param != nil && param.Gamma == 0 && max_idx > 0 {
		param.Gamma = 1.0 / float64(max_idx)
	}

	return problem, nil
}
