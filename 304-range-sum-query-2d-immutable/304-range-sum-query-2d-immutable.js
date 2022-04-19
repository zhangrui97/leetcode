function buildPrefixSums(m) {
  const fstRow = new Array(m[0].length+1)
  fstRow.fill(0)
  const result = [fstRow]
  m.forEach((row, i) => {
    const newRow = [0]
    row.forEach((v, j) => {
      newRow.push(v + result[i][j+1] + newRow[j] - result[i][j])
    })
    result.push(newRow)
  })
  return result
}
/**
 * @param {number[][]} matrix
 */
var NumMatrix = function(matrix) {
  this.matrix = buildPrefixSums(matrix)
}
  
/** 
 * @param {number} row1 
 * @param {number} col1 
 * @param {number} row2 
 * @param {number} col2
 * @return {number}
 */
NumMatrix.prototype.sumRegion = function(row1, col1, row2, col2) {
  return this.matrix[row2+1][col2+1] - this.matrix[row1][col2+1] - this.matrix[row2+1][col1] + this.matrix[row1][col1]
};

/** 
 * Your NumMatrix object will be instantiated and called as such:
 * var obj = new NumMatrix(matrix)
 * var param_1 = obj.sumRegion(row1,col1,row2,col2)
 */