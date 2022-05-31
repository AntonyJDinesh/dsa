package find_and_union

type quickFindImpl struct {
	rootList []int
}

func newQuickFindImpl(size int) *quickFindImpl {
	qfi := &quickFindImpl{rootList: make([]int, size)}
	for i := 0; i < size; i++ {
		qfi.rootList[i] = i
	}

	return qfi
}

func (qfi *quickFindImpl) Find(vertex int) (root int) {
	return qfi.rootList[vertex]
}

func (qfi *quickFindImpl) Union(vertex1, vertex2 int) {

	if !qfi.Connected(vertex1, vertex2) {
		newRoot := qfi.Find(vertex1)
		oldRoot := qfi.Find(vertex2)

		for i := 0; i < len(qfi.rootList); i++ {
			if qfi.rootList[i] == oldRoot {
				qfi.rootList[i] = newRoot
			}
		}
	}
}

func (qfi *quickFindImpl) Connected(vertex1, vertex2 int) (isConnected bool) {
	return qfi.rootList[vertex1] == qfi.rootList[vertex2]
}
