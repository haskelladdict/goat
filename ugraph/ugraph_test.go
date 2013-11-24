package ugraph

import (
  //"fmt"
  "log"
  "testing"
)


var graph1 Ugraph
func init() {
  var err error
  graph1, err = New_parse("test_files/graph_1.txt")
  if err != nil {
    log.Fatal(err)
  }
}


// TestUGraphBasic tests basic undirected graph properties
func TestUgraphBasic(t *testing.T) {

  num_verts := 13
  if graph1.V() != num_verts {
    t.Errorf("incorrect number of vertices A(%d)/T(%d)", graph1.V(), 
      num_verts);
  }

  num_edges := 13
  if graph1.E() != num_edges {
    t.Errorf("incorrect number of edges A(%d)/T(%d)", graph1.V(), num_edges);
  }

  num_adj_verts_4 := 3
  if len(graph1.Adj(4)) != num_adj_verts_4 {
    t.Errorf("incorrect number of vertices adjacent to vertex 4 A(%d)/T(%d)",
      len(graph1.Adj(4)), num_adj_verts_4);
  }

  max_degree := 4
  if graph1.Max_degree() != max_degree {
    t.Errorf("incorrect max degree A(%d)/T(%d)", graph1.Max_degree(),
      max_degree);
  }

  avg_degree := 2.0
  if graph1.Avg_degree() != avg_degree {
    t.Errorf("incorrect average degree A(%f)/T(%f)", graph1.Avg_degree(),
      avg_degree);
  }

  num_conn_comp := 3
  if len(graph1.Conn_components()) != num_conn_comp {
    t.Errorf("incorrect number of connected components A(%d)/T(%d)",
      len(graph1.Conn_components()), num_conn_comp);
  }
}


// TestUGraphPath tests undirected graph properties related to paths
func TestUgraphPath(t *testing.T) {

  // test paths from vertex 7
  paths_7 := graph1.Compute_paths(7)

  num_7_edges := 1
  if paths_7.num_edges() != num_7_edges {
    t.Errorf("incorrect number of paths from vertex 7 A(%d)/T(%d)",
      paths_7.num_edges(), num_7_edges);
  }

  if !paths_7.Has_path_to(8) {
    t.Errorf("vertex 7 has no path to vertex 8 but should have");
  }

  if paths_7.Has_path_to(0) {
    t.Errorf("vertex 7 has a path to vertex 0 but shouldn't have");
  }


  // test paths from vertex 0
  paths_0 := graph1.Compute_paths(0)

  num_0_edges := 6
  if paths_0.num_edges() != num_0_edges {
    t.Errorf("incorrect number of paths from vertex 0 A(%d)/T(%d)",
      paths_0.num_edges(), num_0_edges);
  }

  if !paths_0.Has_path_to(1) {
    t.Errorf("vertex 0 has no path to vertex 1 but should have");
  }

  if !paths_0.Has_path_to(3) {
    t.Errorf("vertex 0 has no path to vertex 3 but should have");
  }

  if paths_0.Has_path_to(8) {
    t.Errorf("vertex 0 has a path to vertex 8 but shouldn't have");
  }

  if paths_0.Has_path_to(12) {
    t.Errorf("vertex 0 has a path to vertex 12 but shouldn't have");
  }
}



// helper function to determined the number of edges in a Paths
// structure
func (p *Paths) num_edges() int {
  num_edges := 0
  for _, e := range p.edge_to {
    if e != -1 {
      num_edges++
    }
  }
  return num_edges
}



// TestUGraphConnect tests undirected graph properties related to vertex
// connectivities
func TestUgraphConnect (t *testing.T) {

  if graph1.Connected(0, 9) {
    t.Errorf("vertex 0 has a path to vertex 9 but shouldn't have");
  }

  if graph1.Connected(0, 8) {
    t.Errorf("vertex 0 has a path to vertex 8 but shouldn't have");
  }

  if !graph1.Connected(0, 6) {
    t.Errorf("vertex 0 has a no path to vertex 6 but should have");
  }

  if !graph1.Connected(0, 3) {
    t.Errorf("vertex 0 has a no path to vertex 3 but should have");
  }
}
