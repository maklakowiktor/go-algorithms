package datastructures

type Node struct {
	data int
	next *Node
}

func NewNode(data int) *Node {
	return &Node{
		data: data,
		next: nil,
	}
}

type LinkedList struct {
	head *Node
}

func NewLinkedList(data int) *LinkedList {
	return &LinkedList{
		head: NewNode(data),
	}
}

func (list *LinkedList) InsertAtBack(data int) {
	newNode := NewNode(data)
	if list.head == nil {
		list.head = newNode
		return
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (list *LinkedList) InsertAtFront(data int) {
	if list.head == nil {
		newNode := NewNode(data)
		list.head = newNode
		return
	}

	newNode := &Node{data: data, next: list.head}
	list.head = newNode
}

func (list *LinkedList) InsertAfterValue(afterValue, data int) {
	newNode := NewNode(data)
	current := list.head
	for current != nil {
		if current.data == afterValue {
			newNode.next = current.next
			current.next = newNode
			return
		}
		current = current.next
	}
}
