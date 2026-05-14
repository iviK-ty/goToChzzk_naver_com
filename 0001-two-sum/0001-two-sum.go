func twoSum(nums []int, target int) []int {
    indexMap := make(map[int]int)  

    for i , num := range nums {  
        pair := target - num

        if j, ok := indexMap[pair]; ok{
            return []int{j, i}
        }
        indexMap[num] = i
    }    
    return nil
}