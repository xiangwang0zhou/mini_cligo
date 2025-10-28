package main

import (
	"testing"

	"github.com/xiangwang0zhou/mini_cligo/commands"
)

func TestClearInput(t *testing.T) {
	type testCases struct {
		calls string
		input string
		wants []string
	}
	cases := []testCases{
		{
			calls: "common string",
			input: "Initlizations contents in theres	",
			wants: []string{
				"initlizations", "contents", "in", "theres",
			},
		},
		{
			calls: "all plus",
			input: "INITLIZATIONS CONTENTS IN THERE",
			wants: []string{
				"initlizations", "contents", "in", "there",
			},
		},
	}

	for _, ca := range cases {
		t.Run(ca.calls, func(t *testing.T) {
			actualWords := commands.ClearInput(ca.input)
			expectWords := ca.wants
			if len(actualWords) != len(expectWords) {
				t.Errorf("The length of %v is not equal to %v", actualWords, expectWords)
			}
			for i := range actualWords {
				if actualWords[i] != expectWords[i] {
					t.Errorf("The value of %v is not equal to %v ",
						actualWords[i], expectWords[i])
				}
			}
		})
	}
}
