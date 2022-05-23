package main

import (
	"fmt"
	"math/rand"
)

const SIZE = 3

func randU8() uint8 {
	min := 0
	max := 255
	result := rand.Intn(max-min) + min
	return uint8(result)
}

func genRandMatrix() [SIZE][SIZE]uint8 {
	newMatrix := [SIZE][SIZE]uint8{}

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			newMatrix[i][j] = randU8()
		}
	}

	return newMatrix
}

func sumMatrixes(matrix1 [SIZE][SIZE]uint8, matrix2 [SIZE][SIZE]uint8) [SIZE][SIZE]int {
	newMatrix := [SIZE][SIZE]int{}

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			newMatrix[i][j] = int(matrix1[i][j]) + int(matrix2[i][j])
		}
	}

	return newMatrix
}

func truncateMatrix(matrix [SIZE][SIZE]int) [SIZE][SIZE]uint8 {
	newMatrix := [SIZE][SIZE]uint8{}

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if matrix[i][j] > 255 {
				newMatrix[i][j] = 255
			} else {
				newMatrix[i][j] = uint8(matrix[i][j])
			}
		}
	}

	return newMatrix
}

func normalizeMatrix(matrix [SIZE][SIZE]int) [SIZE][SIZE]uint8 {
	var min float32 = 0
	var max float32 = 0

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			floatCell := float32(matrix[i][j])

			if floatCell < min {
				min = floatCell
			}

			if floatCell > max {
				max = floatCell
			}
		}
	}

	var factor float32 = 255 / (max - min)

	newMatrix := [SIZE][SIZE]uint8{}

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			newMatrix[i][j] = uint8(factor * (float32(matrix[i][j]) - min))
		}
	}

	return newMatrix
}

func main() {
	fmt.Println("Running...")

	A := genRandMatrix()

	B := genRandMatrix()

	C := sumMatrixes(A, B)

	truncated := truncateMatrix(C)

	normalized := normalizeMatrix(C)

	fmt.Print("A: ")
	fmt.Println(A)

	fmt.Print("B: ")
	fmt.Println(B)

	fmt.Print("C: ")
	fmt.Println(C)

	fmt.Print("truncated: ")
	fmt.Println(truncated)

	fmt.Print("normalized: ")
	fmt.Println(normalized)
}
