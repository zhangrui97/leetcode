func moveZeroes(nums []int)  {
  s := 0
  for _, v := range nums {
    if v != 0 {
      nums[s] = v
      s++
    }
  }
  for ; s < len(nums); s++ {
    nums[s] = 0
  }
}