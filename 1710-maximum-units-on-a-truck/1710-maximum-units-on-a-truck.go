func maximumUnits(boxTypes [][]int, truckSize int) int {
  sort.Slice(boxTypes, func(i, j int) bool {
    return boxTypes[i][1] > boxTypes[j][1]})
  total := 0
  for _, v := range boxTypes {
    if v[0] < truckSize {
      total += v[0] * v[1]
      truckSize -= v[0]
    } else {
      total += truckSize * v[1]
      break
    }
  }
  return total
}