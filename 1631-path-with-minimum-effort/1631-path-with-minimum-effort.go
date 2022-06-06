import (
  "container/heap"
)
type Path struct {
  Type int
  Index int
  Effort int
}
type PathHeap struct {
  Paths []Path
}
func (this PathHeap) Len()int { return len(this.Paths) }
func (this PathHeap) Less(i, j int)bool { return this.Paths[i].Effort < this.Paths[j].Effort }
func (this PathHeap) Swap(i, j int) { this.Paths[i], this.Paths[j] = this.Paths[j], this.Paths[i] }

func (this *PathHeap) Push(a interface{}) {
  this.Paths = append(this.Paths, a.(Path))
}
func (this *PathHeap) Pop()interface{} {
  old := this.Paths
  n := this.Len() - 1
  v := old[n]
  this.Paths = old[:n]
  return (interface{})(v)
}

func diff(a, b int)int {
  if a > b {
    return a - b
  } else {
    return b - a
  }
}

func max(a, b int)int {
  if a > b {
    return a
  } else {
    return b
  }
}

func minimumEffortPath(heights [][]int) int {
  r, c := len(heights), len(heights[0])
  visited := [2]map[int]int{make(map[int]int), make(map[int]int)}
  ph := PathHeap{[]Path{Path{0, 0, 0}, Path{1, (r-1)<<16|(c-1), 0}}}
  heap.Init(&ph)
  updateEffort := func(path Path, l, k int)int {
    i, j := path.Index >> 16, path.Index & 0xFFFF
    return max(path.Effort, diff(heights[i][j], heights[l][k]))    
  }
  for ph.Len() > 0 {
    path := heap.Pop(&ph).(Path)
    if path.Type == 2 { return path.Effort }
    visited[path.Type][path.Index] = path.Effort
    if v, ok := visited[(path.Type+1)%2][path.Index]; ok {
      heap.Push(&ph, Path{2, path.Index, max(path.Effort, v)}) 
      continue
    }
    i, j := path.Index >> 16, path.Index & 0xFFFF
    if i > 0 { 
      if _, ok := visited[path.Type][(i-1)<<16|j]; !ok {
        heap.Push(&ph, Path{path.Type, (i-1)<<16|j, updateEffort(path, i-1, j)}) 
      }
    }
    if i < r-1 {
      if _, ok := visited[path.Type][(i+1)<<16|j]; !ok { 
        heap.Push(&ph, Path{path.Type, (i+1)<<16|j, updateEffort(path, i+1, j)}) 
      }
    }
    if j > 0 {
      if _, ok := visited[path.Type][i<<16|(j-1)]; !ok { 
        heap.Push(&ph, Path{path.Type, i<<16|(j-1), updateEffort(path, i, j-1)}) 
      }
    }
    if j < c-1 {
      if _, ok := visited[path.Type][i<<16|(j+1)]; !ok { 
        heap.Push(&ph, Path{path.Type, i<<16|(j+1), updateEffort(path, i, j+1)}) 
      }
    }
  }
  return -1
}