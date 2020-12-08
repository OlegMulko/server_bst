package trees

import (
	"sync"
	"testing"
)

type testDataInsDel struct {
	inValue  []int
	outValue []int
	level    []int
}

type testSetSearch struct {
	searchValue int
	result      bool
}

type testDataSearch struct {
	values []int
	sets   []testSetSearch
}

type testParam struct {
	tD *testDataInsDel
	t  *testing.T
}

type testDataDelete struct {
	tDID        testDataInsDel
	deleteValue int
}

var tDI []testDataInsDel = []testDataInsDel{
	testDataInsDel{
		inValue:  []int{8, 5, 10, 2, 7, 6},
		outValue: []int{8, 5, 2, 7, 6, 10},
		level:    []int{0, 1, 2, 2, 3, 1},
	},
}

var tDS testDataSearch = testDataSearch{
	values: []int{8, 5, 10, 2, 7, 6},
	sets: []testSetSearch{
		testSetSearch{
			searchValue: 5,
			result:      true,
		},
		testSetSearch{
			searchValue: 9,
			result:      false,
		},
	},
}

var tDD []testDataDelete = []testDataDelete{
	testDataDelete{
		tDID: testDataInsDel{
			inValue:  []int{8, 5, 10, 2, 7, 6},
			outValue: []int{8, 6, 2, 7, 10},
			level:    []int{0, 1, 2, 2, 1},
		},
		deleteValue: 5,
	},
}

func TestInsertTree(t *testing.T) {
	for _, test := range tDI {
		tree := &TreeBst{
			Mux: &sync.RWMutex{},
		}
		tp := &testParam{
			tD: &test,
			t:  t,
		}
		nodeID := 0
		tree.InsertTree(test.inValue...)
		testLeftTravelTree(tree.head, &nodeID, 0, tp)
	}
}

func testLeftTravelTree(n *nodeBst, nodeID *int, level int, tp *testParam) {
	if n.value != tp.tD.outValue[*nodeID] ||
		level != tp.tD.level[*nodeID] {
		tp.t.Fatalf("Expected value %v, current value %v, expected level %v, current level %v\n",
			tp.tD.outValue[*nodeID],
			n.value,
			level,
			tp.tD.level[*nodeID])
		return
	}
	if n.left != nil {
		(*nodeID)++
		testLeftTravelTree(n.left, nodeID, level+1, tp)
	}
	if n.right != nil {
		(*nodeID)++
		testLeftTravelTree(n.right, nodeID, level+1, tp)
	}
}

func TestSearchTree(t *testing.T) {
	tree := &TreeBst{
		Mux: &sync.RWMutex{},
	}
	tree.InsertTree(tDS.values...)
	for _, test := range tDS.sets {
		res := tree.SearchTree(test.searchValue)
		if res != test.result {
			t.Fatalf("Value %v, expected result %v, current result %v\n",
				test.searchValue,
				test.result,
				res)
			continue
		}
	}
}

func TestDeleteTree(t *testing.T) {
	for _, test := range tDD {
		tree := &TreeBst{
			Mux: &sync.RWMutex{},
		}
		tp := &testParam{
			tD: &test.tDID,
			t:  t,
		}
		nodeID := 0
		tree.InsertTree(test.tDID.inValue...)
		tree.DeleteTree(test.deleteValue)
		testLeftTravelTree(tree.head, &nodeID, 0, tp)
	}
}
