package trees

import (
	"sync"
)

const (
	rootNode  = 0
	leftNode  = -1
	rightNode = 1
)

// nodeBst ...
type nodeBst struct {
	value int
	left  *nodeBst
	right *nodeBst
}

// TreeBst ...
type TreeBst struct {
	head *nodeBst
	Mux  *sync.RWMutex
}

// InsertTree ...
func (t *TreeBst) InsertTree(numbers ...int) {
	for _, number := range numbers {
		if t.head == nil {
			t.head = &nodeBst{
				value: number,
				left:  nil,
				right: nil,
			}
			continue
		}
		insertNodeBst(t, number)
	}
}

func insertNodeBst(t *TreeBst, number int) {
	t.Mux.Lock()
	insertNodeBstRec(t.head, number)
	t.Mux.Unlock()
}

func insertNodeBstRec(n *nodeBst, number int) {
	switch {
	case number == n.value:
		return
	case number < n.value && n.left != nil:
		insertNodeBstRec(n.left, number)
	case number < n.value && n.left == nil:
		n.left = &nodeBst{
			value: number,
			left:  nil,
			right: nil,
		}
		return
	case number > n.value && n.right != nil:
		insertNodeBstRec(n.right, number)
	case number > n.value && n.right == nil:
		n.right = &nodeBst{
			value: number,
			left:  nil,
			right: nil,
		}
		return
	}
}

// SearchTree ...
func (t *TreeBst) SearchTree(number int) bool {
	n := searchNodeBst(t, number)
	if n == nil {
		return false
	}
	return true
}

func searchNodeBst(t *TreeBst, number int) *nodeBst {
	t.Mux.RLock()
	n := searchNodeBstRec(t.head, number)
	t.Mux.RUnlock()
	return n
}

func searchNodeBstRec(n *nodeBst, number int) *nodeBst {
	switch {
	case n == nil:
		return nil
	case number == n.value:
		return n
	case number < n.value && n.left != nil:
		return searchNodeBstRec(n.left, number)
	case number > n.value && n.right != nil:
		return searchNodeBstRec(n.right, number)
	default:
		return nil
	}
}

// DeleteTree ...
func (t *TreeBst) DeleteTree(number int) bool {
	if t.head == nil {
		return false
	}
	t.Mux.Lock()
	res := deleteNodeBst(t, number)
	t.Mux.Unlock()
	return res
}

func deleteNodeBst(t *TreeBst, number int) bool {
	searchNode, previousNode, dir := searchDeleteNodeBst(t.head, t.head, 0, number)
	switch {
	case searchNode == nil:
		return false
	case searchNode.left != nil && searchNode.right == nil:
		writeSubNodeBst(searchNode.left, previousNode, dir)
		return true
	case searchNode.left != nil && searchNode.right != nil:
		subNode := searchNode.right
		if subNode.left == nil {
			subNode.left = searchNode.left
			writeSubNodeBst(subNode, previousNode, dir)
			return true
		}
		subNode = searchMaxLeftSubNodeBst(subNode, &subNode)
		subNode.left = searchNode.left
		subNode.right = searchNode.right
		writeSubNodeBst(subNode, previousNode, dir)
		return true
	default:
		return false
	}
}

func writeSubNodeBst(n *nodeBst, pre *nodeBst, dir int) {
	switch dir {
	case rootNode:
		pre = n
	case leftNode:
		pre.left = n
	case rightNode:
		pre.right = n
	}
}

func searchDeleteNodeBst(n *nodeBst, pre *nodeBst, dir int, number int) (*nodeBst, *nodeBst, int) {
	switch {
	case n == nil:
		return nil, pre, dir
	case number == n.value:
		return n, pre, dir
	case number < n.value && n.left != nil:
		return searchDeleteNodeBst(n.left, n, leftNode, number)
	case number > n.value && n.right != nil:
		return searchDeleteNodeBst(n.right, n, rightNode, number)
	default:
		return nil, nil, dir
	}
}

func searchMaxLeftSubNodeBst(n *nodeBst, pre **nodeBst) *nodeBst {
	if n != nil && n.left != nil {
		return searchMaxLeftSubNodeBst(n.left, &n.left)
	}
	*pre = n.right
	n.right = nil
	return n
}
