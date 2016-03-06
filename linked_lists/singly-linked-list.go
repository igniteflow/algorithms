package main

import "fmt"

type Node struct {
    data int
    next *Node
}

type SingleList struct {
  head *Node
  tail *Node
}

func (s SingleList) insertAfter(node, newNode *Node) {
  newNode.next = node.next
  node.next = newNode
}

func PrintList(node *Node) {
    if node == nil {
        return
    }
    fmt.Println(node.data)
    PrintList(node.next)
}

func main() {
    node3 := &Node{data: 3}
    node2 := &Node{data: 2}
    node1 := &Node{data: 1}

    list := SingleList{head: node1}
    list.insertAfter(node1, node2)
    list.insertAfter(node2, node3)

    PrintList(node1)
}
