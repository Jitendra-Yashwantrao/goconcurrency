package main

import (
	"fmt"
	"goconcurrency/welcome"
	"strings"

	"github.com/pborman/uuid"
)

func main() {
	uuidWithHyphen := uuid.NewRandom()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	fmt.Println(uuidWithHyphen, "  ", uuid)
	fmt.Println(welcome.Welcome())
	inputNumbers := []int{12, 4, 5, 6, 1, 7}
	fmt.Println(twoSum(inputNumbers, 7))

}

func twoSum(nums []int, target int) []int {

	var answers []int
	numbersMap := make(map[int]int)
	for j := 0; j < len(nums); j++ {
		numbersMap[nums[j]] = j

	}

	//finding number with matching sum
	for j := 0; j < len(nums); j++ {
		nextMatch := target - nums[j]

		value, present := numbersMap[nextMatch]

		if present && value != j {
			answers = append(answers, j, value)
			break
		}

	}

	return answers
}
