package find_and_union

import "fmt"

type DisjointSet interface {
	Find(vertex int) (root int)
	Union(vertex1, vertex2 int)
	Connected(vertex1, vertex2 int) bool
	GetRootList() []int
}

type QuickFindDisjointSet struct {
	rootList []int
}

func NewQuickFindDisjointSet(size int) DisjointSet {
	qfds := &QuickFindDisjointSet{rootList: make([]int, size)}
	for i := 0; i < len(qfds.rootList); i++ {
		qfds.rootList[i] = i
	}
	return qfds
}

func (qfds QuickFindDisjointSet) Find(vertex int) (root int) {
	return qfds.rootList[vertex]
}

func (qfds QuickFindDisjointSet) Union(vertex1, vertex2 int) {
	rootV1 := qfds.Find(vertex1)
	rootV2 := qfds.Find(vertex2)

	if rootV1 != rootV2 {
		for i := 0; i < len(qfds.rootList); i++ {
			if qfds.rootList[i] == rootV2 {
				qfds.rootList[i] = rootV1
			}
		}
	}
}

func (qfds QuickFindDisjointSet) Connected(vertex1, vertex2 int) bool {
	return qfds.Find(vertex1) == qfds.Find(vertex2)
}

func (qfds QuickFindDisjointSet) GetRootList() []int {
	return qfds.rootList
}

func findCircleNum(isConnected [][]int) int {
	qfds := NewQuickFindDisjointSet(len(isConnected))

	for cityA, connectedData := range isConnected {
		for cityB, connected := range connectedData {
			if connected == 1 {
				qfds.Union(cityA, cityB)
				fmt.Printf("cityA: %d, cityB: %d, root: %#v\n", cityA, cityB, qfds.GetRootList())
			}
		}
	}

	fmt.Printf("root: %#v", qfds.GetRootList())
	uniqueRoots := make(map[int]struct{}, 0)
	var count int
	for _, root := range qfds.GetRootList() {
		if _, found := uniqueRoots[root]; !found {
			uniqueRoots[root] = struct{}{}
			count++
		}
	}

	return count
}
