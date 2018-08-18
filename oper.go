package main

import (
	"fmt"
	"bufio"
	"os"
)

const OPER_ADD = '+'
const OPER_SUB = '-'
const OPER_IPTR = '>'
const OPER_DPTR = '<'
const OPER_OUTP = '.'
const OPER_INP = ','
const OPER_BEGIN = '['
const OPER_END = ']'

type Operator func(*State)

var operMap = map[rune]Operator{
	OPER_ADD:   modifyCurrent(1),
	OPER_SUB:   modifyCurrent(-1),
	OPER_IPTR:  IncMemPtr,
	OPER_DPTR:  DecMemPtr,
	OPER_OUTP:  Output,
	OPER_INP:   Input,
	OPER_BEGIN: Noop,
	OPER_END:   LoopEnd,
}

func Noop(s *State) {}

func modifyCurrent(n int) func(*State) {
	return func(s *State) {
		s.mem[s.mptr] += n
	}
}

func IncMemPtr(s *State) {
	s.mptr += 1
	if s.mptr+1 > len(s.mem) {
		s.mem = append(s.mem, 0)
	}
}

func DecMemPtr(s *State) {
	if s.mptr == 0 {
		s.mem = append([]int{0}, s.mem...)
	} else {
		s.mptr -= 1
	}
}

func Output(s *State) {
	fmt.Print(rune(s.mem[s.mptr]))
}

func Input(s *State) {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	if len(text) > 0 {
		s.mem[s.mptr] = int([]rune(text)[0])
	}
}

var heightMap = map[rune]int{'[': -1, ']': 1}

func LoopEnd(s *State) {
	if s.mem[s.mptr] != 0 {
		for height := 0; s.inst[s.iptr] != '[' || height != 1; {
			if adj, ok := heightMap[s.inst[s.iptr]]; ok {
				height += adj
			}
			s.iptr -= 1
		}
	}
}
