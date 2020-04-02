package stackqueue

type Stack []string

func (s *Stack) Push(value string) *Stack {
	*s = append([]string{value}, *s...)
	return s
}

func (s *Stack) Pop() string {
	value := (*s)[0]
	*s = (*s)[1:]
	return value
}
