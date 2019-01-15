package treenode

type GraphNode struct {
	Order    int
	Value    int
	Children []int
}

type GraphNodes []GraphNode

func (s GraphNodes) Len() int {
	return len(s)
}

func (s GraphNodes) Less(i, j int) bool {
	return s[i].Order > s[j].Order
}

func (s GraphNodes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
