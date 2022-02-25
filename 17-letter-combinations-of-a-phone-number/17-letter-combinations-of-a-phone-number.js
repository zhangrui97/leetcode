/**
 * @param {string} digits
 * @return {string[]}
 */
var letterCombinations = function(digits) {
  const map = ["","","abc","def","ghi","jkl","mno","pqrs","tuv","wxyz"]
  let result = []
  for (let d of digits) {
    if (result.length === 0) {
      for (let ch of map[+d]) result.push(ch)
    } else {
      const temp = []
      for (let s of result) {
        for (let ch of map[+d]) temp.push(s+ch)
      }
      result = temp
    }
  }
  return result
};