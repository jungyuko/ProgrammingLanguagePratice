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

func (t *Tree) Traverse() string{
	if t == nil{
		return "()"
	}
	s := ""
	if t.Left != nil{
		s += t.Left.Traverse() +""
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil{
		s += ""+ t.Right.Traverse()
	}
	return "(" + s + ")"
}

func main() {

	newTree:=NewTree(1)
	fmt.Println(newTree.Traverse())
}
