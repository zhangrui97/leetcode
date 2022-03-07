/**
 * @param {number[]} arr
 * @return {number[]}
 */
var replaceElements = function(arr) {
  const last = arr.length - 1
  let maxIndex = -1
  for (let i = 0; i < last; i++) {
    if (maxIndex > i) {
      arr[i] = arr[maxIndex]
     } else {
       const sub = arr.slice(i + 1)
       const max = Math.max(...sub)
       maxIndex = i + 1 + sub.indexOf(max)
       arr[i] = max
     }
  }
  arr[last] = -1
  return arr
};