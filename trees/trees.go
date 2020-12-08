package trees

// Tree ...
type Tree interface {
	InsertTree(...int)
	SearchTree(int) bool
	DeleteTree(int) bool
}
