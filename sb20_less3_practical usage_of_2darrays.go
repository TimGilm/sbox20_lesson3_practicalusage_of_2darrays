package main

import "fmt"

const rows = 3
const cols = 5

func maximum(input [cols]int) int {
	maxx := input[0]
	for i := 0; i < len(input); i++ {
		if input[i] > maxx {
			maxx = input[i]
		}
	}
	return maxx
}

var result [rows]int

func findMaxInMatrix(matrix [rows][cols]int) [rows]int {
	for i := 0; i < len(matrix); i++ {
		lmax := maximum(matrix[i]) //встроенная функция maximum!
		result[i] = lmax
	}
	return result //[5 101 50] вывод на печать
}

func sumAll(matrix [rows][cols]int) int {
	sum := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			sum = sum + matrix[i][j]
		}
	}
	return sum
}

func sumNotEvenCol(matrix [rows][cols]int) int {
	sum := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if j%2 != 0 {
				sum = sum + matrix[i][j]
			}
		}
	}
	return sum
}

func main() {
	matrix := [rows][cols]int{
		{1, 2, 3, 4, 5},
		{3, 2, 10, 100, 101},
		{40, 50, -10, -20, -30},
	}
	fmt.Println(findMaxInMatrix(matrix))
	fmt.Println(sumAll(matrix))
	fmt.Println(sumNotEvenCol(matrix))
}
