/**
 * @param {number} n
 */
var OrderedStream = function(n) {
  this.n = n
  this.buffer = Array(n+1)
  this.ptr = 1
};

/** 
 * @param {number} idKey 
 * @param {string} value
 * @return {string[]}
 */
OrderedStream.prototype.insert = function(idKey, value) {
  this.buffer[idKey] = value
  if (idKey > this.ptr) {
    return []
  } else if (idKey == this.ptr) {
    let largest = idKey
    while (largest <= this.n && this.buffer[largest]) largest++
    const start = this.ptr
    this.ptr = largest
    return this.buffer.slice(start, largest)
  } else { // idKey < this.ptr
    // impossible
  }
};

/** 
 * Your OrderedStream object will be instantiated and called as such:
 * var obj = new OrderedStream(n)
 * var param_1 = obj.insert(idKey,value)
 */