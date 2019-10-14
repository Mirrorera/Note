package d2

import "fmt"

type Node struct {
	data int
	next *Node
}

func (this *Node) Add(p *Node) {
	p.next = this.next
	this.next = p
}
func (this *Node) Next() *Node {
	return this.next
}
func d226() {
	var A, B *Node
	var sa, sb string

	fmt.Scanln(sa)
	fmt.Scanln(sb)
	for _, x := range sa {
		A.Add(Node{sa - '0', nil})
	}
}
