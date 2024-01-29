package generics

import (
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		intStack := new(Stack[int32])

		AssertTrue(t, intStack.IsEmpty())
		intStack.Push(2)
		AssertFalse(t, intStack.IsEmpty())

		val, popped := intStack.Pop()
		AssertTrue(t, popped)
		AssertEqual(t, val, 2)
	})

	t.Run("string stack", func(t *testing.T) {
		stringStack := new(Stack[string])

		AssertTrue(t, stringStack.IsEmpty())
		stringStack.Push("asdf")
		AssertFalse(t, stringStack.IsEmpty())

		val, popped := stringStack.Pop()
		AssertTrue(t, popped)
		AssertEqual(t, val, "asdf")
	})
}

func TestAssertFunctions(t *testing.T) {
	t.Run("Asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertEqual(t, "ich", "ich")
		AssertNotEqual(t, 1, 2)
	})
}

func AssertEqual[T comparable](t *testing.T, a, b T) {
	t.Helper()
	if a != b {
		t.Errorf("a = %v not equal b = %v", a, b)
	}
}

func AssertNotEqual[T comparable](t *testing.T, a, b T) {
	t.Helper()
	if a == b {
		t.Errorf("a = %v equal b = %v", a, b)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got is false")
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got is true")
	}
}

type Stack[T any] struct {
	values []T
}

func (s *Stack[T]) Push(val T) {
	s.values = append(s.values, val)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	val := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return val, true
}

// type StackOfStrings struct {
// 	values []string
// }

// func (s *StackOfStrings) Push(val string) {
// 	s.values = append(s.values, val)
// }

// func (s *StackOfStrings) IsEmpty() bool {
// 	return len(s.values) == 0
// }

// func (s *StackOfStrings) Pop() (string, bool) {
// 	if s.IsEmpty() {
// 		return "", false
// 	}
// 	val := s.values[len(s.values)-1]
// 	s.values = s.values[:len(s.values)-1]
// 	return val, true
// }
