/**
 * @param {number} n
 * @return {number}
 */
var tribonacci = function(n) {
  switch (n){
    case 0: return 0
    case 1: return 1
    case 2: return 1
    default:
      let t0 = 0
      let t1 = 1
      let t2 = 1
      for (let i = 3; i <= n; i++) {
        [t0, t1, t2] = [t1, t2, t0 + t1 + t2]
      }
      return t2
  }
};