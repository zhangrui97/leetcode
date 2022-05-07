type Solution struct {
  size int
  blackMap map[int]int  
}


func Constructor(n int, blacklist []int) Solution {
  rand.Seed(42)
  l := len(blacklist)
  dictSize := l
  if l + l > n {
    dictSize = n - l
  }
  dict := make(map[int]int, dictSize)
  if l == 0 {
    return Solution{n, dict}
  }
  size := n - l
  sort.Ints(blacklist)
  l, r := 0, l - 1
  last := n -1
  for l <= r && blacklist[l] < last {
    for l < r && last == blacklist[r] {
      last--
      r--
    }
    if l <= r {
      dict[blacklist[l]] = last
      l++
      last--
    }
  }
  return Solution{size, dict}
}


func (this *Solution) Pick() int {
  val := rand.Intn(this.size)  
  if v, ok := this.blackMap[val]; ok {
    return v
  }
  return val
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */