package main

import (
	"data_structure/tree"
	"encoding/json"
	"fmt"
)

func main() {

	//fibbonnaciNumbers := []int{0, 1, 2, 1, 8, 5, 3, 144, 21, 34, 55, 610, 13, 233, 377, 89, 987, 89}
	//fibbonnaciNumbers := []int{0, 1, 2, 3, 4, 5}
	fibbonnaciNumbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	orderedAndNoDuplicatesSlice := tree.OrderSliceAndRemoveDuplicates(fibbonnaciNumbers)

	binaryTree := tree.NewBinaryTree(orderedAndNoDuplicatesSlice)

	//printBinaryTree(binaryTree)

	isFound := binaryTree.Search(11)

	if isFound {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found")
	}
}

func PrintBinaryTree(binaryTree *tree.BinaryTree) {
	data, err := json.Marshal(binaryTree)
	if err != nil {
		fmt.Println("Error converting binary tree into json")
	}

	fmt.Println(string(data))
}
