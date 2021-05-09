package main

import (
	"fmt"
)

type Queue struct {
	elements []int
}

func (q *Queue) Enqueue(i int) {
	q.elements = append(q.elements, i)
}
func (q *Queue) Dequeue() int {
	toRemove := q.elements[0]
	q.elements = q.elements[1:]
	return toRemove
}

func main() {
	myQueue := Queue{}
	fmt.Println(myQueue)
	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)
}



