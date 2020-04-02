package stackqueue

type Queue []string

func (q *Queue) Enqueue(value string) *Queue {
	*q = append([]string{value}, *q...)
	return q
}

func (q *Queue) Dequeue() string {
	value := (*q)[len(*q) - 1]
	*q = (*q)[0:len(*q) - 1]
	return value
}
