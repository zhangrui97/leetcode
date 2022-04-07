type NumMatrix struct {
  sums [][]int  
}


func Constructor(matrix [][]int) NumMatrix {
  for i, r := range matrix {
    for j, v := range r {
      if i == 0 {
        if j == 0 {
          matrix[0][0] = v
        } else {
          matrix[i][j] = v + matrix[i][j-1]
        }
      } else if j == 0 {
        matrix[i][0] = v + matrix[i-1][0]
      } else {
        matrix[i][j] = v + matrix[i-1][j] + matrix[i][j-1] - matrix[i-1][j-1]
      }
    }
  }
  return NumMatrix{matrix}  
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
  result := this.sums[row2][col2]
  if row1 > 0 {
    result -= this.sums[row1-1][col2]
  } 
  if col1 > 0 {
    result -= this.sums[row2][col1-1]
  }
  if row1 > 0 && col1 > 0 {
    result += this.sums[row1-1][col1-1]
  }
  return result
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */