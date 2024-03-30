package main_test

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {

	a := []int{1, 2}
	pivot := int(len(a) / 2)
	fmt.Println(pivot)

	fmt.Println(a[:pivot])
	fmt.Println(a[pivot+1:])
}