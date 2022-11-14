package main

import (
	"fmt"
	"os"
)

func main() {
	//行 列
	mapArr := readFile()
	printMap(mapArr)
	getStepCount(mapArr, 0, 0, 5, 4)
	printMap(mapArr)
}

func readFile() [][]int {
	open, err := os.Open("./src/crawler/maze/maze.in")
	if err != nil {
		panic(err)
	}
	var row, col int
	fmt.Fscanf(open, "%d %d", &row, &col) //注意换行符调整成\r，不然会读到0，windows下换行符为\r\n
	fmt.Println(row, col)
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(open, "%d", &maze[i][j])
		}
	}
	return maze
}

func printMap(mapArr [][]int) {
	for _, row := range mapArr {
		for _, col := range row {
			fmt.Printf("%3d", col)
		}
		fmt.Println()
	}
}
func getStepCount(mapArr [][]int, x int, y int, eX int, eY int) {
	row := len(mapArr)    //行
	col := len(mapArr[0]) //列
	if x < 0 || y < 0 || x >= row || y >= col {
		fmt.Println("起始坐标错误：", x, x)
		return
	}
	if eX < 0 || eY < 0 || eX >= row || eY >= col {
		fmt.Printf("%d*%d终点坐标错误：x=%d,y=%d\n", row, col, eX, eY)
		return
	}

	if mapArr[x][y] == -1 {
		fmt.Println("坐标错误：", x, x, "-墙")
		return
	}
	points := [][]int{
		{x, y},
	}
	var step = 0
	i := 0
	for len(points) > 0 {
		i++
		px, py := points[0][0], points[0][1]
		points = points[1:]
		if px == eX && py == eY { //终点
			fmt.Println("抵达终点：", px, py)
			break
		}
		step = mapArr[px][py] + 1                                          //步数
		if px-1 >= 0 && mapArr[px-1][py] == 0 && !(px-1 == x && py == y) { //上
			mapArr[px-1][py] = step //当前步数+1
			points = append(points, []int{px - 1, py})
		}
		if py-1 >= 0 && mapArr[px][py-1] == 0 && !(px == x && py-1 == y) { //左
			mapArr[px][py-1] = step //当前步数+1
			points = append(points, []int{px, py - 1})
		}
		if px+1 < row && mapArr[px+1][py] == 0 && !(px+1 == x && py == y) { //下
			mapArr[px+1][py] = step //当前步数+1
			points = append(points, []int{px + 1, py})
		}
		if py+1 < col && mapArr[px][py+1] == 0 && !(px == x && py+1 == y) { //右
			mapArr[px][py+1] = step //当前步数+1
			points = append(points, []int{px, py + 1})
		}
	}
	fmt.Println(i)
	//mapArr[x][y] = 0
}
