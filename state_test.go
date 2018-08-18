package main

import (
	"testing"
	"reflect"
)

func complexState() *State {
	return &State{
		mem:  []int{1, 2, 3, 4, 5},
		mptr: 2,
		inst: []rune("+++[>+++<]+."),
		iptr: 0,
	}
}

func memEq(s1, s2 *State) bool {
	return reflect.DeepEqual(s1.mem, s2.mem)
}

func iptrEq(s1, s2 *State) bool {
	return s1.iptr == s2.iptr
}

func mptrEq(s1, s2 *State) bool {
	return s1.mptr == s2.mptr
}

func all(s1, s2 *State, cmps ... func(s1, s2 *State) bool) bool {
	for _, cmp := range cmps {
		if !cmp(s1, s2) {
			return false
		}
	}
	return true
}

func TestState_Copy(t *testing.T) {
	state := complexState()
	sc := state.Copy()

	if !all(state, sc, memEq, iptrEq, mptrEq) {
		t.Fail()
	}
}

func TestState_WithMem(t *testing.T) {
	state := complexState()
	sc := state.WithMem(3, 2, 1)
	if sc.mem[0] != 3 && sc.mem[1] != 2 && sc.mem[2] != 1 {
		t.Fail()
	}
	if !all(state, sc, iptrEq, mptrEq) {
		t.Fail()
	}

}

func TestState_WithMemPtr(t *testing.T) {
	state := complexState()
	sc := state.WithMemPtr(3)
	if sc.mptr != 3 {
		t.Fail()
	}
	if !all(state, sc, memEq, iptrEq) {
		t.Fail()
	}
}

func TestState_WithInstPtr(t *testing.T) {
	state := complexState()
	sc := state.WithInstPtr(4)
	if sc.iptr != 4 {
		t.Fail()
	}
	if !all(state, sc, memEq, mptrEq) {
		t.Fail()
	}
}

func TestState_Withs(t *testing.T) {
	a := complexState().WithMem(1,2,3).WithMemPtr(4)
	b := complexState().WithMemPtr(4).WithMem(1,2,3)

	if !all(a, b, memEq, iptrEq, mptrEq) {
		t.Fail()
	}
}
