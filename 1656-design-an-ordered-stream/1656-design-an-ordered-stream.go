type OrderedStream struct {
  n int
  buffer []string
  ptr int
}


func Constructor(n int) OrderedStream {
  return OrderedStream{n, make([]string, n), 0}  
}


func (this *OrderedStream) Insert(idKey int, value string) []string {
  this.buffer[idKey - 1] = value
  if idKey - this.ptr > 1 {
    return []string{}
  } else {
    largest := idKey
    for largest < this.n && this.buffer[largest] != "" {
      largest++
    }
    from := this.ptr
    this.ptr = largest
    return this.buffer[from:this.ptr]
  }
}


/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */