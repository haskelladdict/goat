// queue implements a simple fifo queue based on a doubly linked list
package queue

import (
  "fmt"
)


// Element stores a single queue item
type Element struct {
  next *Element
  prev *Element
  Value interface{}
}


// Queue represents the queue data structure
type Queue struct {
  front Element     // we keep a sentinel and use a circular
                    // list to simplify them implementation
  len int           // current length of linked list
}


// New returns a pointer to a new queue
func New() *Queue {
  queue := new(Queue)
  queue.front.next = &queue.front
  queue.front.prev = &queue.front

  return queue
}


// Len returns the current length of the queue
func (q *Queue) Len() int {
  return q.len
}


// Print prints the current content of the queue
func (q *Queue) Print() {
  for e := q.front.next; e != &q.front; e = e.next {
    fmt.Print(e.Value," ")
  }
  fmt.Print("\n")
}


// Enqueue adds another element to the queue
func (q *Queue) Enqueue(e interface{}) {
  elem := Element{q.front.next, &q.front, e}
  q.front.next.prev = &elem
  q.front.next = &elem
  q.len++
}


// Dequeue removes the oldest element from the queue
func (q *Queue) Dequeue() interface{} {
  if q.front.next == &q.front {
    return nil
  }

  out_elem := q.front.prev
  q.front.prev = out_elem.prev
  out_elem.prev.next = &q.front
  q.len--

  return out_elem.Value
}

