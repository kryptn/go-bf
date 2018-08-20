package main

import (
	"reflect"
	"testing"
)

type Test struct {
	inst  string
	begin State
	end   State
}

func (tc Test) StepTest(t *testing.T) {
	machine, _ := NewMachine(tc.inst)
	machine.States = []*State{&tc.begin}
	machine.Step()
	if !reflect.DeepEqual(machine.LatestState(), &tc.end) {
		t.Logf("%s\nstart\t%+v\nend\t%+v\nexp\t%+v", tc.inst, tc.begin, machine.LatestState(), tc.end)
		t.Fail()
	}
}

func (tc Test) RunTest(t *testing.T) {
	machine, _ := NewMachine(tc.inst)
	machine.Run()

	if !reflect.DeepEqual(machine.LatestState().mem, tc.end.mem) {
		t.Logf("%s: %+v != %+v", tc.inst, machine.LatestState().mem, tc.end.mem)

		t.Fail()
	}

}

// Empty: State{[]int{0}, 0, 0}
var operatorTests = []Test{
	{
		"+ add",
		State{[]int{0}, 0, 0},
		State{[]int{1}, 0, 1},
	}, {
		"- sub",
		State{[]int{0}, 0, 0},
		State{[]int{-1}, 0, 1},
	}, {
		"> mem ptr right || new cell",
		State{[]int{1}, 0, 0},
		State{[]int{1, 0}, 1, 1},
	}, {
		"> mem ptr right || existing cell",
		State{[]int{0, 1}, 0, 0},
		State{[]int{0, 1}, 1, 1},
	}, {
		"< mem ptr left || new cell",
		State{[]int{1}, 0, 0},
		State{[]int{0, 1}, 0, 1},
	}, {
		"< mem ptr left || existing cell",
		State{[]int{0, 1}, 1, 0},
		State{[]int{0, 1}, 0, 1},
	}, {
		"[] loop start || false || bypass",
		State{[]int{0}, 0, 0},
		State{[]int{0}, 0, 2},
	}, {
		"[] loop start || true || enter",
		State{[]int{1}, 0, 0},
		State{[]int{1}, 0, 1},
	}, {
		"[] loop end || false || exit",
		State{[]int{0}, 0, 1},
		State{[]int{0}, 0, 2},
	}, {
		"[] loop end || true || reenter",
		State{[]int{1}, 0, 1},
		State{[]int{1}, 0, 1},
	},
}

func TestOperators(t *testing.T) {
	for _, test := range operatorTests {
		t.Run(test.inst, test.StepTest)
	}
}

var runTests = []Test{
	{
		"+++++ add to 5", State{},
		State{[]int{5}, 0, 0},
	}, {
		"-----+++++ -5+5", State{},
		State{[]int{0}, 0, 0},
	}, {
		"+++[>+++++<-] 3 * 5 || loop", State{},
		State{[]int{0, 15}, 0, 0},
	}, {
		"+++[>+++++[>++++<-]<-] 3 * 5 * 4 || nested loop", State{},
		State{[]int{0, 0, 60}, 0, 0},
	},
}

func TestMachine_Run(t *testing.T) {
	for _, test := range runTests {
		t.Run(test.inst, test.RunTest)
	}
}
