/**
 * @param {number[]} temperatures
 * @return {number[]}
 */
var dailyTemperatures = function(temperatures) {
  const ans = Array(temperatures.length).fill(0)
  const stack = []
  temperatures.forEach((v, i) => {
    while (stack.length > 0 && v > temperatures[stack.at(-1)]) {
      const j = stack.pop()
      ans[j] = i - j
    }
    stack.push(i)
  })
  return ans
};