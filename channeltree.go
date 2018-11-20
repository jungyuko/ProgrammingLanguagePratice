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

func NewTree(i int) *Tree{
	var t *Tree
	for _,v := range rand.Perm(10){
		t = insert(t, (1+v) * i)
	}
	return t
}

func insert(t *Tree, v int) *Tree{
	if t == nil{
		return &Tree{nil, v, nil}
	}
	if v < t.Value{
		t.Left = insert(t.Left, v)
		return t
	}else{
		t.Right = insert(t.Right, v)
		return t
	}
	return t
}

func channel_tree(a int, c chan int){
	prints := a
	c <- prints
}

func Traverse(t *Tree){
	if t == nil{
		return
	}

	Traverse(t.Left)
	var a int
	c := make(chan int)
	a = t.Value
	go channel_tree(a, c)
	x := <- c
	fmt.Print(" ",x)
	Traverse(t.Right)
}

func main(){

	newTree := NewTree(1)
	Traverse(newTree)
	fmt.Println()
}
