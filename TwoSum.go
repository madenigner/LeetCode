func twoSum(nums []int, target int) []int {
    res := []int{}
  
    for i := 0 ; i <= len(nums)-1 ; i++ {
        for j := i+1 ; j <= len(nums)-1 ; j++ {
            if nums[i] + nums[j] == target {
                res = append(res , i,j)
            }
        }
    }
    return res

}

