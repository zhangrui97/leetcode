impl Solution {
  pub fn corp_flight_bookings(bookings: Vec<Vec<i32>>, n: i32) -> Vec<i32> {
    let mut ans = vec![0; n as usize];
    for r in bookings {
      ans[(r[0]-1) as usize] += r[2];
      if r[1] < n {
        ans[r[1] as usize] -= r[2];
      }
    }
    println!("{:?}", ans);
    let mut s = 0;
    for n in &mut ans {
      s += *n;
      *n = s;
    }
    ans
  }
}