func findFinalValue(nums []int, original int) int {
  last := len(nums) - 1
  sort.Ints(nums)

  var search func (l, r, target int) int 
  search = func (l, r, target int) int {
    if l > r {
      return target
    }
    mid := (l + r) / 2
    if nums[mid] < target {
      return search(mid + 1, r, target)
    } else if nums[mid] == target {
      result := target * 2
      if mid >= last {
        return result
      }
      next := search(mid+1, last, result)
      if next == result {
        return result
      } else {
        return next
      }
    } else {
      return search(l, mid-1, target)
    }
  }
  
  return search(0, last, original)
}