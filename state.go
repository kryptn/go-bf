package main

import (
	"errors"
)

type Machine struct {
	States []*State
	inst   []rune

	//reader *bufio.Reader
	//writer *bufio.Writer

	operators map[rune]Operator
}

func (m *Machine) LatestOperator() Operator {
	latest := m.LatestState()
	return m.operators[m.inst[latest.iptr]]
}

func (m *Machine) Step() {
	current := m.LatestState()
	operator := m.LatestOperator()
	m.States = append(m.States, current.Apply(m.inst, operator))
}

func (m *Machine) LatestState() *State {
	return m.States[len(m.States)-1]
}

func (m *Machine) Run() {
	latest := m.LatestState()
	for latest.iptr < len(m.inst) {
		m.Step()
		latest = m.LatestState()
	}
}

func NewMachine(instructions string) (*Machine, error) {
	cleaned := CleanInput(instructions)
	if !Validate(cleaned) {
		return nil, errors.New("Invalid loop state detected")
	}

	return &Machine{
		[]*State{
			{[]int{0}, 0, 0}},
		[]rune(cleaned),
		OperatorMap(nil, nil),
	}, nil
}

type State struct {
	mem        []int
	mptr, iptr int
}

func (s *State) Apply(inst []rune, op Operator) *State {
	next := s.Copy()
	op(inst, next)
	next.iptr += 1
	return next
}

func (s *State) Copy() *State {
	return &State{
		append([]int{}, s.mem...),
		s.mptr,
		s.iptr,
	}
}

//func (s *State) Print(context string) {
//	fmt.Printf("\n\ndebug -- %s\n", context)
//	fmt.Printf("\t%+v\n", s.mem)
//	fmt.Printf("mem ptr: %d\n", s.mptr)
//	fmt.Printf("%s\n", string(s.inst))
//	fmt.Printf("%s^\n", strings.Repeat(" ", s.iptr))
//	fmt.Printf("%d\n\n", s.iptr)
//}
