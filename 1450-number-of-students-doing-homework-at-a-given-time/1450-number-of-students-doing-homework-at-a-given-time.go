func busyStudent(startTime []int, endTime []int, queryTime int) int {
  ans := 0
  for i := 0; i < len(startTime); i ++ {
    if queryTime >= startTime[i] && queryTime <= endTime[i] {
      ans ++
    }
  }
  return ans
}