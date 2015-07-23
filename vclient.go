package main

import (
			"net"
			"bufio"
			"github.com/nsf/termbox-go"
			"strconv"
		)

type game struct {
	bY, bX	int
	board		[3][3]int
	cX, cY	int
	player	int
	server	int
	myMove	bool
}
const SERVER = "games.recurse.com"

//	Port number for games
const (
		TICTACTOE	=	(iota + 7) * 1000
		SUPERTICTACTOE
		CHECKERS
		)
const (cd = termbox.ColorDefault
		bg = 0x07
	)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetOutputMode(termbox.Output216)
	termbox.Clear(termbox.ColorDefault, bg)
	termbox.Flush()
	g := &game{}

	conn, err := net.Dial("tcp", SERVER + ":7000")	// TODO Use constants for game number
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	m := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	//as, _, _ := bufio.NewReader(conn).ReadLine()
	as, _, _ := m.ReadLine()
	pn, _ := strconv.Atoi(string(as[:]))
	g.player = pn
	// Replace with function or something
	g.server = 1
	termbox.SetCell(0, 0, '2', cd, cd)
	g.myMove = false
	if g.player == 1 {
		termbox.SetCell(0, 0, '1', cd, cd)
		g.myMove = true
		g.server = 2
	}
if !g.myMove {
			servMove, _, _ := m.ReadLine()
			sx, sy, _ := toSpace(servMove)
			g.board[sy][sx] = g.server
			g.myMove = true

		}
	g.drawBoard(3, 3, 0x00)
	g.drawBoard(2, 2, 0x05)
	termbox.Flush()
	s := true
	for s {
		switch e := termbox.PollEvent(); e.Type {
		case termbox.EventKey:
			termbox.Clear(cd, bg)
			switch e.Key {
			case termbox.KeyArrowUp:
				g.cY--
			case termbox.KeyArrowDown:
				g.cY++
			case termbox.KeyArrowRight:
				g.cX++
			case termbox.KeyArrowLeft:
				g.cX--
			case termbox.KeySpace:
				g.board[g.cX][g.cY] = g.player

				// Replace with function or something
				if g.myMove {
					m.WriteString(toString(g.cY, g.cX))
					m.Flush()
					g.myMove = false
				}
			case termbox.KeyEsc:
				s = false
			}
		}
		g.drawBoard(3, 3, 0x00)
		g.drawBoard(2, 2, 0x05)
		termbox.Flush()

		if !g.myMove {
			servMove, _, _ := m.ReadLine()
			sx, sy, state := toSpace(servMove)
			if state != 0 {
				break
			}
			g.board[sy][sx] = g.server
			g.myMove = true
			g.drawBoard(3, 3, 0x00)
			g.drawBoard(2, 2, 0x05)
			termbox.Flush()

		}
	}
}

func toSpace(in []byte) (int, int, int) {
	if	len(in) == 1 {
		return 0, 0, int(in[0] - 48)
	}
	return int(in[0] - 48), int(in[2] - 48), 0
}

func toString(x, y int) (string) {
	return strconv.Itoa(x) + " " + strconv.Itoa(y) + "\n"
}
/*
func gameLogicInit() {
	conn, err := net.Dial("tcp", SERVER + ":7000")
	if err != nil {
		panic(err)
	}

	p, _, _:= bufio.NewReader(conn).ReadLine()
	if string(p) == "2" {
		temp, _, _ := bufio.NewReader(os.Stdin).ReadLine()
		fmt.Println(temp)
	}
	for {
		q := bufio.NewWriter(conn)
		//q.WriteString(getString(g))
		q.Flush()
		temp, _, _ := bufio.NewReader(os.Stdin).ReadLine()
		fmt.Println(temp)
	}
}
*/

func (g game) drawBoard(x, y int, c termbox.Attribute) {
	s := 30
	for rx, vr := range g.board {
		for ry, v := range vr {
			if v == 1 {
				// Temporary below
				if int(c) != 0 {
					drawX(x + (rx * 10) + 1, y + (ry * 10) + 1, 0x79)
				}else{
					drawX(x + (rx * 10) + 1, y + (ry * 10) + 1, c)
				}
			}
			if v == 2 {
				// Temporary below
				if int(c) != 0 {
					drawO(x + (rx * 10) + 1, y + (ry * 10) + 1, 0x79)
				}else{
					drawO(x + (rx * 10) + 1, y + (ry * 10) + 1, c)
				}
			}
		}
	}
	for i := 0; i < s; i++ {
		for z := 0; z < 2; z++{
			termbox.SetCell(i + x, (z * 10) + 10 + y, ' ', cd, c)
			termbox.SetCell((z * 10) + 10 + x, i + y, ' ', cd, c)
		}
	}
	termbox.SetCell(x + (g.cX * 10) + 5, y + (g.cY * 10) + 5, ' ', cd, c)
}

func drawX(x, y int, c termbox.Attribute) {
	s := 8
	for i := 0; i < s; i++ {
		termbox.SetCell(x + i, y + i, ' ', cd, c)
		termbox.SetCell(x - i + s, y + i, ' ', cd, c)
	}
}

func drawO(x, y int, c termbox.Attribute) {
	s := 8
	for i := 0; i < s; i++ {
		termbox.SetCell(x, y + i, ' ', cd, c)
		termbox.SetCell(x + s, y + i, ' ', cd, c)
		termbox.SetCell(x + i, y, ' ', cd, c)
		termbox.SetCell(x + i, y + s, ' ', cd, c)

	}
}
