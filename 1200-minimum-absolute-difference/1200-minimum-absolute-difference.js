/**
 * @param {number[]} arr
 * @return {number[][]}
 */
var minimumAbsDifference = function(arr) {
  if (arr.length === 2) return [arr]
  arr.sort((a, b) => a-b)
  const diff = new Array(arr.length -1)
  for (let i = 1; i < arr.length; i++) {
    diff[i - 1] = [arr[i] - arr[i-1], i]
  }
  diff.sort(([a, i], [b, j]) => a - b)
  const min = diff[0][0]
  return diff.filter(([a, i]) => a === min).map(([_, i]) => [arr[i-1], arr[i]])
};