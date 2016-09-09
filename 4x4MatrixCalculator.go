/*TODO:
1. matrix generating function - 1
2. create matrix class - 1
3. add - 1
4. subtract - 1
5. Multiply - 1
6. Transpone - 1
7. String method for easy-to-read output - 1
Optional :
Gopher output :O - 0
Some unit tests - 0
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
}

type MatrixMethods interface {
	add(inputMatrix [][]int64) [][]int64
	subtract(inputMatrix [][]int64) [][]int64
	Multiplication(inputMatrix [][]int64) [][]int64
	Transpone() [][]int64
	String()
}

type Matrix [][]int64

//constructor
func NewThing(inMatrix [][]int64) *Matrix {
	p := new(Matrix)
	return p
}

func (m Matrix) add(inputMatrix [][]int64) [][]int64 {
	sum := GenerateEmptyMatrix()
	//fmt.Println("Base Matrix = ", m)
	//fmt.Println("Input Matrix = ", inputMatrix)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sum[i][j] = m[i][j] + inputMatrix[i][j]
			fmt.Println(m[i][j], " + ", inputMatrix[i][j], " = ", sum[i][j])
		}
	}
	return sum
}

func (m Matrix) subtract(inputMatrix [][]int64) [][]int64 {
	sub := GenerateEmptyMatrix()
	fmt.Println("Base Matrix = ", m)
	fmt.Println("Input Matrix = ", inputMatrix)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sub[i][j] = m[i][j] - inputMatrix[i][j]
			fmt.Println(m[i][j], " - ", inputMatrix[i][j], " = ", sub[i][j])
		}
	}
	return sub
}

func (m Matrix) Multiplication(inputMatrix [][]int64) [][]int64 {
	mul := GenerateEmptyMatrix()
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				mul[i][j] += m[i][k] * inputMatrix[k][j]
			}
		}
	}
	return mul
}

func (m Matrix) String() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 1; j++ {
			fmt.Println("-------------------------")
			fmt.Println("|", m[i][j], " | ", m[i][j+1], " | ", m[i][j+2], " | ", m[i][j+3])
		}
	}

}

func (m Matrix) Transpone() [][]int64 {
	transponedMatrix := GenerateEmptyMatrix()
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			transponedMatrix[j][i] = m[i][j]
		}
	}
	return transponedMatrix
}

// Helper functions , to avoid redundant code
func GenerateEmptyMatrix() [][]int64 {
	emptyMatrix := make([][]int64, 4)
	for i := 0; i < 4; i++ {
		emptyMatrix[i] = make([]int64, 4)
		for j := 0; j < 4; j++ {
			emptyMatrix[i][j] = 0
		}
	}
	return emptyMatrix
}

//automatically generates nonsingular Matrix
func nonsingularRandomMatrix() [][]int64 {
	rand.Seed(time.Now().UTC().UnixNano())
	var matrixDeterminant int64 = 0
	nm := make([][]int64, 4)
	for i := 0; i < 4; i++ {
		nm[i] = make([]int64, 4)
		for j := 0; j < 4; j++ {
			nm[i][j] = rand.Int63n(100)

		}
	}
	matrixDeterminant =
		nm[0][0]*nm[1][1]*nm[2][2]*nm[3][3] - nm[0][0]*nm[1][1]*nm[2][3]*nm[3][2] -
			nm[0][0]*nm[1][2]*nm[2][1]*nm[3][3] + nm[0][0]*nm[1][2]*nm[2][3]*nm[3][1] +
			nm[0][0]*nm[1][3]*nm[2][1]*nm[3][2] - nm[0][0]*nm[1][3]*nm[2][2]*nm[3][1] -
			nm[0][1]*nm[1][0]*nm[2][2]*nm[3][3] + nm[0][1]*nm[1][0]*nm[2][3]*nm[3][2] +
			nm[0][1]*nm[1][2]*nm[2][0]*nm[3][3] - nm[0][1]*nm[1][2]*nm[2][3]*nm[3][0] -
			nm[0][1]*nm[1][3]*nm[2][0]*nm[3][2] + nm[0][1]*nm[1][3]*nm[2][2]*nm[3][0] +
			nm[0][2]*nm[1][0]*nm[2][1]*nm[3][3] - nm[0][2]*nm[1][0]*nm[2][3]*nm[3][1] -
			nm[0][2]*nm[1][1]*nm[2][0]*nm[3][3] + nm[0][2]*nm[1][1]*nm[2][3]*nm[3][0] +
			nm[0][2]*nm[1][3]*nm[2][0]*nm[3][1] - nm[0][2]*nm[1][3]*nm[2][1]*nm[3][0] -
			nm[0][3]*nm[1][0]*nm[2][1]*nm[3][2] + nm[0][3]*nm[1][0]*nm[2][2]*nm[3][1] +
			nm[0][3]*nm[1][1]*nm[2][0]*nm[3][2] - nm[0][3]*nm[1][1]*nm[2][2]*nm[3][0] -
			nm[0][3]*nm[1][2]*nm[2][0]*nm[3][1] + nm[0][3]*nm[1][2]*nm[2][1]*nm[3][0]
	if matrixDeterminant == 0 {
		nonsingularRandomMatrix()
	}
	return nm
}
