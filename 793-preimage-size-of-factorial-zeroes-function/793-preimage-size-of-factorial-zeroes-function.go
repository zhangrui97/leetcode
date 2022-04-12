func trailingZeros(n int64) int {
  result := int64(0)
  for n >= 5 {
    n /= 5
    result += n
  }
  return int(result)
}

func searchUpper(k int, l, r int64) int64 {
  if l == r { return l }
  m := (l + r) / 2
  if trailingZeros(m) < k {
    l = m + 1
  } else {
    r = m
  }
  return searchUpper(k, l, r)
}

func preimageSizeFZF(k int) int {
  if k == 0 { return 5 }
  if k == 1 { return 5 }
  lo := searchUpper(k, int64(0), int64(5*k))      
  hi := searchUpper(k+1, lo, int64(5*k+5))
  return int(hi - lo)
}