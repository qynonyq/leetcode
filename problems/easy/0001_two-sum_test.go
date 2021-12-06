package problems

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := map[string]struct {
		nums     []int
		target   int
		expected []int
	}{
		"ok":        {nums: []int{2, 11, 7, 15}, target: 9, expected: []int{0, 2}},
		"not found": {nums: []int{2, 11, 7, 15}, target: 5, expected: nil},
		"one num":   {nums: []int{2}, target: 9, expected: nil},
		"no num":    {nums: []int{}, target: 9, expected: nil},
	}

	for name, in := range tests {
		t.Run(name, func(t *testing.T) {
			result := TwoSum(in.nums, in.target)
			if !reflect.DeepEqual(in.expected, result) {
				t.Errorf("expected %#v, got %#v", in.expected, result)
			}
		})
	}
}
