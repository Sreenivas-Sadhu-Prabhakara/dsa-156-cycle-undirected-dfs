package main

import "testing"

func TestHasCycle(t *testing.T) {
	cases := []struct {
		n     int
		edges [][2]int
		want  bool
	}{
		{4, [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 1}}, true},
		{4, [][2]int{{0, 1}, {1, 2}, {2, 3}}, false},
		{3, [][2]int{{0, 1}, {1, 2}, {2, 0}}, true},
		{6, [][2]int{{0, 1}, {2, 3}, {4, 5}}, false}, // forest
		{5, [][2]int{{0, 1}, {2, 3}, {3, 4}, {4, 2}}, true},
		{1, nil, false},
	}
	for _, c := range cases {
		if got := HasCycle(c.n, c.edges); got != c.want {
			t.Errorf("HasCycle(n=%d, %v) = %v, want %v", c.n, c.edges, got, c.want)
		}
	}
}
