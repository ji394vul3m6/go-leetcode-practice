package solutions

import (
	"fmt"
	"sort"
)

func BullsandCows() bool {
	type testCase struct {
		Secret string
		Guess  string
		Expect string
	}
	testCases := []testCase{
		testCase{"1807", "7810", "1A3B"},
		testCase{"1123", "0111", "1A1B"},
		testCase{"11233", "33211", "1A4B"},
	}
	ret := true
	for idx, testcase := range testCases {
		result := getHint(testcase.Secret, testcase.Guess)
		if result != testcase.Expect {
			ret = false
			fmt.Printf("Case %d fail: %s, %s, '%s', get '%s'\n", idx, testcase.Secret, testcase.Guess, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func getHint(secret string, guess string) string {
	if len(secret) != len(guess) {
		// invalid format
		return ""
	}

	bulls := 0
	cows := 0
	secretRemain := []int{}
	guessRemain := []int{}

	for idx := range secret {
		if secret[idx] == guess[idx] {
			bulls++
		} else {
			secretRemain = append(secretRemain, int(secret[idx]))
			guessRemain = append(guessRemain, int(guess[idx]))
		}
	}

	sort.Ints(secretRemain)
	sort.Ints(guessRemain)

	idxSecret := 0
	idxGuess := 0
	for idxSecret < len(secretRemain) && idxGuess < len(guessRemain) {
		if secretRemain[idxSecret] == guessRemain[idxGuess] {
			cows++
			idxSecret++
			idxGuess++
		} else if secretRemain[idxSecret] < guessRemain[idxGuess] {
			idxSecret++
		} else {
			idxGuess++
		}
	}

	return fmt.Sprintf("%dA%dB", bulls, cows)
}

/*
nAnB game
You are playing the following Bulls and Cows game with your friend: You write down a number and ask your friend to guess what the number is. Each time your friend makes a guess, you provide a hint that indicates how many digits in said guess match your secret number exactly in both digit and position (called "bulls") and how many digits match the secret number but locate in the wrong position (called "cows"). Your friend will use successive guesses and hints to eventually derive the secret number.

For example:

Secret number:  "1807"
Friend's guess: "7810"
Hint: 1 bull and 3 cows. (The bull is 8, the cows are 0, 1 and 7.)
Write a function to return a hint according to the secret number and friend's guess, use A to indicate the bulls and B to indicate the cows. In the above example, your function should return "1A3B".

Please note that both secret number and friend's guess may contain duplicate digits, for example:

Secret number:  "1123"
Friend's guess: "0111"
In this case, the 1st 1 in friend's guess is a bull, the 2nd or 3rd 1 is a cow, and your function should return "1A1B".
You may assume that the secret number and your friend's guess only contain digits, and their lengths are always equal.
*/
