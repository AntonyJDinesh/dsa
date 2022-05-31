package graph

func validPath(n int, edges [][]int, source int, destination int) bool {

	graph := NewGraph(n)
	graph.AddEdges(edges)

	stack := NewStack()
	stack.Push(source)
	visited := make([]bool, n)
	var pathFound bool

	for stack.Size() > 0 {
		nodeId := stack.Pop()
		if visited[nodeId] {
			continue
		}

		visited[nodeId] = true

		if nodeId == destination {
			pathFound = true
			break
		}

		for connectedNodeId, _ := range graph.Nodes[nodeId].Edges {
			stack.Push(connectedNodeId)
		}
	}

	return pathFound

}
