/**
 * @param {number[]} arr
 * @return {number}
 */
var sumOddLengthSubarrays = function(arr) {
  const l = arr.length
  const sum = arr.reduce((a, b) => a + b)
  const subSub = len => {
    if (len === 1) return sum
    if (len === l) return sum
    const factor = l >= 2 * len ? len - 1: l - len
    let ans = (factor + 1) * sum
    for (let i = 0; i < factor; i++) {
      ans -= (factor - i) * (arr[i] + arr[l - 1 - i])
    }
    return ans
  }
  let result = 0
  for (let i = 1; i <= l; i += 2) {
    result += subSub(i)
  }
  return result
};