func sumRest(piles []int, start, n int) int {
  sum := 0
  for i := start; i < len(piles); i++ {
    m := piles[i]
    if m % n == 0 {
      sum += m / n
    } else {
      sum += m / n + 1
    }
  }
  return sum
}

func minEatingSpeed(piles []int, h int) int {
  pl := len(piles)
  if pl == h {
    max := piles[0]
    for i := 1; i < pl; i++  {
      if piles[i] > max {
        max = piles[i]
      }
    }
    return max
  }
  if pl == 1 {
    if piles[0] % h == 0 { 
      return piles[0] / h
    } else { 
      return piles[0] / h + 1 
    }
  }
  sort.Ints(piles)
  cache := make(map[int]int)
  l, r := 0, pl
  if h < 2 * pl {
    l = 2*pl - h
  }
  for l < r {
    m := l + (r-l) / 2
    n := piles[m]
    sum := m + 1
    if val, ok := cache[n]; ok {
      sum = val
    } else {
      sum += sumRest(piles, m+1, n)
      cache[n] = sum
    }
    if sum < h {
      r = m
    } else {
      l = m + 1
    }
  }
  min, max := 0, piles[l] + 1
  if l > 0 {
    min = piles[l-1]
  }
  for min < max {
    m := min + (max - min) / 2
    sum := l + sumRest(piles, l, m)
    if sum > h {
      min = m + 1
    } else {
      max = m
    }
  }
  return min
}