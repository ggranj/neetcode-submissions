type Node struct {
	val int
	next *Node
}

type LinkedList struct {
	head *Node
	len int // сколько всего списков с каким-то val внутри
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: nil,
		len: 0,
	}
}

func (ll *LinkedList) Get(index int) int {
	if index >= ll.len {
		return -1
	}

	if index == 0 {
		return ll.head.val
	}

	cur := ll.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}

	return cur.val
}

func (ll *LinkedList) InsertHead(val int) {
	newHead := &Node{
		val: val,
		next: ll.head,
	}

	ll.head = newHead
	ll.len = ll.len + 1
	return
}

func (ll *LinkedList) InsertTail(val int) {
	if ll.head == nil {
		ll.head = &Node{
			val: val,
			next: nil,
		}

		ll.len = ll.len + 1
		return
	}

	newTail := &Node{
		val: val,
		next: nil,
	}

	cur := ll.head
	for cur.next != nil {
		cur = cur.next
	}

	cur.next = newTail
	ll.len = ll.len + 1
	
	return
}

func (ll *LinkedList) Remove(index int) bool {
	if index >= ll.len { // len с 1 стартует, а index с 0, так что ">="
		return false
	}

	if index == 0 {
		ll.head = ll.head.next
		ll.len = ll.len - 1
		return true
	}

	cur := ll.head // стоим на [0] - первый элемент в списке
	for i := 0; i < index-1; i++ {
		cur = cur.next
	}

	cur.next = cur.next.next
	ll.len = ll.len - 1

	return true
}

func (ll *LinkedList) GetValues() []int {
	total := make([]int, 0, ll.len)

	for cur := ll.head; cur != nil; cur = cur.next {
		total = append(total, cur.val)
	}

	return total
}
