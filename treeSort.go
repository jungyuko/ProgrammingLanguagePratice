package main

import(
	"fmt"
	"math/rand"
)

type Tree struct{
	Left *Tree
	Value int
	Right *Tree
}

func NewTree(k int) *Tree{
	var t *Tree
	for _, v := range rand.Perm(10){
		t = insert(t, (1+v) * k)
	}
	return t
}

func insert(t *Tree, v int) *Tree{
	if t == nil{
		return &Tree{nil, v, nil}
	}
	if v < t.Value{
		t.Left = insert(t.Left, v)
	}else{
		t.Right = insert(t.Right, v)
	}
	return t
}

func Traverse(t *Tree){
	if t == nil{
		return
	}
	Traverse(t.Left)
	a := t.Value
	fmt.Print(a," ")
	Traverse(t.Right)
}

func main() {

	newTree:=NewTree(1)
	Traverse(newTree)
	fmt.Println()
}
