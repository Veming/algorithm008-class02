func majorityElement(nums []int) int {
    count:= 0
    var major int
    for _, n := range nums {
        if count == 0 {
            major = n
        }
        if major == n {
            count +=1
        } else {
            count -=1
        }
    }
    return major
}
