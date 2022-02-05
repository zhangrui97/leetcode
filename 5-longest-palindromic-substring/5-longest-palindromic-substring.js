/**
 * @param {string} s
 * @return {string}
 */
var longestPalindrome = function(s) {
  const str = ['^', ...([].concat(...([...s].map(ch => ['#', ch])))), '#', '$']
  const len = str.length
  const dp = new Array(len).fill(1)
  let c = 0
  let r = 0
  let index = 1
  for (let i = 1; i < len - 1; i++) {
    if (i < (r - 1)) {
      dp[i] = Math.min(dp[2*c-i], r-i)
    }
    while (str[i-dp[i]] === str[i+dp[i]]) {
      dp[i] += 1
    }
    if ((i + dp[i]) > r) {
      c = i
      r = c + dp[i]
    }
    if (dp[i] > dp[index]) {
      index = i
    }
  }
  return s.substr(Math.floor((index-dp[index])/2), dp[index]-1)
};