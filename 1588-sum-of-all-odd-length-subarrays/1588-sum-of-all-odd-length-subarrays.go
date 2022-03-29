func sumOddLengthSubarrays(arr []int) int {
  if len(arr) == 1 {
    return arr[0]
  }
  pre := arr[0]
  ans := pre
  for i := 2; i < len(arr); i += 2 {
    pre += arr[i-1] + arr[i]
    ans += pre
  }
  return ans + sumOddLengthSubarrays(arr[1:])
}