struct NumArray {
  sum: Vec<i32>
}


/** 
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl NumArray {

    fn new(nums: Vec<i32>) -> Self {
      let mut sum = vec![0; nums.len() + 1];
      for (i, &num) in nums.iter().enumerate() {
        sum[i+1] = sum[i] + num;
      }
      NumArray {
        sum
      }
    }
    
    fn sum_range(&self, left: i32, right: i32) -> i32 {
      self.sum[(right+1) as usize] - self.sum[left as usize]
    }
}

/**
 * Your NumArray object will be instantiated and called as such:
 * let obj = NumArray::new(nums);
 * let ret_1: i32 = obj.sum_range(left, right);
 */