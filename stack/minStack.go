package stack

import "fmt"

type MinStack struct {
	stkRoot    *node
	minStkRoot *node
}

type node struct {
	val int

	stkNext *node
	stkPrev *node

	minStkNext *node
	minStkPrev *node
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {

	newNode := &node{val: val}

	if this.isEmpty() {
		this.stkRoot = newNode
		this.minStkRoot = newNode
	} else {
		// stack
		newNode.stkNext = this.stkRoot
		this.stkRoot.stkPrev = newNode
		this.stkRoot = newNode

		// min stack
		curNode := this.minStkRoot
		var prevNode *node
		for curNode != nil && curNode.val < val {
			prevNode = curNode
			curNode = curNode.minStkNext
		}

		if curNode == nil { // rear
			prevNode.minStkNext = newNode
			newNode.minStkPrev = prevNode
			fmt.Println("rear")
		} else if prevNode == nil { // front
			newNode.minStkNext = curNode
			curNode.minStkPrev = newNode
			this.minStkRoot = newNode
			fmt.Println("front")
		} else { // middle
			prevNode.minStkNext = newNode
			curNode.minStkPrev = newNode

			newNode.minStkPrev = prevNode
			newNode.minStkNext = curNode
			fmt.Println("middle")
		}
	}
}

func (this *MinStack) Pop() {
	if this.isEmpty() {
		panic("stack is empty")
	}

	nodeToBeRemoved := this.stkRoot

	this.stkRoot = nodeToBeRemoved.stkNext
	nodeToBeRemoved.stkNext = nil

	if this.stkRoot != nil {
		this.stkRoot.stkPrev = nil
	}

	// min stack
	if nodeToBeRemoved.minStkPrev == nil { // Front
		this.minStkRoot = nodeToBeRemoved.minStkNext
		nodeToBeRemoved.minStkNext = nil
		nodeToBeRemoved.minStkPrev = nil
		fmt.Println("\nFront")
	} else if nodeToBeRemoved.minStkNext == nil { // Rear
		nodeToBeRemoved.minStkPrev.minStkNext = nil
		nodeToBeRemoved.minStkPrev = nil
		nodeToBeRemoved.minStkNext = nil
		fmt.Println("\nRear")
	} else { // middle
		prev := nodeToBeRemoved.minStkPrev
		nxt := nodeToBeRemoved.minStkNext

		prev.minStkNext = nxt
		nxt.minStkPrev = prev

		nodeToBeRemoved.minStkPrev = nil
		nodeToBeRemoved.minStkNext = nil
		fmt.Printf("\nmiddle: %d", nodeToBeRemoved.val)
	}
}

func (this *MinStack) Top() int {
	if this.isEmpty() {
		panic("")
	}

	return this.stkRoot.val
}

func (this *MinStack) GetMin() int {
	if this.isEmpty() {
		panic("")
	}

	return this.minStkRoot.val
}

func (this *MinStack) isEmpty() bool {
	return this.stkRoot == nil
}

func (this *MinStack) print() {
	n := this.stkRoot
	fmt.Print("\nStk: ")
	for n != nil {
		fmt.Printf("%d->", n.val)
		n = n.stkNext
	}

	n = this.minStkRoot
	fmt.Print("\nMinStk: ")
	for n != nil {
		fmt.Printf("%d->", n.val)
		n = n.minStkNext
	}
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
