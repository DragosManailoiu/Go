package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	//make a board

	const (
		width  = 50
		height = 10

		cellEmpty = ' '
		cellBall  = '⚽'
		maxFrame  = 1200
		speed     = time.Second / 20
	)

	var (
		cell   rune
		vx, vy = 1, 1
		px, py int
	)

	board := make([][]bool, width)
	for row := range board {
		board[row] = make([]bool, height)
	}

	screen.Clear()

	for i := 0; i < maxFrame; i++ {
		px += vx
		py += vy

		if px <= 0 || px >= width-1 {
			vx *= -1
		}
		if py <= 0 || py >= height-1 {
			vy *= -1
		}
		for y := range board[0] {
			for x := range board {
				board[x][y] = false
			}
		}

		board[px][py] = true

		buf := make([]rune, 0, width*height)
		buf = buf[:0]

		for y := range board[0] {
			for x := range board {
				cell = cellEmpty
				if board[x][y] {
					cell = cellBall
				}
				buf = append(buf, cell, ' ')

			}
			buf = append(buf, '\n')
		}

		screen.MoveTopLeft()
		fmt.Print(string(buf))

		time.Sleep(speed)
	}
}
