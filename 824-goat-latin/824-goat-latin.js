/**
 * @param {string} sentence
 * @return {string}
 */
var toGoatLatin = function(sentence) {
  const isVowel = ch => 'aeiouAEIOU'.includes(ch)
  return sentence.split(' ').map((word, i) => {
    if (isVowel(word[0])) return word + 'ma' + Array(i+1).fill('a').join('')
    else return word.substring(1) + word[0] + 'ma' + Array(i+1).fill('a').join('')
  }).join(' ')  
};