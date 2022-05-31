package queue

import "strconv"

func openLock(deadends []string, target string) int {

	res := -1

	deMap := make(map[string]struct{}, len(deadends))
	for _, de := range deadends {
		deMap[de] = struct{}{}
	}

	q := &Queue{}
	visited := make(map[string]struct{}, 0)

	q.Enqueue("0000")
	visited["0000"] = struct{}{}

outer:
	for true {

		qSize := q.Len()
		res++

	inner:
		for i := 0; i < qSize; i++ {

			l := q.Dequeue()

			if target == l {
				break outer
			}

			if _, isDeadEnd := deMap[l]; isDeadEnd {
				continue inner
			}

			for i := 0; i < len(l); i++ {

				iVal := int(l[i] - '0')

				p := l[:i] + strconv.Itoa((iVal+1)%10) + l[i+1:]
				if _, found := visited[p]; !found {
					q.Enqueue(p)
					visited[p] = struct{}{}
				}

				n := l[:i] + strconv.Itoa((iVal+9)%10) + l[i+1:]
				if _, found := visited[n]; !found {
					q.Enqueue(n)
					visited[n] = struct{}{}
				}
			}
		}
	}

	return res
}

type Queue []string

func (q *Queue) Enqueue(val string) {
	*q = append(*q, val)
}

func (q *Queue) Dequeue() (val string) {
	val = (*q)[0]
	copy(*q, (*q)[1:])
	*q = (*q)[:len(*q)-1]

	return
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Len() int {
	return len(*q)
}
