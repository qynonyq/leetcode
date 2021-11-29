package problems

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

type input struct {
	desc     string
	nums     []int
	target   int
	expected []int
}

func TestTwoSum(t *testing.T) {
	bigAmount := make([]int, 10001, 10001)
	for i := 1; i < len(bigAmount); {
		bigAmount[i-1] = 2
		bigAmount[i] = 3
		i += 2
	}
	maxNum := int(math.Pow10(10))
	minNum := -int(math.Pow10(10))
	halfBillion := 5 * int(math.Pow10(8))
	halfBillionNeg := -5 * int(math.Pow10(8))
	inputs := []input{
		{
			desc:     "ok",
			nums:     []int{2, 3},
			target:   5,
			expected: []int{0, 1},
		}, {
			desc:     "not found",
			nums:     []int{2, 3},
			target:   6,
			expected: []int{},
		}, {
			desc:     "small amount",
			nums:     []int{2},
			target:   5,
			expected: nil,
		}, {
			desc:     "big amount",
			nums:     bigAmount,
			target:   5,
			expected: nil,
		}, {
			desc:     "max num",
			nums:     []int{0, maxNum},
			target:   5,
			expected: nil,
		}, {
			desc:     "min num",
			nums:     []int{0, minNum},
			target:   5,
			expected: nil,
		}, {
			desc:     "max target",
			nums:     []int{halfBillion, halfBillion},
			target:   maxNum,
			expected: nil,
		}, {
			desc:     "min target",
			nums:     []int{halfBillionNeg, halfBillionNeg},
			target:   minNum,
			expected: nil,
		},
	}

	for _, in := range inputs {
		t.Run(in.desc, func(t *testing.T) {
			result := TwoSum(in.nums, in.target)
			assert.Equal(t, in.expected, result)
		})
	}
}
