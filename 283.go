func moveZeroes(nums []int)  {
    length := len(nums)
    
    if length <= 1 {
        return
    }
    
    p := 1
    if nums[0] == 0 {
        p = 0
    }
    
    for i := 1; i<length; i++ {
        if nums[i] != 0 {      
            nums[p] = nums[i] 
            p++
        }
    }
    
    for ; p<length; p++ {
        nums[p] = 0
    }
}