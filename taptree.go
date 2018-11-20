package main

import(	
	"fmt"
	// "unsafe"
	"math/rand"
	// "strconv"
)

var N_nodes int= 0

type Node struct{
	data int
	nchild int
	down *Node
	right *Node
}

func put_tree(node *Node, depth int){
	var p *Node

	if node == nil{
		return
	}
	for j:= 0; j< depth; j++{
		fmt.Print("		")
	}
	fmt.Println(node.data,"(",node.nchild,")")

	if node.nchild > 0{
		p = node.down
		for i:= node.nchild-1; i>=0; i--{
			put_tree(p, depth+1)
			if p == nil{return}
			p = p.right
		}
	}
}

func del_tree(node *Node){
	if node == nil{return}

	if(node.down != nil){
		del_tree(node.down)
	}
	if(node.right != nil){
		del_tree(node.right)
	}

	N_nodes = N_nodes-1
}

func get_newnode() *Node{
	var node *Node = new(Node)

	node.nchild = rand.Intn(4)+1
	node.down = nil
	node.right = nil

	N_nodes = N_nodes+1
	return node
}

func build_tree(n int) *Node{

	var node *Node = get_newnode()
	node.data = n * 1000

	if n-1 > 0{
		node.down = build_tree(n-1)
		node.down.data = node.down.data +1

		n2 := node.down
		for i:=0; i < node.nchild-1; i++{
			n2.right = build_tree(n-1)
			n2.right.data = (n2.right.data + (i+2))

			n2 = n2.right
		}
	}else{
		node.nchild = 0
	}
	return node
}

func main(){
	depth := 4
	var root *Node

	// if argc >= 2 && argv[1][0] == "-"{
	// 	depth, err := strconv.Atoi(argv[1]+1)
	// }

	root = build_tree(depth)
	fmt.Println("Total", N_nodes," nodes are created!")

	put_tree(root,0)
	// fmt.Println("throu put_tree")
	del_tree(root)
	// fmt.Println("throu del_tree")
	if(N_nodes != 0){
		fmt.Println("Total %d nodes are remaining!\n", N_nodes)
	}
	// fmt.Println("throu for")
}
