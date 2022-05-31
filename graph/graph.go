package graph

type Graph struct {
	size  int
	Nodes []*GraphNode
}

type GraphNode struct {
	Id    int
	Edges map[int]struct{}
}

func NewGraph(size int) *Graph {
	graph := &Graph{size: size, Nodes: make([]*GraphNode, size)}
	for i := 0; i < size; i++ {
		graph.Nodes[i] = &GraphNode{Id: i, Edges: make(map[int]struct{}, 0)}
	}

	return graph
}

func (graph *Graph) AddEdges(edges [][]int) {
	for _, edge := range edges {
		graph.Nodes[edge[0]].Edges[edge[1]] = struct{}{}
		graph.Nodes[edge[1]].Edges[edge[0]] = struct{}{}
	}
}

type Stack interface {
	Push(ele int)
	Pop() (ele int)
	Size() (size int)
}

func NewStack() (stack Stack) {
	return &StackImpl{}
}

type StackImpl struct {
	container []int
}

func (stackImpl *StackImpl) Push(ele int) {
	stackImpl.container = append(stackImpl.container, ele)
}

func (stackImpl *StackImpl) Pop() (ele int) {
	ele = stackImpl.container[len(stackImpl.container)-1]
	stackImpl.container = stackImpl.container[:len(stackImpl.container)-1]
	return ele
}

func (stackImpl *StackImpl) Size() (size int) {
	return len(stackImpl.container)
}
