func needMoreDays(weights []int, days, cap int) bool {
  lastWeight := 0
  l, r := 0, len(weights)
  for days > 0 {
    for l < r {
      m := (l + r)/2
      if weights[m] - lastWeight < cap {
        l = m + 1
      } else {
        r = m
      }
    }
    if l >= len(weights) {
      return false
    }
    if weights[l] - lastWeight == cap {
      l++
      if l == len(weights) {
        return false
      }
    }
    days--
    lastWeight = weights[l-1]
    r = len(weights)
  }
  return true
}

func shipWithinDays(weights []int, days int) int {
  leng := len(weights)
  for i := 1; i < leng; i++ {
    weights[i] += weights[i-1]
  }  
  max := weights[leng-1]
  if days == 1 { return max }
  l, r := max / days, max + 1
  if l < weights[0] {
    l = weights[0]
  }
  for l < r {
    m := (l + r) / 2
    if needMoreDays(weights, days, m) {
      l = m + 1
    } else {
      r = m
    }
  }
  return l
}