func peakIndexInMountainArray(arr []int) int {
  l, r := 1, len(arr) - 2
  for l <= r {
    mid := (l + r) / 2
    if arr[mid-1] < arr[mid] && arr[mid+1] < arr[mid] {
      return mid
    } else if arr[mid-1] < arr[mid] {
      l = mid + 1
    } else {
      r = mid - 1
    }
  }
  return -1
}