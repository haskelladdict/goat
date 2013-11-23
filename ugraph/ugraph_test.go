package ugraph

import (
  //"fmt"
  "log"
  "testing"
)


var graph Ugraph
func init() {
  var err error
  graph, err = New_parse("graph_1.txt")
  if err != nil {
    log.Fatal(err)
  }
}


func TestUgraph(t *testing.T) {

  if graph.V() != 13 {
    t.Errorf("%d != 13 -- incorrect number of vertices", graph.V());
  }
}
