//not exactly correct I dont think but good enough

package main

import "fmt"

type BSTNode struct {

	left *BSTNode 
	right *BSTNode
	value int
}

func (n *BSTNode) insert(v int){
	if n.value == 0{
		n.value = v 
		return
	}
	if n.value == v{
		return 
	}


	if n.value > v {
		if n.left == nil{
			n.left = &BSTNode{value: v}
			return

		}
		n.left.insert(v)
	}
	if n.value < v {
		if n.right == nil{
			n.right = &BSTNode{value: v}
			return
		}
		n.right.insert(v)
		
	}


}

 
func main(){
	
	node := &BSTNode{value:1}
	
	vals := []int{2,3,4,5,6,87,8,4,6,7,4}

	for _,val := range vals{
		node.insert(val)
	}
		
	fmt.Println(node.left.left.value)


}
