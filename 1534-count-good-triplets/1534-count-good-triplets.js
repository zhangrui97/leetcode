/**
 * @param {number[]} arr
 * @param {number} a
 * @param {number} b
 * @param {number} c
 * @return {number}
 */
var countGoodTriplets = function(arr, a, b, c) {
  const l = arr.length
  let ans = 0
  for (let i = 0; i < l - 2; i++) {
    const ai = arr[i]
    for (let j = i + 1; j < l - 1; j++) {
      const bj = arr[j]
      if (ai - bj > a || bj - ai > a) continue
      for (let k = j + 1; k < l; k++) {
        const ck = arr[k]
        if (bj - ck >= -b && bj - ck <= b && ai - ck >= -c && ai - ck <= c) ans++
      }
    }
  }
  return ans
};