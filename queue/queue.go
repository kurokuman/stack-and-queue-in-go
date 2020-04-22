package queue

import "errors"

type item struct {
	value string
	prev  *item
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

	if queue.Size() < 1 {
		newItem := item{
			value: value,
			prev:  nil,
			next:  nil,
		}
		queue.top = &newItem
		queue.bottom = queue.top
	} else {
		newItem := item{
			value: value,
			prev:  queue.bottom,
			next:  nil,
		}
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
