package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	input := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
		},
	}
	codec := Constructor()
	serializedData := codec.serialize(input)
	deserializedData := codec.deserialize(serializedData)
	fmt.Println(serializedData)
	fmt.Println(deserializedData)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// serialize a tree to a single string
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "nil,"
	}
	str := strconv.Itoa(root.Val) + ","
	str += this.serialize(root.Left)
	str += this.serialize(root.Right)

	return str
}

// deserialize your encoded data to tree
func (this *Codec) deserialize(data string) *TreeNode {
	datalist := strings.Split(data, ",")
	var recursiveFn func(*[]string) *TreeNode
	recursiveFn = func(l *[]string) *TreeNode {
		if (*l)[0] == "nil" {
			*l = (*l)[1:]
			return nil
		}

		v, err := strconv.Atoi((*l)[0])
		if err != nil {
			log.Fatal(err)
		}
		root := &TreeNode{Val: v}
		*l = (*l)[1:]
		root.Left = recursiveFn(l)
		root.Right = recursiveFn(l)

		return root
	}
	return recursiveFn(&datalist)
}
