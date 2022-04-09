func findAnagrams(s string, p string) []int {
  sl, pl := len(s), len(p)
  if sl < pl { return []int{} }
  target := [26]int{}
  for _, v := range p { target[v-'a']++ }
  window := [26]int{}
  for i := 0; i < pl; i++ { window[s[i]-'a']++ }
  ans := make([]int, 0, sl /pl + 1)
  if window == target { ans = append(ans, 0) }
  for l, r := 0, pl; r < sl; r++{
    window[s[r]-'a']++
    window[s[l]-'a']--
    l++ 
    if window == target { ans = append(ans, l) }
  }
  return ans
}