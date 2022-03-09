/**
 * @param {string[]} ops
 * @return {number}
 */
var calPoints = function(ops) {
  const buffer = []
  ops.forEach(op => {
    switch (op) {
      case 'C':
        buffer.pop()
        return
      case 'D':
        buffer.push(2 * buffer.at(-1))
        return
      case '+':
        buffer.push(buffer.at(-1) + buffer.at(-2))
        return
      default: buffer.push(+op)
    }
  })
  return buffer.reduce((a, b) => a + b)
};