package main

import "fmt"

// HasCycle reports whether an undirected graph contains a cycle, via DFS
// carrying the parent: reaching a visited vertex by any edge other than the
// one we came in on closes a cycle.
func HasCycle(n int, edges [][2]int) bool {
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	visited := make([]bool, n)
	var walk func(v, parent int) bool
	walk = func(v, parent int) bool {
		visited[v] = true
		for _, w := range adj[v] {
			if !visited[w] {
				if walk(w, v) {
					return true
				}
			} else if w != parent {
				return true
			}
		}
		return false
	}
	for s := 0; s < n; s++ {
		if !visited[s] && walk(s, -1) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(HasCycle(4, [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 1}}))
	fmt.Println(HasCycle(4, [][2]int{{0, 1}, {1, 2}, {2, 3}}))
}
