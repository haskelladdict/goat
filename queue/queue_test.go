// queue implements a simple fifo queue based on a doubly linked list
package queue

import (
  "testing"
)

// a few simple tests
func TestQueue(t *testing.T) {

  q := New()

  q.Enqueue(1)
  q.Enqueue(2)
  q.Enqueue(3)

  if q.Len() != 3 {
    t.Error("Incorrect queue length")
  }

  if q.Dequeue() != 1 {
    t.Error("Error dequeing")
  }

  if q.Dequeue() != 2 {
    t.Error("Error dequeing")
  }

  if q.Dequeue() != 3 {
    t.Error("Error dequeing")
  }

  if q.Dequeue() != nil {
    t.Error("Error dequeing")
  }

  if q.Dequeue() != nil {
    t.Error("Error dequeing")
  }


}

