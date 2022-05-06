type Node struct {
  key int
  freq int
  value int
  pre *Node
  nxt *Node
}

type NodeSet struct {
  head *Node
  tail *Node
}

type LFUCache struct {
  keyDict map[int]*Node
  freqDict map[int]*NodeSet
  cap int
  size int
  leastFreq int
}


func Constructor(capacity int) LFUCache {
  return LFUCache{make(map[int]*Node, capacity),
                  make(map[int]*NodeSet, capacity),
                  capacity, 0, capacity}
}

func (this *LFUCache) IncFrequency(node *Node) {
  set := this.freqDict[node.freq]
  if set.head == set.tail {
    delete(this.freqDict, node.freq)
    if this.leastFreq == node.freq {
      this.leastFreq++
    }
  } else {
    if node.pre != nil && node.nxt != nil {
      node.nxt.pre = node.pre
      node.pre.nxt = node.nxt
    } else if node.pre != nil {
      node.pre.nxt = nil
      set.tail = node.pre
    } else {
      node.nxt.pre = nil
      set.head = node.nxt
    }
  }
  
  node.freq++
  
  if newSet, ok := this.freqDict[node.freq]; ok {
    node.pre = nil
    node.nxt = newSet.head
    newSet.head.pre = node
    newSet.head = node
  } else {
    node.pre = nil
    node.nxt = nil
    this.freqDict[node.freq] = &NodeSet{node, node}
  }
}

func (this *LFUCache) RemoveLeastFreq() {
  set := this.freqDict[this.leastFreq]
  tail := set.tail
  if set.head == set.tail {
    delete(this.freqDict, this.leastFreq)
  } else {
    tail.pre.nxt = nil
    set.tail = tail.pre
  }
  delete(this.keyDict, tail.key)
}

func (this *LFUCache) Get(key int) int {
  if v, ok := this.keyDict[key]; ok {
    this.IncFrequency(v)
    return v.value
  }
  return -1
}


func (this *LFUCache) Put(key int, value int)  {
  if v, ok := this.keyDict[key]; ok {
    v.value = value
    this.IncFrequency(v)
  } else {
    if this.cap == 0 {
      return
    }
    node := &Node{key, 1, value, nil, nil}
    this.keyDict[key] = node
    if set, ok := this.freqDict[1]; ok {
      node.nxt = set.head
      set.head.pre = node
      set.head = node
    } else {
      this.freqDict[1] = &NodeSet{node, node}
    }
    if this.size < this.cap {
      this.size++
    } else {
      this.RemoveLeastFreq()
    }
    if this.leastFreq > 1 {
      this.leastFreq = 1
    }
  }
}


/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */