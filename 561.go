func arrayPairSum(nums []int) int {
    insertSort(nums)
    
    result := 0
    
    for index,item := range nums {
        if index % 2 == 0 {
            result += item
        }
    }
    
    return result
}

func insertionSort(nums []int) {
    length := len(nums)
    
    for i := 1; i < length; i++ {
        temp := nums[i]
        
        j := i-1
        
        for ;j>=0 && nums[j] > temp; j-- {
            nums[j+1] = nums[j]
        }
        
        nums[j+1] = temp
    }
}