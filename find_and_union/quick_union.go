package find_and_union

type quickUnionImpl struct {
	rootList []int
}

func newQuickUnionImpl(size int) *quickUnionImpl {
	qui := &quickUnionImpl{rootList: make([]int, size)}
	for i := 0; i < size; i++ {
		qui.rootList[i] = i
	}

	return qui
}

func (qui *quickUnionImpl) Find(vertex int) (root int) {
	root = vertex
	for root != qui.rootList[root] {
		root = qui.rootList[root]
	}
	return
}

func (qui *quickUnionImpl) Union(vertex1, vertex2 int) {
	rootV1 := qui.Find(vertex1)
	rootV2 := qui.Find(vertex2)
	if rootV1 != rootV2 {
		qui.rootList[rootV1] = rootV2
	}
}

func (qui *quickUnionImpl) Connected(vertex1, vertex2 int) (isConnected bool) {
	return qui.Find(vertex1) == qui.Find(vertex2)
}
