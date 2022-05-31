package main

import "fmt"

func main() {

    nums1 := []int{4,9,5}
    nums2 := []int{9,4,9,8,4}

    fmt.Printf("Res: %+v\n", intersection(nums1, nums2))

    a := []int{1, 5, 10, 20, 40, 80}
    b := []int{6, 7, 20, 80, 100}
    c := []int{3, 4, 15, 20, 30, 70, 80, 120}
    fmt.Printf("Res: %+v\n", intersectionThreeSortedArrays(a, b, c))
}

func intersection(nums1 []int, nums2 []int) []int {
    m := make(map[int]struct{}, 0)

    for _, num := range nums1 {
        m[num] = struct{}{}
    }

    res := make([]int, 0)
    for _, num := range nums2 {
        if _, exists := m[num]; exists {
            res = append(res, num)
            delete(m, num) 
        }
    }

    return res
}

func intersectionThreeSortedArrays(nums1 []int, nums2 []int, nums3 []int) []int {
   
   i, j, k := 0, 0, 0
   res := make([]int, 0)
   for i < len(nums1) && j < len(nums2) && k < len(nums3) {
       if nums1[i] == nums2[j] && nums3[k] == nums2[j] {
           res = append(res, nums1[i])
           i++
           j++
           k++
       } else {
           if nums1[i] < nums2[j] && nums1[i] < nums3[k] {
               i++
           } else if nums2[j] < nums3[k] && nums2[j] < nums1[i] {
               j++
           } else {
               k++
           }
       }
   }

   return res
}