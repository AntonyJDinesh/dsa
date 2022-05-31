package graph

import "testing"

func Test_validPath(t *testing.T) {
	testData := []struct {
		n, source, destination int
		edges                  [][]int
		found                  bool
	}{
		{3, 0, 2, [][]int{{0, 1}, {1, 2}, {2, 0}}, true},
		{6, 0, 5, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, false},
	}

	for _, td := range testData {
		res := validPath(td.n, td.edges, td.source, td.destination)

		if res != td.found {
			t.Errorf("td.n: %d, td.edges: %#v, td.source: %d, td.destination: %d, Exp: %t, Got: %t", td.n, td.edges, td.source, td.destination, td.found, res)
		}
	}
}
