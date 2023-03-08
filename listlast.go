package piscine

// import "fmt"

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListLast(l *List) interface{} {
	if l.Head != nil {
		return l.Tail.Data
	}
	return nil
}

// func ListPushBack(l *List, data interface{}) {
// 	n := &NodeL{Data: data}
// 	if l.Head == nil {
// 		l.Head = n
// 		l.Tail = n
// 	} else {
// 		l.Tail.Next = n
// 		l.Tail = n
// 	}
// }

// func main() {
// 	link := &List{}
// 	link2 := &List{}

// 	ListPushBack(link, "three")
// 	ListPushBack(link, 3)
// 	ListPushBack(link, "2")

// 	fmt.Println(ListLast(link))
// 	fmt.Println(ListLast(link2))
// }
