package treesort

import (
	"bytes"
	"fmt"
	"strconv"
)

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	var sorted = make([]int, 0)
	sorted = appendValues(sorted, t)
	var buffer bytes.Buffer
	for i, v := range sorted {
		if i > 0 {
			buffer.WriteString(" ")
		}
		buffer.WriteString(strconv.Itoa(v))
	}
	return buffer.String()
}

func PrintValues(values []int) {
	// make tree
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	fmt.Println(root)
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
