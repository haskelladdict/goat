package main

import (
  "fmt"
  "github.com/haskelladdict/goat/queue"
  "github.com/haskelladdict/goat/ugraph"
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
  fmt.Println("conntected 0 9 ", g.Connected(0, 9))

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

  paths := g.Compute_paths(7)
  fmt.Println(paths)
  fmt.Printf("0 %t 1 %t   2 %t   3 %t   4 %t   5 %t   6 %t  7 %t  8 %t   9 %t \n",
    paths.Has_path_to(0),
    paths.Has_path_to(1), paths.Has_path_to(2), paths.Has_path_to(3),
    paths.Has_path_to(4), paths.Has_path_to(5), paths.Has_path_to(6),
    paths.Has_path_to(7), paths.Has_path_to(8), paths.Has_path_to(9))

  fmt.Println(g.Path_to(7, 8))

  g2, err := ugraph.New_parse("graph_2.txt")
  fmt.Println(g2.Num_cycles())

}
