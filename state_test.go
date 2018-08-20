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

func TestState_Copy(t *testing.T) {
	state := complexState()
	sc := state.Copy()

	if !reflect.DeepEqual(state, sc) {
		t.Fail()
	}
}

func TestNewMachine__InvalidLoopState(t *testing.T) {
	_, err := NewMachine("[")

	if err == nil {
		t.Fail()
	}

}
