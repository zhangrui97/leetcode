/**
 * @param {number} num
 * @return {string}
 */
var intToRoman = function(num) {
    const rn = "IVXLCDM"
    const convert = (n, level) => {
      const next = ~~(n / 10)
      const result = next === 0 ? "" : convert(next, level + 1)
      const one = rn[2*level]
      const five = level < 3 && rn[2*level+1]
      const ten = level < 3 && rn[2*level+2]
      switch (n % 10) {
        case 0: return result
        case 1: return result + one
        case 2: return result + one + one
        case 3: return result + one + one + one
        case 4: return result + one + five
        case 5: return result + five
        case 6: return result + five + one
        case 7: return result + five + one + one
        case 8: return result + five + one + one + one
        case 9: return result + one + ten
      }
    }
    return convert(num, 0)
};