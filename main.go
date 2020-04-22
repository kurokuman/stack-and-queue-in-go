package main

import (
	"fmt"

	"github.com/kurokuman/stack-and-queue-in-go/queue"
	"github.com/kurokuman/stack-and-queue-in-go/stack"
)

func main() {
	stack := stack.Stack{}
	queue := queue.Queue{}

	valueList := []string{"a", "b", "c", "d", "e", "f", "g"}

	for _, value := range valueList {
		stack.Push(value)
		queue.Enqueue(value)
	}

	for {
		sValue, sErr := stack.Pop()
		if sErr != nil {
			fmt.Println(sErr)
			break
		}
		fmt.Printf("stack pop value : %s \n", sValue)
	}

	fmt.Println("---------------------------------")

	for {
		qValue, qErr := queue.Dequeue()
		if qErr != nil {
			fmt.Println(qErr)
			break
		}
		fmt.Printf("queue dequeue value : %s \n", qValue)
	}
}
