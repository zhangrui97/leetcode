func getRow(rowIndex int) []int {  
  ans := make([]int, rowIndex + 1)
  ans[0] = 1
  ans[rowIndex] = 1
  for i := 1; i <= rowIndex / 2; i++ {
    ans[i] = ans[i-1] * (rowIndex + 1 - i) / i
    ans[rowIndex - i] = ans[i]
  }
  return ans
}