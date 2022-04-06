type NumArray struct {
  sums []int  
}


func Constructor(nums []int) NumArray {
  result := make([]int, len(nums)+1)
  result[0] = 0
  sum := 0
  for i, v := range nums {
    sum += v
    result[i+1] = sum
  }
  return NumArray{result}
}


func (this *NumArray) SumRange(left int, right int) int {
  return this.sums[right+1] - this.sums[left]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */