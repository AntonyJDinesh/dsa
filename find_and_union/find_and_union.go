package find_and_union

type FindAndUnion interface {
	Find(vertex int) (root int)
	Union(vertex1, vertex2 int)
	Connected(vertex1, vertex2 int) (isConnected bool)
}

type FindAndUnionType int

const (
	QuickFind  FindAndUnionType = 1
	QuickUnion FindAndUnionType = 2
)

func NewFindAndUnion(size int, typ FindAndUnionType) (fau FindAndUnion) {
	if typ == QuickFind {
		fau = newQuickFindImpl(size)
	}

	return
}
