type UF struct {
  size map[[2]int]int
  parent map[[2]int][2]int
  max int
}

func (this *UF)Find(a [2]int)[2]int {
  for this.parent[a] != a {
    this.parent[a] = this.parent[this.parent[a]]
    a = this.parent[a]
  }
  return a
}

func (this *UF)Union(a, b [2]int) {
  rootA := this.Find(a)
  rootB := this.Find(b)
  if rootA != rootB {
    this.parent[rootB] = rootA
    size := this.size[rootA] + this.size[rootB]
    this.size[rootA] = size
    if size > this.max {
      this.max = size
    }
  }
}

func (this *UF)MaxSize() int {
  return this.max
}

func NewUF(grid [][]int)*UF {
  size := make(map[[2]int]int)
  parent := make(map[[2]int][2]int)
  for i, r := range grid {
    for j, v := range r {
      if v == 1 {
        parent[[2]int{i, j}] = [2]int{i, j}
        size[[2]int{i, j}] = 1
      }
    }
  }
  s := 1
  if len(size) == 0 { s = 0 }
  return &UF{size, parent, s}
}

func maxAreaOfIsland(grid [][]int) int {
  uf := NewUF(grid)
  for k, _ := range uf.parent {
    r, c := k[0], k[1]
    if r > 0 && grid[r-1][c] == 1 {
      uf.Union([2]int{r-1, c}, k)
    }
    if c > 0 && grid[r][c-1] == 1 {
      uf.Union([2]int{r, c-1}, k)
    }
  }
  return uf.MaxSize()
}