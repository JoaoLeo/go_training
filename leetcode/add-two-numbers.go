/*
	You are given two non-empty linked lists representing two non-negative integers.

The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example 1:

Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
*/
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, n := range nums[1:] {
		current.Next = &ListNode{Val: n}
		current = current.Next
	}
	return head
}

func printList(l *ListNode) {
	fmt.Print("[")
	for l != nil {
		fmt.Print(l.Val)
		if l.Next != nil {
			fmt.Print(",")
		}
		l = l.Next
	}
	fmt.Println("]")
}

func reverseList(l *ListNode) *ListNode {
	var list []int

	for l != nil {
		if l.Next != nil {
			list = append(list, l.Val)
		}
		l = l.Next
	}

	result := buildList(list)
	return result
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return l1
}

func main() {
	l1 := buildList([]int{2, 4, 3})
	l2 := buildList([]int{5, 6, 4})
	printList(reverseList(l1))
	printList(reverseList(l2))
}
