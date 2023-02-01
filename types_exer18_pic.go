package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	rows := make([][]uint8, dy)
	for y := range rows {
		rows[y] = make([]uint8, dx)
		for x := range rows[y] {
			rows[y][x] = uint8(x*y)
		}
	}
	return rows
}

func main() {
	pic.Show(Pic)
}

