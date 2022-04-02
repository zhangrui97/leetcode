func isEvenDigits(num int) bool {
  if num > 999 {
    return num <= 9999 || num == 100000
  } else {
    return num > 9 && num <= 99
  }
}

func findNumbers(nums []int) int {
  ans := 0
  for _, v := range nums {
    if isEvenDigits(v) {
      ans ++
    }
  }
  return ans
}