impl Solution {
  pub fn can_partition(nums: Vec<i32>) -> bool {
    let (max, sum) = (nums.iter().max().unwrap(), nums.iter().sum());
    if sum % 2 != 0 || max+max > sum { return false }
    let target = (sum / 2) as usize;
    let mut dp = vec![false; target + 1];
    dp[0usize] = true;
    for &n in nums.iter() {
      for i in (n as usize..=target).rev() {
        dp[i] = dp[i] || dp[i-(n as usize)]
      }
    }
    dp[target]
  }
}