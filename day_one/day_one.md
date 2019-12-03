## leetcode01  
Given an array of integers, return indices of the two numbers such that they add up to a specific target.  

You may assume that each input would have exactly one solution, and you may not use the same element  
twice.

Example:

```
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```  
## code
```
func twoSum(nums []int, target int) []int {
    
}
```
这道题主要考的是数组以及循环的使用  
第一种解法，暴力法使用两个for循环来解决问题，但是这个算法在leetcode的判题中会出现错误，没想明白  

```
func twoSum(nums []int, target int) []int {
    res := []int{}
    for k1,v1 := range nums{
        for k2,v2 := range nums{
            if((v1 + v2) == target){
            res = []int{k1,k2}
            breake
            }
        }
    }
    return res
}
```

第二种解法

```
func twoSum(nums []int, target int) []int {
    m := make(map[int]int, len(nums))
    res := []int{}
    for k,v := range nums {
        if _,val2 := m[target-v];val2 {
            res = []int{k,m[target-v]}
            break
        }
        m[v] = k
    }
    return res
}
```


## leetcode02
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in  
reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a  
linked list.  

You may assume the two numbers do not contain any leading zero, except the number 0 itself.
```
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
```
## code
```
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    
}
```