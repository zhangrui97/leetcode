func maxProduct(nums []int) int {
  var m1, m2 int
  for _, v := range nums {
    if v > m1 {
      m1, m2 = v, m1
    } else if v > m2 {
      m2 = v
    }
  }  
  return (m1-1)*(m2-1)
}