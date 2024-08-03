package main

import (
	"bufio"
	"fmt"
	"os"
)

type TreeNode struct {
	Value    int
	Children []*TreeNode
	Parent   *TreeNode
}

func findAllElements(treeCode []int) map[int]bool {
	parent := make(map[int]bool)

	for i := 0; i < len(treeCode); {

		val := treeCode[i]

		//fmt.Printf("Родитель: %d\n", val)

		if _, ok := parent[val]; !ok {
			parent[val] = false
		}

		count := treeCode[i+1]

		//fmt.Printf("Количество детей: %d\n", count)

		children := treeCode[i+2 : i+2+count]

		//fmt.Printf("Дети: %v\n\n", children)

		for _, child := range children {
			parent[child] = true
		}

		i += count + 2
	}

	return parent
}

func main() {
	out := bufio.NewWriter(os.Stdout)

	result := findAllElements([]int{3, 0, 1, 0, 5, 2, 2, 6, 4, 3, 5, 1, 3, 2, 0, 6, 0})

	for key, v := range result {
		if !v {
			fmt.Fprintf(out, "%d", key)
			out.Flush()
		} else {
			continue
		}
	}

}
