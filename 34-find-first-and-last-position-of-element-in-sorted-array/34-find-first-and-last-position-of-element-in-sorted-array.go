func search(nums []int, target int) (int, int, int) {
  l, r := 0, len(nums) - 1
  for l <= r {
    mid := (l + r) / 2
    if nums[mid] == target {
      return l, mid, r
    } else if nums[mid] < target {
      l = mid + 1
    } else {
      r = mid - 1
    }
  }
  return 0, -1, 0
}

func searchLeft(nums []int, target, l, r int) int {
  for l < r {
    mid := (l + r) / 2
    if nums[mid] < target {
      l = mid + 1
    } else {
      r = mid
    }
  }
  return l
}

func searchRight(nums []int, target, l, r int) int {
  for l < r {
    mid := (l + r) / 2
    if nums[mid] > target {
      r = mid
    } else {
      l = mid + 1
    }
  }
  return l - 1
}

func searchRange(nums []int, target int) []int {
  l, idx, r := search(nums, target)
  if idx == -1 {
    return []int{-1, -1}
  } else {
    return []int{searchLeft(nums, target, l, idx),
                searchRight(nums, target, idx, r+1)}
  }
}