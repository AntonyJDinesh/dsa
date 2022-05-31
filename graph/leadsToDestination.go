package graph

func leadsToDestination(n int, edges [][]int, source int, destination int) bool {

	if source == destination {
		return true
	}

	nodes := make([][]int, n)
	for i := 0; i < n; i++ {
		nodes[i] = make([]int, 0)
	}

	for _, edge := range edges {
		nodes[edge[0]] = append(nodes[edge[0]], edge[1])
	}

	if len(nodes[source]) == 0 {
		return false
	}

	for _, neighbour := range nodes[source] {

		nList := make([]int, 0)
		nList = append(nList, neighbour)
		visited := make(map[int]struct{}, 0)

		found := false
	inner:
		for len(nList) > 0 {
			node := nList[len(nList)-1]
			nList = nList[0 : len(nList)-1]

			if _, found := visited[node]; found {
				continue inner
			}

			if node == destination {
				found = true
				break inner
			}

			nList = append(nList, nodes[node]...)
			visited[node] = struct{}{}

		}

		if !found {
			return false
		}
	}

	return true
}
