package main

import (
	"reflect"
	"testing"
)

// type SpyFn struct {
// 	calls int
// }

// func (s *SpyFn) fn(string) {
// 	s.calls++
// }

func TestWalk(t *testing.T) {
	testCases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"}},
		{"struct with one string field",
			struct {
				Name string
				City string
			}{"Chris", "New York"},
			[]string{"Chris", "New York"}},
		{"struct with one string field",
			struct {
				Name string
				Age  int
				City string
			}{"Chris", 33, "New York"},
			[]string{"Chris", "New York"}},
		{"struct with nested struct",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"}},

		{"pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"}},
		{"Slices",
			[]Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			[]string{"London", "Reykjavik"}},
		{"Arrays",
			[2]Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			[]string{"London", "Reykjavik"}},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			var got []string
			walk(tc.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, tc.ExpectedCalls) {
				t.Errorf("got %q, want %q", got, tc.ExpectedCalls)
			}
		})

	}

	t.Run("with Maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}
		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})
		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)
		go func() {
			aChannel <- Profile{33, "London"}
			aChannel <- Profile{22, "Berlin"}
			close(aChannel)
		}()

		var got []string
		want := []string{"London", "Berlin"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with functions", func(t *testing.T) {
		aFunc := func() (Profile, Profile) {
			return Profile{33, "London"}, Profile{22, "Berlin"}
		}

		var got []string
		want := []string{"London", "Berlin"}

		walk(aFunc, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("%q doesn't contain %q", haystack, needle)
	}
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

// func TestWalk(t *testing.T) {
// 	x := "first"
// 	spy := SpyFn{}
// 	Walk(x, spy.fn)
// 	want := 1
// 	got := spy.calls
// 	if got != want {
// 		t.Errorf("Spy.calls = \"%v\"; want \"%v\"", got, want)
// 	}
// }
