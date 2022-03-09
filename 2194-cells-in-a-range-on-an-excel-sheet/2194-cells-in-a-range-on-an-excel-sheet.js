/**
 * @param {string} s
 * @return {string[]}
 */
var cellsInRange = function(s) {
  let ans = []
  const [start, end] = s.split(':')
  for (let i = start.charCodeAt(0); i <= end.charCodeAt(0); i++) {
    for (let j = start[1]; j <= end[1]; j++) {
      ans.push(String.fromCharCode(i) + j)
    }
  }
  return ans
};