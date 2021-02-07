package main

import (
	"log"
	"math"
	"sort"
)

func main() {
	log.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}

func threeSumClosest(nums []int, target int) (nearestNumber int) {
	if len(nums) < 2 {
		return
	}
	smallestRange := math.MaxInt32
	sort.Ints((nums))

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] >= target && target > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				sumNums := nums[i] + nums[j] + nums[k]
				deviationNumber := sumNums - target
				if deviationNumber < 0 {
					deviationNumber *= -1
				}
				if deviationNumber == 0 {
					nearestNumber = sumNums
					return
				} else if deviationNumber < smallestRange {
					smallestRange = deviationNumber
					nearestNumber = sumNums
				}
			}
		}
	}
	return
}
