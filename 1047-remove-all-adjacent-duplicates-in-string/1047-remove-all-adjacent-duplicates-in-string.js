/**
 * @param {string} s
 * @return {string}
 */
var removeDuplicates = function(s) {
  const stack = []
  for (let ch of s) {
    if (stack.at(-1) === ch) stack.pop()
    else stack.push(ch)
  }
  return stack.join('')
};