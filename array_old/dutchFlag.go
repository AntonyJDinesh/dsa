package main

import "fmt"

func main() {
    nums := []int{2,0,2,1,1,0}
    sortColors(nums)
    fmt.Printf("%+v", nums)
}

func sortColors(nums []int)  {

    zero, one, two := 0, 0, len(nums)-1

    for one <= two {
        switch nums[one] {
            case 0:
                nums[zero], nums[one] = nums[one], nums[zero]
                zero++
                one++
            case 1:
                one++
            case 2:
                nums[two], nums[one] = nums[one], nums[two]
                two--
        }
    }
}