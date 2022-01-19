# Tricks

### 交換傳入參數位置
由於函數傳入值不一定是前面的長度比較短或比較小，你可能想要確保前面的比較短，可以用遞迴的方式確保位置，便於後面實作。

```golang
func intersection(nums1 []int, nums2 []int) []int {
    if len(nums1) > len(nums2)
        return intersection(nums2, nums1)
    ...
    ...
    ...
}
```
