impl Solution {
  pub fn car_pooling(trips: Vec<Vec<i32>>, capacity: i32) -> bool {
    let mut diff = [0; 1001];
    let mut max = 0usize;
    for trip in trips {
      let (num, from, to) = (trip[0], trip[1] as usize, trip[2] as usize);
      if num > capacity { return false }
      if to > max { max = to }      
      diff[from] += num;
      diff[to] -= num;
    }
    let mut sum = 0;
    for i in 0usize..max {
      sum += diff[i];
      if sum > capacity {
        return false
      }
    }
    return true
  }
}