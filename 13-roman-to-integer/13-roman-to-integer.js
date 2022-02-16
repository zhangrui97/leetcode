/**
 * @param {string} s
 * @return {number}
 */
var romanToInt = function(s) {
  const ch2n = ch => {
    switch (ch) {
      case 'I': return 1
      case 'V': return 5
      case 'X': return 10
      case 'L': return 50
      case 'C': return 100
      case 'D': return 500
      case 'M': return 1000
    }
  }
  let a = []
  for (let ch of s) {
    a.push(ch2n(ch))
  }
  let sum = 0
  for (let i = 0; i < a.length; i++) {
    if (a[i] < a[i+1]) {
      sum -= a[i]
    } else {
      sum += a[i]
    }
  }
  return sum
};