package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*读内容*/
func printFileContents(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col) //注意换行符调整成\r，不然会读到0，windows下换行符为\r\n
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

type point struct {
	i, j int
}

//上左下右
var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func (p point) add(step point) point {
	return point{p.i + step.i, p.j + step.j}
}
func (p point) at(grid [][]int) (int, bool) {
	if p.i >= len(grid) ||
		p.j >= len(grid[0]) ||
		p.i < 0 ||
		p.j < 0 {
		return 0, false
	}

	return grid[p.i][p.j], true
}

func walk(maze [][]int, star, end point) [][]int {

	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[0]))
	}
	i := 0
	Q := []point{star}
	s := false
	for len(Q) > 0 {
		q := Q[0]
		Q = Q[1:]
		if q == end {
			s = true
			break
		}
		for _, d := range dirs {
			next := q.add(d)

			if next == star { //跳过起点
				continue
			}

			val, ok := next.at(maze)
			if !ok || val == 1 { //跳过墙和越界值
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 { //跳过走过点
				continue
			}
			currVal, _ := q.at(steps)           //获取当前点步数
			steps[next.i][next.j] = currVal + 1 //新点加步数
			if next == end {
				s = true
				break
			}
			i++
			Q = append(Q, next) //新点加入循环

		}
		if s {
			break
		}
	}
	fmt.Println(i)
	return steps
}

func main() {
	maze := readMaze("./src/maze/maze.in")
	for _, row := range maze {
		for _, col := range row {
			fmt.Printf("%3d ", col)
		}
		fmt.Println()
	}
	star := point{0, 0}
	end := point{5, 4}
	step := walk(maze, star, end)
	for _, row := range step {
		for _, col := range row {
			fmt.Printf("%3d ", col)
		}
		fmt.Println()
	}
	fmt.Println(step[end.i][end.j])
}
