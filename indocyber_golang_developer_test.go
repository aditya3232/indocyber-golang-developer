// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func GetPalindrome(input string) bool {
	formatString := strings.ToLower(strings.ReplaceAll(input, " ", ""))

	for i := 0; i <= len(formatString)/2; i++ {
		if formatString[i] != formatString[len(formatString)-1-i] {
			return false
		}
	}

	return true
}

func TestGetPalindrome(t *testing.T) {
	result_1 := GetPalindrome("malam")
	expected_1 := true

	if result_1 != expected_1 {
		t.Errorf("result_1 returned %t, expected result %t", result_1, expected_1)
	}

	fmt.Println(result_1)
	fmt.Print("\n")
}

func GetMaxArrayValue(input []int) int {
	sort.Slice(input, func(i, j int) bool {
		return input[i] > input[j]
	})

	return input[0]
}

func TestGetMaxArrayValue(t *testing.T) {
	input := []int{3, 5, 1, 9, 2}
	result_2 := GetMaxArrayValue(input)
	expected_2 := 9

	if result_2 != expected_2 {
		t.Errorf("result_2 returned %d, expected result %d", result_2, expected_2)
	}

	fmt.Println(result_2)
	fmt.Print("\n")
}

func TestGetRightTriangle(t *testing.T) {
	input := 5

	for i := 0; i < input; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Print("\n")
}

func GetFactorialWithRecursive(input int) int {
	if input == 1 {
		return 1
	} else {
		return input * GetFactorialWithRecursive(input-1)
	}
}

func TestFactorialWithRecursive(t *testing.T) {
	result_4 := GetFactorialWithRecursive(5)
	expected_4 := 120

	if result_4 != expected_4 {
		t.Errorf("result_4 returned %d, expected result %d", result_4, expected_4)
	}

	fmt.Println(result_4)
	fmt.Print("\n")
}
