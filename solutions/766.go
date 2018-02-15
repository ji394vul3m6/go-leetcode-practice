package solutions

import (
	"bytes"
	"fmt"
)

func IsToeplitzMatrix() bool {
	type testCase struct {
		Matrix [][]int
		Expect bool
	}
	testCases := []testCase{
		testCase{[][]int{
			[]int{1, 2, 3, 4},
			[]int{5, 1, 2, 3},
			[]int{9, 5, 1, 2},
		}, true},
		testCase{[][]int{
			[]int{1, 2, 3, 4},
			[]int{5, 1, 2, 3},
			[]int{9, 4, 1, 2},
		}, false},
		testCase{[][]int{
			[]int{1, 2},
			[]int{2, 2},
		}, false},
		testCase{[][]int{
			[]int{1},
		}, true},
		testCase{[][]int{}, true},
	}
	ret := true
	for idx, testcase := range testCases {
		result := isToeplitzMatrix(testcase.Matrix)
		if result != testcase.Expect {
			ret = false
			fmt.Printf("Case %d fail:\nInput:\n%s\nExcept %t, Get %t\n", idx, arrayToString(testcase.Matrix), testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func isToeplitzMatrix(matrix [][]int) bool {
	if len(matrix) == 0 {
		return true
	}
	rows := len(matrix)
	columns := len(matrix[0])

	for i := 0; i < columns; i++ {
		standard := matrix[0][i]
		for x, y := i, 0; x < columns && y < rows; x, y = x+1, y+1 {
			if matrix[y][x] != standard {
				return false
			}
		}
	}
	for i := 1; i < rows; i++ {
		standard := matrix[i][0]
		for x, y := 0, i; x < columns && y < rows; x, y = x+1, y+1 {
			if matrix[y][x] != standard {
				return false
			}
		}
	}

	return true
}

func arrayToString(matrix [][]int) string {
	var buffer bytes.Buffer
	for _, row := range matrix {
		buffer.WriteString("[")
		for _, col := range row {
			buffer.WriteString(fmt.Sprintf("% 4d ", col))
		}
		buffer.WriteString("]\n")
	}
	return buffer.String()
}
