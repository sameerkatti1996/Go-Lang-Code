/*
Set of linked lists operations:
1. Add at the end
2. Add at the start
3. Delete at the end
4. Delete at the start
*/
package main

import "fmt"

type node struct {
	value int
	next  *node
}

func addAtLast(head *node, value int) *node {
	//If head in nil, create a node that'd be head and return
	if head == nil {
		return &node{
			value: value,
			next:  nil,
		}
	}

	//Traverse through the LL upto last node
	temp := head
	for temp.next != nil {
		temp = temp.next
	}

	//Create a new node that'd be next of last node
	temp.next = &node{
		value: value,
		next:  nil,
	}

	return head
}

func addAtFront(head *node, value int) *node {
	//If head in nil, create a node that'd be head and return
	if head == nil {
		return &node{
			value: value,
			next:  nil,
		}
	}

	//Create a node whose next is head
	temp := &node{
		value: value,
		next:  head,
	}

	//Make this new node as head
	head = temp

	return head
}

func deleteAtLast(head *node) *node {
	//If LL is empty or has only one node, return nil
	if head == nil || head.next == nil {
		return nil
	}

	//Traverse till that last node with prev pointing to penultimate node
	temp := head
	var prev *node
	for temp.next != nil {
		prev = temp
		temp = temp.next
	}

	//next of penultimate node is made nil so that last node is deleted
	prev.next = nil

	return head
}

func deleteAtFront(head *node) *node {
	//If LL is empty or has only one node, return nil
	if head == nil || head.next == nil {
		return nil
	}
	//First node has to be deleted. Make head point to next node.
	return head.next
}

func printLL(head *node) {
	temp := head
	for temp != nil {
		fmt.Printf("%v ->", temp.value)
		temp = temp.next
	}
	fmt.Printf(" nil\n")
}

func main() {
	var head *node
	head = addAtLast(head, 10)
	printLL(head)
	head = addAtFront(head, 20)
	printLL(head)
	head = addAtLast(head, 30)
	printLL(head)
	head = addAtFront(head, 40)
	printLL(head)
	head = deleteAtFront(head)
	printLL(head)
	head = deleteAtLast(head)
	printLL(head)
	head = deleteAtFront(head)
	printLL(head)
}

/*
Output for this code:
10 -> nil                           Add 10 at end
20 ->10 -> nil                      Add 20 at start
20 ->10 ->30 -> nil                 Add 30 at end
40 ->20 ->10 ->30 -> nil            Add 40 at start
20 ->10 ->30 -> nil                 Delete at front
20 ->10 -> nil                      Delete at end
10 -> nil                           Delete at front
*/
