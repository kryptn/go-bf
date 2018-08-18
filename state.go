package main

import (
	"log"
	"fmt"
	"strings"
)

type Machine struct {
	States []*State
	inst   []rune
}

func (m *Machine) Next() {
	latest := m.States[len(m.States)-1]

	latest.Step()
}

type State struct {
	mem  []int
	mptr int

	inst []rune
	iptr int
}

func (s *State) Copy() *State {
	memCopy := make([]int, len(s.mem))
	copy(s.mem, memCopy)
	return &State{memCopy, s.mptr, s.inst, s.iptr}

}

func (s *State) WithMem(m ... int) *State {
	c := s.Copy()
	c.mem = m
	//c.mem = append([]int{}, m...)
	return c
}

func (s *State) WithMemPtr(mptr int) *State {
	c := s.Copy()
	c.mptr = mptr
	return c
}

func (s *State) WithInstPtr(iptr int) *State {
	c := s.Copy()
	c.iptr = iptr
	return c
}

func (s *State) Print(context string) {
	fmt.Printf("\n\ndebug -- %s\n", context)
	fmt.Printf("\t%+v\n", s.mem)
	fmt.Printf("mem ptr: %d\n", s.mptr)
	fmt.Printf("%s\n", string(s.inst))
	fmt.Printf("%s^\n", strings.Repeat(" ", s.iptr))
	fmt.Printf("%d\n\n", s.iptr)
}

func (s *State) Step() {
	//s.Print("Before")
	if s.iptr < len(s.inst) {
		operator := operMap[s.inst[s.iptr]]
		operator(s)
		s.iptr += 1
		//s.Print("During")
	}
}

func NewState(instructions string) *State {
	cleaned := CleanInput(instructions)
	if ! Validate(cleaned) {
		log.Fatalf("Invalid loop state detected")
	}
	return &State{[]int{0}, 0, []rune(cleaned), 0}

}

func (s *State) Run() {
	for s.iptr = 0; s.iptr < len(s.inst); s.iptr++ {
		opcode := s.inst[s.iptr]
		oper := operMap[opcode]
		oper(s)
	}
	fmt.Print("\n")
}
