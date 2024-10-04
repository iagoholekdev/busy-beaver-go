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
	return 0 // valor padrão se a célula estiver vazia
}

// Move move a posição da fita para a esquerda ou para a direita
func (t *Tape) Move(dir int) {
	t.pos += dir
}

// busyBeaver simula a máquina de Turing
func busyBeaver(steps int) int {
	tape := NewTape()
	state := A
	transitionCount := 0

	for transitionCount < steps {
		switch state {
		case A:
			if tape.Read() == 0 {
				tape.Write(1) // escreve 1
				tape.Move(1)  // move para a direita
				state = B     // muda para o estado B
			} else {
				tape.Write(0) // escreve 0
				tape.Move(-1) // move para a esquerda
				state = HALT  // muda para o estado de parada
			}
		case B:
			if tape.Read() == 0 {
				tape.Write(1) // escreve 1
				tape.Move(-1) // move para a esquerda
				state = C     // muda para o estado C
			} else {
				tape.Write(1) // escreve 1
				tape.Move(1)  // move para a direita
				state = HALT  // muda para o estado de parada
			}
		case C:
			if tape.Read() == 0 {
				tape.Write(1) // escreve 1
				tape.Move(1)  // move para a direita
				state = HALT  // muda para o estado de parada
			} else {
				tape.Write(0) // escreve 0
				tape.Move(-1) // move para a esquerda
				state = A     // volta para o estado A
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
	steps := 100 // Defina o número máximo de passos permitidos
	operations := busyBeaver(steps)
	fmt.Printf("A máquina executou %d operações antes de parar\n", operations)
}
