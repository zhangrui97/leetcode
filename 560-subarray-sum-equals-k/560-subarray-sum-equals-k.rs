use std::collections::HashMap;

impl Solution {
  pub fn subarray_sum(nums: Vec<i32>, k: i32) -> i32 {
    let mut dict = HashMap::with_capacity(nums.len()+1);
    dict.insert(0, 1);
    let mut sum = 0;
    let mut counter = 0;
    for v in nums {
      sum += v;
      if let Some(c) = dict.get(&(sum - k)) {
        counter += c;
      }
      let e = dict.entry(sum).or_insert(0);
      *e += 1;
    }
    counter
  }
}