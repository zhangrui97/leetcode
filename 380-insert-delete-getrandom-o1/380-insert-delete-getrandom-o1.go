type RandomizedSet struct {
  dict map[int]int
  buffer []int
}


func Constructor() RandomizedSet {
  rand.Seed(42)
  return RandomizedSet{make(map[int]int), make([]int, 0)}
}


func (this *RandomizedSet) Insert(val int) bool {
  if _, ok := this.dict[val]; ok {
    return false
  } else {
    this.dict[val] = len(this.buffer)
    this.buffer = append(this.buffer, val)
    return true
  }
}


func (this *RandomizedSet) Remove(val int) bool {
  if v, ok := this.dict[val]; ok {
    l := len(this.buffer)
    tail := this.buffer[l-1]
    this.buffer[v] = tail
    this.buffer = this.buffer[:l-1]
    this.dict[tail] = v
    delete(this.dict, val)
    return true
  } else {
    return false
  }
}


func (this *RandomizedSet) GetRandom() int {
  return this.buffer[rand.Intn(len(this.buffer))]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */