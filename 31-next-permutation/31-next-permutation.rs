impl Solution {
  pub fn next_permutation(nums: &mut Vec<i32>) {
    let n:i32 = nums.len() as i32;
    let mut k:i32 = n - 2;
    while k >= 0 && nums[k as usize] >= nums[(k+1) as usize] {
      k -= 1;
    }
    if k < 0 {
      nums.reverse();
      return
    }
    let mut l = n - 1;
    let v = nums[k as usize];
    while l > k && nums[l as usize] <= v {
      l -= 1;
    }
    nums[k as usize] = nums[l as usize];
    nums[l as usize] = v;
    let mut i = k+1;
    let mut j = n-1;
    while i < j {
      nums.swap(i as usize, j as usize);
      i += 1;
      j -= 1;
    }
  }
}