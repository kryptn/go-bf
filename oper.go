package main

import (
	"bufio"
	"log"
)

type Operator func([]rune, *State)

func OperatorMap(reader *bufio.Reader, writer *bufio.Writer) map[rune]Operator {
	return map[rune]Operator{
		'+': modifyCurrentMemory(1),
		'-': modifyCurrentMemory(-1),
		'>': IncMemPtr,
		'<': DecMemPtr,
		'.': Output(writer),
		',': Input(reader),
		'[': seekForPair(1),
		']': seekForPair(-1),
	}
}

func modifyCurrentMemory(n int) Operator {
	return func(_ []rune, s *State) {
		s.mem[s.mptr] += n
	}
}

func IncMemPtr(_ []rune, s *State) {
	s.mptr += 1
	if s.mptr >= len(s.mem) {
		s.mem = append(s.mem, 0)
	}
}

func DecMemPtr(_ []rune, s *State) {
	if s.mptr == 0 {
		s.mem = append([]int{0}, s.mem...)
	} else {
		s.mptr -= 1
	}
}

func Output(writer *bufio.Writer) Operator {
	return func(_ []rune, s *State) {
		writer.WriteString(string(rune(s.mem[s.mptr])))
	}
}

func Input(reader *bufio.Reader) Operator {
	return func(_ []rune, s *State) {
		r, _, err := reader.ReadRune()
		if err != nil {
			log.Fatalf("hmm")
		}
		s.mem[s.mptr] = int(r)
	}

}

var goBackMap = map[rune]int{'[': -1, ']': 1}
var goForwardMap = map[rune]int{'[': 1, ']': -1}

func seekForPair(dir int) Operator {
	var adjMap map[rune]int
	var other rune
	if dir == -1 {
		adjMap, other = goBackMap, '['
	} else {
		adjMap, other = goForwardMap, ']'
	}
	return func(inst []rune, s *State) {
		memory := s.mem[s.mptr]
		if memory != 0 && dir == 1 {
			// when [ and not zero, noop
			return
		}

		if memory == 0 && dir == -1 {
			// when ] and zero, noop
			return
		}

		for level := 0; inst[s.iptr] != other || level != 1; {
			if adj, ok := adjMap[inst[s.iptr]]; ok {
				level += adj
			}
			s.iptr += dir
		}
	}
}
