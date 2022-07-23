use std::collections::HashMap;

impl Solution {
  pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
    let mut map = HashMap::new();
    for (i, n) in nums.iter().enumerate() {
      match map.get(n) {
        Some(&v) => return vec![i as i32, v],
        _ => map.insert(target - n, i as i32),
      };
    }
    vec![]
  }
}