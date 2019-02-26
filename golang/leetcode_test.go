package golang

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	if sum := twoSum([]int{2, 7, 11, 15}, 9); sum == nil || len(sum) != 2 {
		t.Fail()
	}
}
