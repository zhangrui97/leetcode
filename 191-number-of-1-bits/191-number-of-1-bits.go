func hammingWeight(num uint32) int {
  counter := 0
  for num != 0 {
    num &= num-1
    counter++
  }
  return counter
}