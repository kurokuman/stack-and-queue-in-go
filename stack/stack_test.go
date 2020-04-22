package stack

import (
	"testing"
)

var sizeTests = []struct {
	input    []string
	expected int
}{
	{
		input:    []string{"a", "b", "c"},
		expected: 3,
	},
}

func TestSize(t *testing.T) {
	stack := Stack{}
	for _, tt := range sizeTests {
		for _, input := range tt.input {
			stack.Push(input)
		}
		size := stack.Size()
		if size != tt.expected {
			t.Fatalf("FAIL: expected %d, got %d", tt.expected, size)
		}
		t.Logf("PASS: PoP %q", tt.input)
	}
}

var pushTests = []struct {
	input    []string
	expected string
}{
	{
		input:    []string{"a", "b", "c"},
		expected: "c",
	},
}

func TestPush(t *testing.T) {
	stack := Stack{}
	for _, tt := range pushTests {
		for _, input := range tt.input {
			stack.Push(input)
		}
		top := stack.top.value
		if top != tt.expected {
			t.Fatalf("FAIL: expected %s, got %s", tt.expected, tt.input[len(tt.input)-1])
		}
		t.Logf("PASS: Push %q", tt.input)
	}
}

var popTests = []struct {
	input    []string
	expected string
}{
	{
		input:    []string{"a", "b", "c"},
		expected: "c",
	},
}

func TestPop(t *testing.T) {
	stack := Stack{}
	for _, tt := range popTests {
		for _, input := range tt.input {
			stack.Push(input)
		}
		pop, _ := stack.Pop()
		if pop != tt.expected {
			t.Fatalf("FAIL: expected %s, got %s", tt.expected, tt.input[len(tt.input)-1])
		}
		t.Logf("PASS: PoP %q", tt.input)
	}
}
