/*
Doubly linked lists operations:
1. Add at the end
2. Add at the start
3. Delete at the end
4. Delete at the start
*/
package main

import "fmt"

type node struct {
	value int
	left  *node
	right *node
}

func addAtLast(head *node, value int) *node {
	//If head is nil, create a node that'd be head
	if head == nil {
		return &node{
			value: value,
			left:  nil,
			right: nil,
		}
	}

	//Traverse till last node
	temp := head
	for temp.right != nil {
		temp = temp.right
	}

	//At last node, create a new node whose left will be current last(temp)
	//and current last's(temp's) right will be new node
	temp.right = &node{
		value: value,
		left:  temp,
		right: nil,
	}

	return head
}

func addAtFront(head *node, value int) *node {
	//If head is nil, create a node that'd be head
	if head == nil {
		return &node{
			value: value,
			left:  nil,
			right: nil,
		}
	}

	//Create a new node whose right will be head and change head's
	//left to point to this node.
	//This new node will become new head
	temp := &node{
		value: value,
		left:  nil,
		right: head,
	}
	head.left = temp
	head = temp

	return head
}

func deleteAtLast(head *node) *node {
	//If the DLL is empty or has only one node, return nil
	if head == nil || head.right == nil {
		return nil
	}

	//Traverse till last node
	temp := head
	for temp.right != nil {
		temp = temp.right
	}

	//Make right of penultimate node nil so that last node is deleted
	temp.left.right = nil

	return head
}

func deleteAtFront(head *node) *node {
	//If the DLL is empty or has only one node, return nil
	if head == nil || head.right == nil {
		return nil
	}

	//Make left of second node nil
	head.right.left = nil

	//Make second node the new head
	head = head.right

	return head
}

func printDLL(head *node) {
	fmt.Printf("nil <- ")
	temp := head
	for temp.right != nil {
		fmt.Printf("%v <=> ", temp.value)
		temp = temp.right
	}
	fmt.Printf("%v -> nil\n", temp.value)

}

func main() {
	var head *node
	head = addAtLast(head, 10)
	printDLL(head)
	head = addAtFront(head, 20)
	printDLL(head)
	head = addAtLast(head, 30)
	printDLL(head)
	head = addAtFront(head, 40)
	printDLL(head)
	head = deleteAtFront(head)
	printDLL(head)
	head = deleteAtLast(head)
	printDLL(head)
	head = deleteAtFront(head)
	printDLL(head)
}

/*
Output:
nil <- 10 -> nil                            Add 10 at last
nil <- 20 <=> 10 -> nil                     Add 20 at front
nil <- 20 <=> 10 <=> 30 -> nil              Add 30 at last
nil <- 40 <=> 20 <=> 10 <=> 30 -> nil       Add 40 at front
nil <- 20 <=> 10 <=> 30 -> nil              Delete at front
nil <- 20 <=> 10 -> nil                     Delete at last
nil <- 10 -> nil                            Delete at front
*/
