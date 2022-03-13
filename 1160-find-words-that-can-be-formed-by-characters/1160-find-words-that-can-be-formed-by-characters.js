/**
 * @param {string[]} words
 * @param {string} chars
 * @return {number}
 */
var countCharacters = function(words, chars) {
  const resultWords = words.filter(word => {
    let tester = chars.slice()
    for (let ch of word) {
      if (tester.includes(ch)) {
        tester = tester.replace(ch, '', 1)
      } else return false      
    }
    return true
  })
  console.log(resultWords)
  let result = 0
  for (word of resultWords)
    result += word.length
  return result
};