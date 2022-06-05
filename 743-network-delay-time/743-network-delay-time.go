import (
  "container/heap"
)

func max(a, b int)int {
  if a > b {
    return a
  } else {
    return b
  }
}

type TimeHeap struct {
  times [][2]int
}

func (this TimeHeap)Len()int { return len(this.times) }
func (this TimeHeap)Less(i, j int)bool { return this.times[i][1] < this.times[j][1] }
func (this TimeHeap)Swap(i, j int) { this.times[i], this.times[j] = this.times[j], this.times[i] }

func (this *TimeHeap)Push(a interface{}) {
  this.times = append(this.times, a.([2]int))
}

func (this *TimeHeap)Pop() interface{} {
  old := this.times
  n := this.Len() - 1
  v := old[n]
  this.times = old[:n]
  return v
}

func networkDelayTime(times [][]int, n int, k int) int {
  ans := -1
  th := TimeHeap{[][2]int{}}
  heap.Init(&th)
  g := make(map[int][][2]int)
  for _, v := range times {
    if v[0] == k {
      heap.Push(&th, [2]int{v[1], v[2]})
    } else if v[1] != k {
      g[v[0]] = append(g[v[0]], [2]int{v[1], v[2]})
    }
  }
  visited := make(map[int]bool)
  for n > 1 && th.Len() > 0 {
    min := heap.Pop(&th).([2]int)
    if visited[min[0]] { continue }
    i, t := min[0], min[1]
    ans = max(ans, t)
    visited[i] = true
    n--
    for _, v := range g[i] {
      heap.Push(&th, [2]int{v[0], t + v[1]})
    }
  }
  if n > 1 {
    return -1
  } else {
    return ans
  }
}