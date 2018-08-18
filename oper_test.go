package main

import "testing"

func memEquals(n int) func(*State) bool {
	return func(s *State) bool {
		return s.mem[s.mptr] == n
	}
}

func memPtrEquals(n int) func(*State) bool {
	return func(s *State) bool {
		return s.mptr == n
	}
}

func InstPtrEquals(n int) func(*State) bool {
	return func(s *State) bool {
		return s.iptr == n
	}
}

func and(fn ... func(*State) bool) func(*State) bool {
	return func(state *State) bool {
		for _, f := range fn {
			if !f(state) {
				return false
			}
		}
		return true
	}
}

type Test struct {
	state  *State
	verify func(*State) bool
}

func (ut Test) StepTest(i int, t *testing.T) {
	ut.state.Step()
	if !ut.verify(ut.state) {
		t.Logf("Failed on %s: %d", t.Name(), i)
		t.Fail()
	}
}

func stepTestWith(tests []Test, t *testing.T) {
	for i, test := range tests {
		test.StepTest(i, t)
	}
}

var operAddTests = []Test{
	{NewState("+"), memEquals(1)},
	{NewState("+").WithMem(1), memEquals(2)},
}

func TestOperAdd(t *testing.T) {
	stepTestWith(operAddTests, t)
}

var operSubTests = []Test{
	{NewState("-"), memEquals(-1)},
	{NewState("-").WithMem(1), memEquals(0)},
}

func TestOperSub(t *testing.T) {
	stepTestWith(operSubTests, t)
}

var memPtrTests = []Test{
	{
		NewState(">"),
		memPtrEquals(1),
	}, {
		NewState("<"),
		memPtrEquals(0),
	}, {
		NewState(">").WithMem(0, 0).WithMemPtr(1),
		memPtrEquals(2),
	},
}

func TestMemPtr(t *testing.T) {
	stepTestWith(memPtrTests, t)
}

var loopTests = []Test{
	{
		NewState("+[>++<]").WithInstPtr(6).WithMem(1, 2),
		InstPtrEquals(2),
	}, {
		NewState("+[>++<]").WithInstPtr(6).WithMem(0, 2),
		InstPtrEquals(7),
	}, {
		NewState("+[>+[>+++<]+<]").WithInstPtr(10).WithMem(1),
		InstPtrEquals(5),
	},
}

func TestLoop(t *testing.T) {
	stepTestWith(loopTests, t)
}
