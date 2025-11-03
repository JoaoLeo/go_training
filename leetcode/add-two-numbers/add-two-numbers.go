/*
You are given two non-empty linked lists representing two non-negative integers.

The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example 1:

Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

Dividir para conquistas
1 - Criar os listnode baseado nas lista - feito
2 - Reverter a ordem dos listnode - feito
3 - Transformar os listnode em numeros inteiros - feito
4 - Somar numeros - feito
5 - Transformar o numero em um listnode
*/
package main

import (
	"fmt"
	"strconv"
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
	current := l
	cont := 1

	// contando quantidade de itens no listnode
	for current.Next != nil {
		current = current.Next
		cont++
	}

	// criando lista do tamanho do listnode e adicionando na lista o listnode ao contrario
	list := make([]int, cont)
	current2 := l
	cont2 := len(list) - 1
	for current2 != nil {
		list[cont2] = current2.Val
		current2 = current2.Next
		cont2--
	}

	result := buildList(list)
	return result
}

func convertListNodeToInt(l *ListNode) int {
	var numeroStr string

	current := l
	for current != nil {
		numeroStr += strconv.Itoa(current.Val)
		current = current.Next
	}

	numero, err := strconv.Atoi(numeroStr)

	if err != nil {
		return 0
	}

	return numero
}

func createListNodeFromInt(num int) *ListNode {
	fmt.Println(num)
	numeroStr := strconv.Itoa(num)
	list := make([]int, len(numeroStr))

	for i, c := range numeroStr {
		numero, err := strconv.Atoi(string(c))
		if err == nil {
			list[i] = numero
		}
	}

	result := buildList(list)

	return result
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverseList(l1)
	l2 = reverseList(l2)
	n1 := convertListNodeToInt(l1)
	n2 := convertListNodeToInt(l2)
	resultInt := n1 + n2
	result := reverseList(createListNodeFromInt(resultInt))
	return result
}

func main() {
	l1 := buildList([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
	l2 := buildList([]int{5, 6, 4})
	printList(addTwoNumbers(l1, l2))
}
