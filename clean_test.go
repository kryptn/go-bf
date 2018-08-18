package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	if CleanInput("  [  ]  .  , wef < erwaslk > c +  -") != "[].,<>+-" {
		t.Fail()
	}
}

type validationTest struct {
	test     string
	expected bool
}

var validationTests = []validationTest{
	{"", true},
	{"[", false},
	{"[]]", false},
	{"[[]", false},
	{"[][][][][]", true},
	{"[[[][]][]]", true},
	{"[[[[][[[]]][]]]]", true},
	{"[[[[][[[]]]][]]]]", false},
}

func TestValidate(t *testing.T) {
	for i, test := range validationTests {
		if test.expected != Validate(test.test) {
			t.Logf("TestValidate %d: Expected %v for %s", i, test.expected, test.test)
			t.Fail()
		}
	}
}
