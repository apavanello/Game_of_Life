package game

import (
	"fmt"
	"math/rand"
	"time"
)

type Field bool

type Board [][]Field

type Pos struct {
	x int
	y int
}

type screen [][]string

func Start() {
	steps := int64(0)
	currentBoard := Board{}
	currentBoard.create(40, 60)
	s := currentBoard.initScreen()
	currentBoard.genLife()

	for {
		currentBoard.print(s, steps)
		time.Sleep(10 * time.Millisecond)
		currentBoard.step()
		steps++
	}
}

func (b *Board) genLife() {

	row := len(*b)
	col := len((*b)[0])
	seed := int64(time.Now().Nanosecond())
	rand.Seed(seed)

	//25% of the cells will be alive
	qnt := row * col / 5

	for i := 0; i < qnt; i++ {
		x := rand.Int() % row
		y := rand.Int() % col
		(*b)[x][y] = true
	}
}

func (b *Board) create(row int, col int) {

	*b = make(Board, row)

	for i := 0; i < row; i++ {
		(*b)[i] = make([]Field, col)
	}
}

func (b *Board) print(s screen, steps int64) {

	line := "\033[H\033[2J"
	dead := "⬛"
	live := "⬜"

	fmt.Println(line)

	for i, row := range *b {
		for j, col := range row {
			if col {
				s[i][j] = live
			} else {
				s[i][j] = dead
			}
		}
		s[i][len(row)] = "\n"
	}
	fmt.Printf("%s", s)
	fmt.Printf("\n\n\nSteps: %v ", steps)
}

func (b *Board) step() {

	nextBoard := Board{}
	nextBoard.create(len(*b), len((*b)[0]))

	for i, _ := range *b {
		for j, _ := range (*b)[i] {
			nextBoard[i][j] = b.getNextState(i, j)
		}
	}

	*b = nextBoard
}

func (b *Board) getNextState(row int, col int) Field {

	p := Pos{}
	neigbor := 0
	ret := (*b)[row][col]

	for i := -1; i <= 1; i++ {
		p.x = row + i

		for j := -1; j <= 1; j++ {
			p.y = col + j

			if i == 0 && j == 0 {
				continue
			}
			if b.isBorder(p) {
				continue
			}

			if (*b)[p.x][p.y] {
				neigbor++
			}

		}
	}

	if neigbor < 2 || neigbor > 3 {
		ret = false
	}

	if neigbor == 3 {
		ret = true
	}

	//	fmt.Printf("{%v,%v} -> ", row, col)
	//	fmt.Printf("%v = %v \n", neigbor, ret)

	return ret
}

func (b *Board) isBorder(p Pos) bool {

	if p.x < 0 || p.x >= len(*b) {
		return true
	}

	if p.y < 0 || p.y >= len((*b)[0]) {
		return true
	}

	return false
}

func (b *Board) initScreen() (s screen) {
	s = make([][]string, len(*b))
	for i := 0; i < len(*b); i++ {
		s[i] = make([]string, len((*b)[0])+1)
	}
	return
}
