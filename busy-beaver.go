package main

import (
	"fmt"
)

const (
	A = iota
	B
	C
	HALT
)

type Tape struct {
	cells map[int]int
	pos   int
}

func NewTape() *Tape {
	return &Tape{cells: make(map[int]int), pos: 0}
}

func (t *Tape) Write(val int) {
	t.cells[t.pos] = val
}

func (t *Tape) Read() int {
	val, ok := t.cells[t.pos]
	if ok {
		return val
	}
	return 0
}

func (t *Tape) Move(dir int) {
	t.pos += dir
}

func busyBeaver(steps int) int {
	tape := NewTape()
	state := A
	transitionCount := 0

	for transitionCount < steps {
		switch state {
		case A:
			if tape.Read() == 0 {
				tape.Write(1)
				tape.Move(1)
				state = B
			} else {
				tape.Write(0)
				tape.Move(-1)
				state = HALT
			}
		case B:
			if tape.Read() == 0 {
				tape.Write(1)
				tape.Move(-1)
				state = C
			} else {
				tape.Write(1)
				tape.Move(1)
				state = HALT
			}
		case C:
			if tape.Read() == 0 {
				tape.Write(1)
				tape.Move(1)
				state = HALT
			} else {
				tape.Write(0)
				tape.Move(-1)
				state = A
			}
		}

		transitionCount++

		if transitionCount > steps {
			break
		}
	}
	return transitionCount
}

func main() {
	steps := 100
	operations := busyBeaver(steps)
	fmt.Printf("A máquina executou %d operações antes de parar\n", operations)
}
