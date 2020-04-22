package queue

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
	queue := Queue{}
	for _, tt := range sizeTests {
		for _, input := range tt.input {
			queue.Enqueue(input)
		}
		size := queue.Size()
		if size != tt.expected {
			t.Fatalf("FAIL: expected %d, got %d", tt.expected, size)
		}
		t.Logf("PASS: Size %q", tt.input)
	}
}

var EnqueueTest = []struct {
	input    []string
	expected string
}{
	{
		input:    []string{"a", "b", "c"},
		expected: "a",
	},
}

func TestEnqueue(t *testing.T) {
	queue := Queue{}
	for _, tt := range EnqueueTest {
		for _, input := range tt.input {
			queue.Enqueue(input)
		}
		top := queue.top.value
		if top != tt.expected {
			t.Fatalf("FAIL: expected %s, got %s", tt.expected, tt.input[len(tt.input)-1])
		}
		t.Logf("PASS: Enqueue %q", tt.input)
	}
}

var dequeueTests = []struct {
	input    []string
	expected string
}{
	{
		input:    []string{"a", "b", "c"},
		expected: "a",
	},
}

func TestDequeue(t *testing.T) {
	queue := Queue{}
	for _, tt := range dequeueTests {
		for _, input := range tt.input {
			queue.Enqueue(input)
		}
		pop, _ := queue.Dequeue()
		if pop != tt.expected {
			t.Fatalf("FAIL: expected %s, got %s", tt.expected, tt.input[len(tt.input)-1])
		}
		t.Logf("PASS: Dequeue %q", tt.input)
	}
}
