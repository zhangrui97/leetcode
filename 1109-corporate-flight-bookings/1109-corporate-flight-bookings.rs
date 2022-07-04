impl Solution {
  pub fn corp_flight_bookings(bookings: Vec<Vec<i32>>, n: i32) -> Vec<i32> {
    let mut ans = vec![0; n as usize];
    for r in bookings {
      ans[(r[0]-1) as usize] += r[2];
      if r[1] < n {
        ans[r[1] as usize] -= r[2];
      }
    }
    for i in 1usize..(n as usize) {
      ans[i] += ans[i-1]
    }
    ans
  }
}