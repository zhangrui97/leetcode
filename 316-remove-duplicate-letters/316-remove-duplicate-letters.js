/**
 * @param {string} s
 * @return {string}
 */
var removeDuplicateLetters = function(s) {
  const counter = Array(127).fill(0)
  for (const ch of s) {
    counter[ch.charCodeAt(0)]++
  }
  const inStack = Array(127).fill(false)
  const stack = []
  for (const ch of s) {
    const code = ch.charCodeAt(0)
    counter[code]--
    if (inStack[code]) continue
    while (stack.length > 0 && counter[stack.at(-1)] > 0 && code < stack.at(-1)) {
      const last = stack.pop()
      inStack[last] = false
    }
    stack.push(code)
    inStack[code] = true
  }
  return String.fromCharCode(...stack)
};