/**
 * @param {string} word
 * @return {number}
 */
var countVowelSubstrings = function(word) {
  let ans = 0
  for (let left = 0; left < word.length - 4; left++) {
    if ('aeiou'.includes(word[left])) {
      let right = left + 1
      while (right < word.length && 'aeiou'.includes(word[right])) {
        if (right >= left + 4 && (new Set([...word.substring(left, right+1)]).size) === 5)
          ans ++
        right++
      }
      if (right <= left + 5) left = right
    }
  }
  return ans
};