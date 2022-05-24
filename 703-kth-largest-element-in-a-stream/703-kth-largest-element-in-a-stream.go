type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
  data *IntHeap
  k int
}

func Constructor(k int, nums []int)KthLargest {
  sort.Slice(nums, func(i, j int)bool {
    return nums[i] < nums[j]
  })
  l := len(nums)
  var d IntHeap
  if l > k {
    d = IntHeap(nums[l-k:])
  } else {
    d = IntHeap(nums)
  }
  heap.Init(&d)
  return KthLargest{&d, k}
}

func (this *KthLargest) Add(val int) int {
  if len(*this.data) < this.k {
    heap.Push(this.data, val)
    return (*this.data)[0]
  }
  if val < (*this.data)[0] { return (*this.data)[0] }
  heap.Pop(this.data)
  heap.Push(this.data, val)
  return (*this.data)[0]
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */