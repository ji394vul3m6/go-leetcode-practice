package solutions

import (
	"bytes"
	"fmt"
	"strings"
)

func IsValidSudoku() bool {
	type testCase struct {
		Nums   [][]byte
		Expect bool
	}
	testCases := []testCase{
		testCase{[][]byte{
			[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
			[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		}, true},
		testCase{[][]byte{
			[]byte{'5', '3', '3', '.', '7', '.', '.', '.', '.'},
			[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		}, false},
		testCase{[][]byte{
			[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
			[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			[]byte{'5', '9', '8', '.', '.', '.', '.', '6', '.'},
			[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		}, false},
		testCase{[][]byte{
			[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
			[]byte{'6', '8', '.', '1', '9', '5', '.', '.', '.'},
			[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		}, false},
	}
	ret := true
	for idx, testcase := range testCases {
		result := isValidSudoku(testcase.Nums)
		if result != testcase.Expect {
			ret = false
			fmt.Printf("Case %d fail: \n%s\n, except %t, get %t\n", idx, printBoard(testcase.Nums), testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func isValidSudoku(board [][]byte) bool {
	if len(board) != 9 {
		return false
	}
	for _, val := range board {
		if len(val) != 9 {
			return false
		}
	}

	showMap := map[byte]bool{
		'1': false,
		'2': false,
		'3': false,
		'4': false,
		'5': false,
		'6': false,
		'7': false,
		'8': false,
		'9': false,
	}
	// check row and column
	for i := 0; i < 9; i++ {
		resetMap(showMap)
		for j := 0; j < 9; j++ {
			val := board[i][j]
			if val == '.' {
				continue
			}
			if showMap[val] {
				return false
			}
			showMap[val] = true
		}

		resetMap(showMap)
		for j := 0; j < 9; j++ {
			val := board[j][i]
			if val == '.' {
				continue
			}
			if showMap[val] {
				return false
			}
			showMap[val] = true
		}
	}

	// check block
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			resetMap(showMap)
			for m := -1; m <= 1; m++ {
				for n := -1; n <= 1; n++ {
					val := board[3*i+1+m][3*j+1+n]
					if val == '.' {
						continue
					}
					if showMap[val] {
						return false
					}
					showMap[val] = true
				}
			}
		}
	}

	return true
}

func printBoard(board [][]byte) string {
	buf := bytes.Buffer{}
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			buf.WriteString(strings.Repeat("-", 25))
			buf.WriteString("\n")
		}
		for j := 0; j < 9; j++ {
			if j%3 == 0 {
				buf.WriteString("| ")
			}
			buf.WriteString(fmt.Sprintf("%c ", board[i][j]))
		}
		buf.WriteString(fmt.Sprintln("|"))
	}
	buf.WriteString(strings.Repeat("-", 24))
	buf.WriteString("\n")
	return buf.String()
}

func resetMap(showMap map[byte]bool) {
	for key := range showMap {
		showMap[key] = false
	}
}
