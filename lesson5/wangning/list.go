package main

import "fmt"

type Student struct {
	Id   int
	Name string
}

type Node struct {
	Val  Student
	Next *Node
}

func main() {

	node_a := Node{Val: Student{Id: 1, Name: "student_a"}}
	node_b := Node{Val: Student{Id: 2, Name: "student_b"}}
	node_c := Node{Val: Student{Id: 3, Name: "student_c"}}
	node_d := Node{Val: Student{Id: 4, Name: "student_d"}}

	node_a.Next = &node_b
	node_b.Next = &node_c
	node_c.Next = &node_d


	p := &node_a

	for p != nil {
		node := *p
		fmt.Println(node)
		fmt.Println(node.Val.Id,node.Val.Name)
		p = node.Next
	}
//	fmt.Println(node_a)
	//	node_a.Next.Next = &node_c

	//	p:= ReverseList(&node_a)
	//	printLlist(p)

}
