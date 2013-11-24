// ugraph provides a data structure and methods for undirected graphs. 
package ugraph

import (
  "bufio"
  "errors"
  "fmt"
  "os"
)


// Ugraph describes an undirected graph as adjacency list. 
// The graph is assumed to be unlabeled and vertices are indexed 
// via ints.
type Ugraph [][]int


// function for creating a new Ugraph with n vertices
func New_empty(num_vertices int) (Ugraph, error) {
  return make(Ugraph, num_vertices), nil
}


// function for creating a new Ugraph based on the content
// of supplied graph description file
//
// NOTE: The input file format is assumed to be
//   #vertices
//   #edges
//   v1 v2
//   v3 v4
//   ...
//
func New_parse(file_name string) (Ugraph, error) {

  // try to open and parse file
  file, ok := os.Open(file_name)
  if ok != nil {
    return nil, errors.New("Failed to open file")
  }
  scanner := bufio.NewScanner(file)

  // parse number of vertices and edges
  var n, m int
  scanner.Scan()
  if _, err := fmt.Sscanf(scanner.Text(), "%d", &n); err != nil {
    return nil, errors.New("Failed to parse number of vertices")
  }

  scanner.Scan()
  if _, err := fmt.Sscanf(scanner.Text(), "%d", &m); err != nil {
    return nil, errors.New("Failed to parse number of edges")
  }

  // initialize graph
  graph := make(Ugraph, n)

  // scan edges
  var x, y int
  for scanner.Scan() {
    if _, err := fmt.Sscanf(scanner.Text(), "%d %d", &x, &y); err != nil {
      return nil, errors.New("Failed to parse edge")
    }

    graph[x] = append(graph[x], y)
    graph[y] = append(graph[y], x)
  }
  if err := scanner.Err(); err != nil {
    return nil, errors.New("Failed to parse edges")
  }

  return graph, nil
}



// V returns the number of vertices in the graph
func (g Ugraph) V() int {
  return len(g)
}



// E returns the number of edges in the graph
func (g Ugraph) E() int {
  edges := 0
  for _, v := range g {
    edges += len(v)
  }
  return edges/2
}



// Adj returns the list of vertices adjacent to vertex v
func (g Ugraph) Adj(v int) []int {
  return g[v]
}



// Add_edge adds an edge between vertices v and w
func (g Ugraph) Add_edge(v int, w int) {
  g[v] = append(g[v], w)
  g[w] = append(g[w], v)
}



// Degree returns the degree of vertex v
func (g Ugraph) Degree(v int) int {
  return len(g[v])
}



// Max_degree returns the maximum degree of the graph
func (g Ugraph) Max_degree() int {
  max_deg := 0
  for _, v := range g {
    if len(v) > max_deg {
      max_deg = len(v)
    }
  }
  return max_deg
}



// Avg_degree returns the average degree of the graph
func (g Ugraph) Avg_degree() float64 {
  return 2.0*float64(g.E())/float64(g.V())
}


// Num_selfloops returns the number of self loops
func (g Ugraph) Num_selfloops() int {
  num_loops := 0
  for v, vs := range g {
    for _, w := range vs {
      if v == w {
        num_loops++
      }
    }
  }
  return num_loops/2
}



// Paths is a data structure keeping track of all connectivities
// from a source vertex in the graph
// NOTE: edge_to is set to -1 if there is no edge
type Paths struct {
  marked []bool
  edge_to []int
}



// compute_path determines all connections (paths) from source
// vertex s in the undirected graph. This functions returns a
// Path object.
func (g *Ugraph) Compute_paths(source int) *Paths {

  paths := Paths{}
  paths.marked = make([]bool, g.V())

  // edge_to slice is initialized to -1
  paths.edge_to = make([]int, g.V())
  for i := 0; i < g.V(); i++ {
    paths.edge_to[i] = -1
  }

  // compute connectivities via depth first search
  dfs_path(g, source, &paths)

  return &paths
}



// dfs_path computes all the vertices reachable in graph g 
// starting from source vertex s
func dfs_path(g *Ugraph, source int, paths *Paths) {

  paths.marked[source] = true
  for _, w := range g.Adj(source) {
    if !paths.marked[w] {
      dfs_path(g, w, paths)
      paths.edge_to[w] = source
    }
  }
}



// Has_path_to returns true/false depending on if there is a 
// path between s and v
func (p *Paths) Has_path_to(v int) bool {
  return p.marked[v]
}


// Path_to returns a vertex path between vertex s and t
// if one exists.
// NOTE: This will not return the shorted path - just a
// path!
func (g *Ugraph) Path_to(s, t int) ([]int, bool) {
  all_paths := g.Compute_paths(s)
  path := make([]int, 0)
  if !all_paths.marked[t] {
    return path, false
  }

  path = append(path, t)
  for p := all_paths.edge_to[t]; p != s; p = all_paths.edge_to[p] {
    path = append(path, p)
  }
  path = append(path, s)

  // reverse path
  for i := 0; i < len(path)/2; i++ {
    path[i], path[len(path)-i-1] = path[len(path)-i-1], path[i]
  }

  return path, true
}



// Connected returns true if two vertices are connected (i.e.
// in the same connected component and false otherwise
func (g *Ugraph) Connected(v, w int) bool {
  return g.Compute_paths(v).Has_path_to(w)
}



// Conn_components returns a slice of slices with vertices in each
// connected component 
func (g Ugraph) Conn_components() [][]int {
  components := make([][]int, 0)
  discovered := make([]bool, g.V())
  for i := 0; i < g.V(); i++ {
    if discovered[i] == false {
      cs := make([]int, 0)
      dfs_conn_components(g, i, discovered, &cs)
      components = append(components, cs)
    }
  }
  return components
}



// dfs helper function for Conn_components
func dfs_conn_components(g Ugraph, x int, discovered []bool, elems *[]int) {
  discovered[x] = true
  *elems = append(*elems, x)
  for _, y := range g.Adj(x) {
    if discovered[y] == false {
      dfs_conn_components(g, y, discovered, elems)
    }
  }
}








