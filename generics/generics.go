package main

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val T
}

// receiver function Push on List
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val : v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// receiver function GetAll on List
func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	m := map[int]string{1: "2", 2: "4", 4: "8"}
	fmt.Println("Keys: ", MapKeys(m))

	_ = MapKeys[int, string](m)

	lst := List[int]{}
	lst.Push(10)
	lst.Push(20)
	lst.Push(30)
	fmt.Println(lst.GetAll())
}
