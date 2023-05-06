package module4

import (
	"errors"
	"fmt"
)

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(a string) {
	*s = append(*s, a)
}

func (s *Stack) Pop() error {
	if s.IsEmpty() {
		return errors.New("Stack is empty")
	} else {
		*s = (*s)[:len(*s)-1]
		return nil
	}
}

func (s *Stack) Last() interface{} {
	if s.IsEmpty() {
		return nil
	} else {
		e := (*s)[len(*s)-1]
		return e
	}
}

func PSPcheck(s string) int {
	if len(s) == 0 {
		return 0
	}
	var st Stack
	st.Push(string(s[0]))
	for i := 1; i < len(s); i++ {
		if string(s[i]) == ")" && st.Last() == "(" {
			st.Pop()
		} else {
			st.Push(string(s[i]))
		}
	}
	return len(st)
}

type Item struct {
	Value int
	Next  *Item
}

type List struct {
	Head *Item
	Tail *Item
	Size int
}

func (l *List) Add(item Item) {
	if l.Size == 0 {
		l.Head = &item
	} else {
		l.Tail.Next = &item
	}
	l.Tail = &item
	l.Size++
}

func (l *List) GetByIndex(index int) (*Item, error) {
	if index >= l.Size {
		return nil, errors.New("Index out of bound")
	}
	if index == l.Size-1 {
		return l.Tail, nil
	}
	if index == 0 {
		return l.Head, nil
	}
	item := l.Head
	for i := 0; i <= index-1; i++ {
		item = item.Next
	}
	return item, nil
}

func (l *List) PopHead() {
	l.Head = l.Head.Next
	l.Size--
}

func (l *List) PopTail() {
	item := l.Head
	for i := 0; i < l.Size-2; i++ {
		item = item.Next
	}
	item.Next = nil
	l.Tail = item
	l.Size--
}

func Nearest_Lowest() {
	arr := List{
		Head: nil,
		Tail: nil,
		Size: 0,
	}
	var n, a int
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		arr.Add(Item{
			Value: a,
			Next:  nil,
		})
	}
	var min int = 1e9 + 1
	c := arr.Head
	for i := 0; i < arr.Size; i++ {
		if min > c.Value {
			min = c.Value
		}
		c = c.Next
	}
	lens := arr.Size
	for i := 0; i < lens; i++ {
		found := false
		b := arr.Head.Value
		arr.PopHead()
		new := arr.Head
		if b == min {
			fmt.Print("-1 ")
			continue
		}
		for a := i + 1; a < lens; a++ {
			if b > new.Value {
				found = true
				fmt.Print(a, " ")
				break
			}
			new = new.Next
		}
		if !found {
			fmt.Print("-1 ")
		}
	}
}

func Nearest_Lowest_Dumb_method() {
	var arr []int
	var n, a int
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		arr = append(arr, a)
	}
	for i := 0; i < len(arr); i++ {
		found := false
		for a := i + 1; a < len(arr); a++ {
			if arr[a] < arr[i] {
				found = true
				fmt.Print(a, " ")
				break
			}
		}
		if !found {
			fmt.Print("-1 ")
		}
	}
}

func MinWindowsLol(n []int, k int) {
	l := List{
		Head: nil,
		Tail: nil,
		Size: 0,
	}
	for i := 0; i < k-1; i++ {
		l.Add(Item{
			Value: n[i],
			Next:  nil,
		})
	}
	for i := 0; i < len(n)-k+1; i++ {
		l.Add(Item{
			Value: n[i+k-1],
			Next:  nil,
		})
		min := l.Head.Value
		it := l.Head
		for i := 0; i < k; i++ {
			if it.Value < min {
				min = it.Value
			}
			it = it.Next
		}
		fmt.Println(min)
		l.PopHead()
	}
}
