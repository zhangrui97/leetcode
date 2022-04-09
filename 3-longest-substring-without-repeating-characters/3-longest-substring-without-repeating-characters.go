func lengthOfLongestSubstring(s string) int {
  seen := [127]int{}
  for i := 0; i < 127; i++ { seen[i] = -1 }
  l := 0
  ans := 0
  for i, v := range s {
    if val := seen[int(v)]; val >= l {
      l = val + 1
    } else {
      if i + 1 - l > ans {
        ans = i + 1 - l
      }
    }
    seen[int(v)] = i
  }
  return ans
}