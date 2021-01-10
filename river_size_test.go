package main

import (
	"fmt"
	"testing"
)

func TestRiverSize(t *testing.T) {
	matrix := [][]int{
		[]int{1, 0, 0, 1, 0},
		[]int{1, 0, 1, 0, 0},
		[]int{0, 0, 1, 0, 1},
		[]int{1, 0, 1, 0, 1},
		[]int{1, 0, 1, 1, 0},
	}
	fmt.Println(RiverSizes(matrix))
	matrix = [][]int{
		[]int{0},
	}
	fmt.Println(RiverSizes(matrix))
	matrix = [][]int{
		[]int{1},
	}
	fmt.Println(RiverSizes(matrix))
	matrix = [][]int{
		[]int{1, 0, 0, 1},
		[]int{1, 0, 1, 0},
		[]int{0, 0, 1, 0},
		[]int{1, 0, 1, 0},
	}
	fmt.Println(RiverSizes(matrix))
	matrix = [][]int{
		[]int{1, 0, 0, 1, 0, 1, 0, 0, 1, 1, 1, 0},
		[]int{1, 0, 1, 0, 0, 1, 1, 1, 1, 0, 1, 0},
		[]int{0, 0, 1, 0, 1, 1, 0, 1, 0, 1, 1, 1},
		[]int{1, 0, 1, 0, 1, 1, 0, 0, 0, 1, 0, 0},
		[]int{1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 0, 1},
	}
	fmt.Println(RiverSizes(matrix))
}
