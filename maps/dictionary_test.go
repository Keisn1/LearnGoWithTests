package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dict := Dictionary{
		"test":    "this is just a test",
		"present": "present word",
	}

	t.Run("Word present", func(t *testing.T) {
		got, err := dict.Search("test")
		want := "this is just a test"
		assertNoError(t, err)
		assertStrings(t, got, want)
	})

	t.Run("Word not present", func(t *testing.T) {
		word := "not present"
		_, gotErr := dict.Search(word)

		assertError(t, gotErr, ErrNotFound)
	})

	t.Run("Add a word", func(t *testing.T) {
		word := "New word"
		def := "definition of new word"
		err := dict.Add(word, def)
		assertNoError(t, err)
		assertDefinition(t, dict, word, def)
	})

	t.Run("Don't Add word, rahter update", func(t *testing.T) {
		word := "present"
		def := "present word"

		err := dict.Add(word, def)
		assertError(t, err, ErrWordPresent)
		assertDefinition(t, dict, word, def)
	})

	t.Run("Update exist word", func(t *testing.T) {
		word := "newWord"
		def := "old definition"
		dict := Dictionary{
			word: def,
		}
		newDef := "new definition"

		dict.Update(word, newDef)

		assertDefinition(t, dict, word, newDef)
	})

	t.Run("Don't add word", func(t *testing.T) {
		word := "NonExisting"
		def := "old definition"

		err := dict.Update(word, def)
		assertError(t, err, ErrWordDoesNotExist)
	})

	t.Run("Delete word", func(t *testing.T) {
		word := "test"

		err := dict.Delete(word)

		_, err = dict.Search(word)
		if err != ErrNotFound {
			t.Errorf("Err was not ErrNotFound; %q", err)
		}
	})

	t.Run("Delete present word", func(t *testing.T) {
		word := "test"

		_ = dict.Delete(word)
		assertErrNotFound(t, dict, word)
	})

	t.Run("Error attempt deleting non present word ", func(t *testing.T) {
		word := "not present"

		err := dict.Delete(word)
		assertError(t, err, ErrWordToDeleteNotPresent)
	})
}

func assertErrNotFound(t testing.TB, d Dictionary, w string) {
	_, err := d.Search(w)
	if err != ErrNotFound {
		t.Errorf("Err was not ErrNotFound; %q", err)
	}
}

func assertDefinition(t testing.TB, d Dictionary, w, def string) {
	t.Helper()
	got, err := d.Search(w)
	assertNoError(t, err)
	assertStrings(t, got, def)
}

func assertError(t testing.TB, gotErr, wantErr error) {
	t.Helper()
	if gotErr != wantErr {
		t.Fatalf("Got error: %q. Want error: %q", gotErr, wantErr)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("Got an unwanted error")
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Search(test) = \"%v\"; want \"%v\"", got, want)
	}
}
