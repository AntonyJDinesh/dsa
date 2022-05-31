package tree

import "testing"

func Test_maxDepth(t *testing.T) {
	three := &TreeNode{Val: 3}
	nine := &TreeNode{Val: 9}
	twenty := &TreeNode{Val: 20}
	fifteen := &TreeNode{Val: 15}
	seven := &TreeNode{Val: 7}

	three.Left = nine
	three.Right = twenty

	twenty.Left = fifteen
	twenty.Right = seven

	res := maxDepth(three)
	if res != 3 {
		t.Errorf("Expexted: %d, Got: %d", 3, res)
	}
}

func Test_isSameTree(t *testing.T) {
	three := &TreeNode{Val: 3}
	nine := &TreeNode{Val: 9}
	twenty := &TreeNode{Val: 20}
	fifteen := &TreeNode{Val: 15}
	seven := &TreeNode{Val: 7}

	three.Left = nine
	three.Right = twenty

	twenty.Left = fifteen
	twenty.Right = seven

	res := isSameTree(three, three)
	if !res {
		t.Errorf("Expexted: %t, Got: %t", true, res)
	}
}

func Test_isSubtree(t *testing.T) {
	three := &TreeNode{Val: 3}
	nine := &TreeNode{Val: 9}
	twenty := &TreeNode{Val: 20}
	fifteen := &TreeNode{Val: 15}
	seven := &TreeNode{Val: 7}

	three.Left = nine
	three.Right = twenty

	twenty.Left = fifteen
	twenty.Right = seven

	res := isSubtree(three, nine)
	if !res {
		t.Errorf("isSubtree failed")
	}

	ten := &TreeNode{Val: 10}
	res = isSubtree(three, ten)
	if res {
		t.Errorf("isSubtree failed")
	}
}

func Test_countNodes(t *testing.T) {
	one := &TreeNode{Val: 1}
	two := &TreeNode{Val: 2}
	three := &TreeNode{Val: 3}
	four := &TreeNode{Val: 4}
	five := &TreeNode{Val: 5}
	six := &TreeNode{Val: 6}
	seven := &TreeNode{Val: 7}
	eight := &TreeNode{Val: 8}
	nine := &TreeNode{Val: 9}
	ten := &TreeNode{Val: 10}

	one.Left = two
	one.Right = three

	two.Left = four
	two.Right = five

	three.Left = six
	three.Right = seven

	four.Left = eight
	four.Right = nine

	five.Left = ten

	res := countNodes(one)
	if res != 10 {
		t.Errorf("Expected: %d, Got: %d", 5, res)
	}
}

func Test_numOfMinutes(t *testing.T) {
	testData := []struct {
		n, headID, res      int
		manager, informTime []int
	}{
		{6, 2, 1, []int{2, 2, -1, 2, 2, 2}, []int{0, 0, 1, 0, 0, 0}},
	}

	for _, td := range testData {

		res := numOfMinutes(td.n, td.headID, td.manager, td.informTime)
		if res != td.res {
			t.Errorf("td.n: %d, td.headID: %d, td.manager: %#v, td.informTime: %#v, td.res: %d, Res: %d", td.n, td.headID, td.manager, td.informTime, td.res, res)
		}
	}
}
