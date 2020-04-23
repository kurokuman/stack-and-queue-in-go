package queue

import "errors"

type item struct {
	value string
	next  *item
}

type Queue struct {
	top    *item
	bottom *item
	size   int
}

func (queue *Queue) Size() int {
	return queue.size
}

func (queue *Queue) Enqueue(value string) {

	newItem := item{
		value: value,
		next:  nil,
	}
	if queue.Size() < 1 {
		queue.top = &newItem
		queue.bottom = queue.top
	} else {
		queue.bottom.next = &newItem
		queue.bottom = &newItem
	}
	queue.size++
}

func (queue *Queue) Dequeue() (string, error) {
	if queue.Size() < 1 {
		return "", errors.New("No more items")
	}
	value := queue.top.value
	queue.top = queue.top.next
	queue.size--
	return value, nil
}
