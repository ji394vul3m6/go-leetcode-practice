package binarytreenode

import (
	"fmt"
	"sort"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) ToArray() []interface{} {
	if node == nil {
		return []interface{}{nil}
	}

	ret := []interface{}{}
	ret = append(ret, node.Val)
	ret = append(ret, node.Left.ToArray()...)
	ret = append(ret, node.Right.ToArray()...)
	return ret
}

func (node *TreeNode) ToString() string {
	ret := node.ToArray()
	return fmt.Sprintf("%+v", ret)
}

type TreeNodes []*TreeNode

func (list TreeNodes) ToString() string {
	strs := []string{}
	for _, node := range list {
		strs = append(strs, node.ToString())
	}
	sort.Strings(strs)
	return strings.Join(strs, ",")
}
