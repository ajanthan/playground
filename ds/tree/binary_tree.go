package tree

type BTNode struct {
	Data  interface{}
	Found bool
	Size  int
	Left  *BTNode
	Right *BTNode
}
type BBTNode struct {
	Data   interface{}
	Left   *BBTNode
	Right  *BBTNode
	Parent *BBTNode
}
