package main

import (
	"os"
	"testing"
)

func resetDataFile(t *testing.T) func() {
	t.Helper()
	prev, err := os.ReadFile(dataFile)
	existed := err == nil
	return func() {
		if existed {
			os.WriteFile(dataFile, prev, 0644)
		} else {
			os.Remove(dataFile)
		}
	}
}

func TestNextID(t *testing.T) {
	cases := []struct {
		todos []Todo
		want  int
	}{
		{[]Todo{}, 1},
		{[]Todo{{ID: 1}, {ID: 3}}, 4},
		{[]Todo{{ID: 5}}, 6},
	}
	for _, c := range cases {
		if got := nextID(c.todos); got != c.want {
			t.Errorf("nextID(%v) = %d, want %d", c.todos, got, c.want)
		}
	}
}

func TestSaveAndLoad(t *testing.T) {
	restore := resetDataFile(t)
	defer restore()

	todos := []Todo{
		{ID: 1, Title: "牛乳を買う", Done: false},
		{ID: 2, Title: "掃除する", Done: true},
	}

	if err := save(todos); err != nil {
		t.Fatal(err)
	}

	got, err := load()
	if err != nil {
		t.Fatal(err)
	}

	if len(got) != len(todos) {
		t.Fatalf("len = %d, want %d", len(got), len(todos))
	}
	for i, want := range todos {
		if got[i].ID != want.ID || got[i].Title != want.Title || got[i].Done != want.Done {
			t.Errorf("todos[%d] = %+v, want %+v", i, got[i], want)
		}
	}
}

func TestLoadEmpty(t *testing.T) {
	restore := resetDataFile(t)
	defer restore()
	os.Remove(dataFile)

	todos, err := load()
	if err != nil {
		t.Fatal(err)
	}
	if len(todos) != 0 {
		t.Errorf("expected empty, got %v", todos)
	}
}
