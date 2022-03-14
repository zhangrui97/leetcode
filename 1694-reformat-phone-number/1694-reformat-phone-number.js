/**
 * @param {string} number
 * @return {string}
 */
var reformatNumber = function(number) {
  const a = number.split('').filter(ch => !isNaN(ch) && ch !== ' ')
  let i = 0
  let ans = ''
  const l = a.length
  for (; i < l - 4; i+=3) {
    ans += a[i] + a[i+1] + a[i+2] + '-'
  }
  switch (l - i) {
    case 2: return ans + a[l-2] + a[l-1]
    case 3: return ans + a[l-3] + a[l-2] + a[l-1]
    case 4: return ans + a[l-4] + a[l-3] + '-' + a[l-2] + a[l-1]
  }
  return ans
};