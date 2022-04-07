type NumMatrix struct {
  sums [][]int  
}


func Constructor(matrix [][]int) NumMatrix {
  m, n := len(matrix), len(matrix[0])
  sums := make([][]int, m+1)
  sums[0] = make([]int, n+1)
  for i, r := range matrix {
    row := make([]int, n+1)
    sums[i+1] = row
    for j, v := range r {
      row[j+1] = v + sums[i][j+1] + row[j] - sums[i][j]
    }
  }
  return NumMatrix{sums}  
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
  return this.sums[row2+1][col2+1] - this.sums[row2+1][col1] - this.sums[row1][col2+1] + this.sums[row1][col1]
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */