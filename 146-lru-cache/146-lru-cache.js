/**
 * @param {number} capacity
 */
var LRUCache = function(capacity) {
  this.cap = capacity
  this.map = new Map()
  this.buffer = Array(capacity).fill([])
  this.head = 0
  this.count = 0
};

LRUCache.prototype.makeHead = function(idx) {
  if (this.head === idx) return
  
  const tail = this.buffer[this.head][2]
  const [_, v, pre, nxt] = this.buffer[idx]
  if (idx !== tail) {
    this.buffer[pre][3] = nxt
    this.buffer[nxt][2] = pre
    this.buffer[idx][2] = tail
    this.buffer[idx][3] = this.head
    this.buffer[this.head][2] = idx
    this.buffer[tail][3] = idx
  }
  this.head = idx
}

/** 
 * @param {number} key
 * @return {number}
 */
LRUCache.prototype.get = function(key) {
  if (this.map.has(key)) {
    const idx = this.map.get(key)
    this.makeHead(idx)
    return this.buffer[this.head][1]
  }
  return -1
};

/** 
 * @param {number} key 
 * @param {number} value
 * @return {void}
 */
LRUCache.prototype.put = function(key, value) {
  if (this.map.has(key)) {
    const idx = this.map.get(key)
    this.buffer[idx][1] = value
    this.makeHead(idx)
  } else {
    const tail = this.count === 0 ? 0 : this.buffer[this.head][2]
    if (this.count < this.cap) {
      this.buffer[this.count] = [key, value, tail, this.head]
      this.buffer[this.head][2] = this.count
      this.buffer[tail][3] = this.count
      this.head = this.count
      this.count++
    } else {
      this.map.delete(this.buffer[tail][0])
      this.buffer[tail][0] = key
      this.buffer[tail][1] = value
      this.head = tail
    }
    this.map.set(key, this.head)
  }
};

/** 
 * Your LRUCache object will be instantiated and called as such:
 * var obj = new LRUCache(capacity)
 * var param_1 = obj.get(key)
 * obj.put(key,value)
 */