package medium

import (
	"fmt"
	"math/big"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(numbers []int) []int {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func getNum(l *ListNode) *big.Int {
	var nums []int
	for {
		if l.Next == nil {
			nums = append(nums, l.Val)
			break
		}
		nums = append(nums, l.Val)
		l = l.Next
	}
	nums = reverse(nums)
	return numsToNum(nums)
}

func numsToNum(nums []int) *big.Int {
	num := big.NewInt(0)
	// for i, n := range nums {
	// 	num.Add()n * 10 * i
	// }
	// num.SetString()
	numStr := strings.Trim(strings.Replace(fmt.Sprint(nums), " ", "", -1), "[]")
	num.SetString(numStr, 10)
	// num, err := strconv.ParseUint(numStr, 10, 64)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	return num
}

func intToSlice(n *big.Int, sequence []int) []int {
	zeroInt := big.NewInt(0)
	tenInt := big.NewInt(10)
	if n.Cmp(zeroInt) != 0 {
		i := n.Mod(n, tenInt)
		sequence = append(sequence, int(i.Int64()))
		return intToSlice(n.Div(n, tenInt), sequence)
	}
	return sequence
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1.Val == 0 && l1.Next == nil && l2.Val == 0 && l2.Next == nil {
		return &ListNode{Val: 0}
	}
	n1 := getNum(l1)
	n2 := getNum(l2)
	sum := big.NewInt(0)
	sum.Add(n1, n2)
	nums := intToSlice(sum, []int{})
	list := make([]*ListNode, len(nums), len(nums))
	for i, n := range nums {
		list[i] = &ListNode{Val: n}
	}
	for i := 0; i < len(nums)-1; i++ {
		list[i].Next = list[i+1]
	}
	return list[0]
}
