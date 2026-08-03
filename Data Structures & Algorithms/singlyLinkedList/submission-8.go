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

	dummy := &Node{next: ll.head}
	cur := dummy // стоим на [-1] - на случай index 0
	for i := 0; i < index; i++ {
		cur = cur.next // если index == 0, тогда остаёмся на dummy
	}

	cur.next = cur.next.next // если index == 0, dummy.next (cur.next) [указатель у -1 элемента] начинается ссылаться на указатель у 1 элемента, [0] забыт
	ll.head = dummy.next
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
