package main

import "fmt"

//structure of heap
type heap struct {
	arr []int
}

//inserting into the heap

func (h *heap) insert(val int) {
	h.arr = append(h.arr, val)
	h.heapifyBtoT()
}

//rearranging the elements from bottom to top to satisfy the heap property

func (h *heap) heapifyBtoT() {
	index := len(h.arr) - 1
	for h.arr[parent(index)] < h.arr[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}
//deleting the root node

func (h *heap) extract() {
	element := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	fmt.Println(element)
	h.heapifyTtoB()
}

//rearranging the elements from top to bottom to satisfy the heap property

func (h *heap) heapifyTtoB() {
	index := 0
	for lchild(index) <= len(h.arr)-1 {
		lc := lchild(index)
		rc := rchild(index)
		if rc > len(h.arr)-1 && h.arr[lc] > h.arr[index] {
			h.swap(lc, index)
			index = lchild(index)
		} else if h.arr[lc] > h.arr[rc] && h.arr[index] < h.arr[lc] {
			h.swap(lc, index)
			index = lchild(index)
		} else if h.arr[rc] > h.arr[lc] && h.arr[index] < h.arr[rc] {
			h.swap(rc, index)
			index = rchild(index)
		} else {
			return
		}
	}
}

func (h *heap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

//to find the parent node of a node

func parent(i int) int {
	return (i - 1) / 2
}

//to find the left child of a node

func lchild(i int) int {
	return (i * 2) + 1
}

//to find the right child of a node

func rchild(i int) int {
	return (i * 2) + 2
}

func main() {
	var h heap
	arr := []int{3, 2, 1}
	for _, v := range arr {
		h.insert(v)
	}
	fmt.Println(h.arr)
	h.extract()
	fmt.Println(h.arr)
}
