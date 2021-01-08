package main

var result []int

// RiverSizes is a function for question river size
func RiverSizes(matrix [][]int) []int {
	// Write your code here.
	result = []int{}
	h := len(matrix)
	if h == 0 {
		return nil
	}
	l := len(matrix[0])

	visited := make([][]bool, h)
	for i := range matrix {
		visited[i] = make([]bool, l)
	}

	for i := range matrix {
		for j := range matrix[i] {
			r := dfsUtil(matrix, i, j, visited)
			if r != 0 {
				result = append(result, r)
			}
		}
	}
	return result
}

func dfsUtil(grid [][]int, row int, col int, visited [][]bool) int {

	H := len(grid)
	L := len(grid[0])

	if row < 0 || col < 0 || row >= H || col >= L || visited[row][col] {
		return 0
	}

	//mark the cell visited
	visited[row][col] = true

	if grid[row][col] == 1 {
		return 1 + dfsUtil(grid, row+1, col, visited) + dfsUtil(grid, row-1, col, visited) + dfsUtil(grid, row, col+1, visited) + dfsUtil(grid, row, col-1, visited)
	}
	return 0

}
