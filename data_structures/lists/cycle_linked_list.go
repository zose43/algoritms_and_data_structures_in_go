package lists

type cycledLinkedNode[T any] struct {
	value T
	prev  *cycledLinkedNode[T]
	next  *cycledLinkedNode[T]
}

func newCycledLinkedNode[T any](value T) *cycledLinkedNode[T] {
	return &cycledLinkedNode[T]{value: value}
}

type CycleLinkedList[T any] struct {
	head *cycledLinkedNode[T]
	size int
}

func NewCycleLinkedList[T any]() *CycleLinkedList[T] {
	return &CycleLinkedList[T]{}
}

func (cl *CycleLinkedList[T]) isEmpty() bool {
	return cl.size < 1
}

func (cl *CycleLinkedList[T]) Size() int {
	return cl.size
}

func (cl *CycleLinkedList[T]) Add(element T) {
	node := newCycledLinkedNode(element)
	if cl.isEmpty() {
		cl.head = node
		cl.head.prev = node
		cl.head.next = node
		cl.size++
		return
	}
	node.next = cl.head
	node.prev = cl.head.prev
	cl.head.prev.next = node
	cl.head.prev = node
	cl.head = node
	cl.size++
}
