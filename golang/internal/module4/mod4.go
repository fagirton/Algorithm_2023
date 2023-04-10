package module4

import "errors"

type Stack []interface{}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(a interface{}) []interface{} {
	return append(*s, a)
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
	var st Stack
	st.Push(s[0])
	for i := 1; i < len(s); i++ {
		if string(s[i]) == ")" && st.Last() == "(" {
			st.Pop()
		} else {
			st.Push(s[i])
		}
	}
	return len(st)
}
