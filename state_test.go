package main

import (
	"reflect"
	"testing"
)

func complexState() *State {
	return &State{
		mem:  []int{1, 2, 3, 4, 5},
		mptr: 2,
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

func all(s1, s2 *State, cmps ...func(s1, s2 *State) bool) bool {
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

func TestNewMachine__InvalidState(t *testing.T) {
	_, err := NewMachine("[")

	if err == nil {
		t.Fail()
	}

}
