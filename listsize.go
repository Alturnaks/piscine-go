package piscine

// import (
// 	"fmt"
// )

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListSize(l *List) int {
	it := 0
	for l.Head != nil {
		l.Head = l.Head.Next
		it++
	}
	return it
}

// func main() {
// 	link := &List{}

// 	ListPushFront(link, "Hello")
// 	ListPushFront(link, "2")
// 	ListPushFront(link, "you")
// 	ListPushFront(link, "man")

// 	fmt.Println(ListSize(link))
// }

// \\func main() {
// 	link := &List{}

// 	ListPushFront(link, "Hello")
// 	ListPushFront(link, "man")
// 	ListPushFront(link, "how are you")

// 	it := link.Head
// 	for it != nil {
// 		fmt.Print(it.Data, " ")
// 		it = it.Next
// 	}
// 	fmt.Println()
// }
