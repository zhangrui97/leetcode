func needDays(weights []int, cap int) int {
  res := 0
  lastWeight := 0
  for i, v := range weights {
    if v > lastWeight + cap {
      res++
      lastWeight = weights[i-1]
    } else if v == lastWeight + cap {
      res++
      lastWeight = weights[i]
    }
  }
  if lastWeight == weights[len(weights)-1] {
    return res
  } else {
    return res + 1
  }
}

func shipWithinDays(weights []int, days int) int {
  leng := len(weights)
  max := weights[0]
  for i := 1; i < leng; i++ {
    v := weights[i]
    if v > max {
      max = v
    }
    weights[i] = weights[i-1] + v
  }  
  sum := weights[leng-1]
  if days == 1 { return sum }
  l, r := max, sum + 1
  for l < r {
    m := (l + r) / 2
    if needDays(weights, m) > days {
      l = m + 1
    } else {
      r = m
    }
  }
  return l
}