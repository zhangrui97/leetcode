/**
 * @param {number[][]} matrix
 */
var NumMatrix = function(matrix) {
  this.matrix = matrix
  for (j = 1; j < matrix[0].length; j++) {    
    this.matrix[0][j] = this.matrix[0][j] + this.matrix[0][j-1] 
  }

  for (i = 1; i < matrix.length; i++) {
    this.matrix[i][0] = this.matrix[i][0] + this.matrix[i-1][0]
    
    
    for (j = 1; j < matrix[0].length; j++) {
    
    this.matrix[i][j] = this.matrix[i - 1][j] 
      + this.matrix[i][j-1] 
      - this.matrix[i-1][j-1] 
      + this.matrix[i][j]
      
    }
  }
  
    
    
    
  /*[[1,1,1,1,1],
     [2,2,2,2,2],
     [3,123,123,12,3],
     [4,123,12312,12,1]] */

};

/** 
 * @param {number} row1 
 * @param {number} col1 
 * @param {number} row2 
 * @param {number} col2
 * @return {number}
 */
NumMatrix.prototype.sumRegion = function(row1, col1, row2, col2) {
  if (row1 === 0 && col1 === 0) { 
    
    return(this.matrix[row2][col2])
    
  }
  if (row1 === 0) {
    return(this.matrix[row2][col2] - this.matrix[row2][col1-1])
    
  }
  if (col1 === 0) {
    
     return(this.matrix[row2][col2] - this.matrix[row1-1][col2])
  }
  return (this.matrix[row2][col2] - this.matrix[row1-1][col2] - this.matrix[row2][col1-1] + this.matrix[row1-1][col1-1])
};

/** 
 * Your NumMatrix object will be instantiated and called as such:
 * var obj = new NumMatrix(matrix)
 * var param_1 = obj.sumRegion(row1,col1,row2,col2)
 */