package main

import (
  "fmt"
  "goat/queue"
  "goat/ugraph"
)

func main() {

  fmt.Println("hi there")

  g, err := ugraph.New_parse("graph_1.txt")
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(g)
  fmt.Println("number of vertices ", g.V())
  fmt.Println("number of edges ", g.E())
  fmt.Println("vertices adjacent to vertex 6 ", g.Adj(6))
  fmt.Println("max degree", g.Max_degree(), " avg degree ", g.Avg_degree())
  fmt.Println("number of self loops ", g.Num_selfloops())
  cc := g.Conn_components()
  fmt.Println("number of connected components ", cc)
  fmt.Println("conntected 0 6 ", cc.Connected(12,9)) 

  q := queue.New()
  q.Print()
  q.Enqueue(5)
  q.Enqueue(12)
  q.Enqueue(77)
  q.Print()
  fmt.Println(q.Dequeue())
  q.Enqueue(3.4)
  e := q.Dequeue()
  for ; e != nil; e = q.Dequeue() {
    fmt.Println(e)
  }
}
