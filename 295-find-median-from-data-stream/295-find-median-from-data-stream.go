import (
    "container/heap"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
  *h = append(*h, x.(int))
}

func (h *MinHeap) Peek() int {
  old := *h
  return old[0]
}

func (h *MinHeap) Pop() interface{} {
  old := *h
  n := len(old)
  x := old[n-1]
  *h = old[0 : n-1]
  return x
}

type MaxHeap struct {
  MinHeap
}

func (h MaxHeap) Less(i, j int) bool { return h.MinHeap[i] > h.MinHeap[j] }

type MedianFinder struct {
  fst *MaxHeap
  snd *MinHeap
}


func Constructor() MedianFinder {
  fst := MaxHeap{}
  heap.Init(&fst)
  snd := MinHeap{}
  heap.Init(&snd)
  return MedianFinder{&fst, &snd}
}


func (this *MedianFinder) AddNum(num int)  {
  fl, sl := this.fst.Len(), this.snd.Len()
  if sl == 0 {
    heap.Push(this.snd, num)
    return
  }
  if fl == 0 {
    if num <= this.snd.Peek() {
      heap.Push(this.fst, num)
    } else {
      heap.Push(this.fst, heap.Pop(this.snd))
      heap.Push(this.snd, num)
    }
    return
  }
  if sl == fl {
    if num >= this.fst.Peek() {
      heap.Push(this.snd, num)
    } else {
      heap.Push(this.snd, heap.Pop(this.fst))
      heap.Push(this.fst, num)
    }
  } else { // sl = fl + 1
    if num >= this.snd.Peek() {
      heap.Push(this.fst, heap.Pop(this.snd))
      heap.Push(this.snd, num)
    } else {
      heap.Push(this.fst, num)
    }
  }
}


func (this *MedianFinder) FindMedian() float64 {
  if this.fst.Len() == this.snd.Len() {
    return (float64(this.fst.Peek()) + float64(this.snd.Peek())) / 2.0
  } else {
    return float64(this.snd.Peek())
  }
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */