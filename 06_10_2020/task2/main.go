package main

import (
	"fmt"
)

func main() {
	var nums [3]int;

	for i := 0; i < len(nums); i++ {
		fmt.Print("Enter ", i ," num ")
		fmt.Scan(&nums[i])
	}

	if nums[0] == nums[1] || nums[1] == nums[2] || nums[0] == nums[2] {
		fmt.Println("nice joke")
	}

	var min = nums[0]
    for _, v := range nums {
            if (v < min) {
                min = v
            }
	}
	
	var max = nums[0]
    for _, v := range nums {
            if (v > max) {
                max = v
            }
	}

	for _, v := range nums {
		if (v != min && v != max) {
			fmt.Println(v)
		}
	}
}
