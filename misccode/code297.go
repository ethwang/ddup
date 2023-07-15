package misccode

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

//func Constructor() Codec {
//	return Codec{}
//}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	res := ""
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			res += "nil,"
			return
		}
		res += fmt.Sprint(node.Val) + ","
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return res
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	datas := strings.Split(data, ",")
	var construct func(datas []string) *TreeNode
	construct = func(datas []string) *TreeNode {
		if datas[0] == "nil" {
			datas = datas[1:]
			return nil
		}
		val, _ := strconv.Atoi(datas[0])
		datas = datas[1:]
		return &TreeNode{Val: val, Left: construct(datas), Right: construct(datas)}
	}
	return construct(datas)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
