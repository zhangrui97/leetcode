import (
  "container/heap"
)

type Prob struct {
  Node int
  Prob float64
}

type ProbHeap struct {
  Data []Prob  
}

func (this ProbHeap)Len()int { return len(this.Data) }
func (this ProbHeap)Less(i, j int)bool { return this.Data[i].Prob > this.Data[j].Prob }
func (this ProbHeap)Swap(i, j int) { this.Data[i], this.Data[j] = this.Data[j], this.Data[i] }
func (this *ProbHeap)Push(v interface{}) { this.Data = append(this.Data, v.(Prob)) }
func (this *ProbHeap)Pop() interface{} {
  old := this.Data
  n := this.Len() - 1
  v := old[n]
  this.Data = old[:n]
  return (interface{})(v)
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
  nodes := make([][]Prob, n)
  for i, v := range edges {
    nodes[v[0]] = append(nodes[v[0]], Prob{v[1], succProb[i]})
    nodes[v[1]] = append(nodes[v[1]], Prob{v[0], succProb[i]})
  }
  
  visited := make([]bool, n)

  ph := ProbHeap{[]Prob{Prob{start, 1}}}
  heap.Init(&ph)
  
  for ph.Len() > 0 {
    prob := heap.Pop(&ph).(Prob)
    if prob.Node == end { return prob.Prob }
    visited[prob.Node] = true
    for _, v := range nodes[prob.Node] {
      if !visited[v.Node] {
        heap.Push(&ph, Prob{v.Node, v.Prob * prob.Prob})
      }
    }
  }
  return 0
}