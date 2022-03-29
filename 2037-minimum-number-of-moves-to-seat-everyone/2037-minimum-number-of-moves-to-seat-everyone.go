func minMovesToSeat(seats []int, students []int) int {
  sort.Ints(seats)
  sort.Ints(students)
  ans := 0
  for i := 0; i < len(students); i++ {
    if students[i] > seats[i] {
      ans += students[i] - seats[i]
    } else {
      ans += seats[i] - students[i]
    }
  }
  return ans
}