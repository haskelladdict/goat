// conncomp provides a data structure for representing the connected
// components of an undirected graph
package conncomp

// ConnComp desribes the connected components of a Ugraph
type ConnComp [][]int


// Connected returns if vertices x and y are connected
func (cc ConnComp) Connected(x int, y int) bool {
  for _, vs := range cc {
    found := false
    for _, v := range vs {
      if v == x || v == y {
        if found {
          return true
        } else {
          found = true
        }
      }
    }
  }
  return false
}
