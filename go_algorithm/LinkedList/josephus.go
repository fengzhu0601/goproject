package main

import "fmt"

// 节点结构
type Node struct {
	Value int
	Next  *Node
}

// 创建循环链表
func CreateCircularLinkedList(n int) *Node {
	if n <= 0 {
		return nil
	}

	// 创建头节点
	head := &Node{
		Value: 1,
	}
	prev := head
	for i := 2; i <= n; i++ {
		node := &Node{
			Value: i,
		}
		prev.Next = node
		prev = node
	}
	prev.Next = head // 将尾节点的Next指向头节点形成循环

	return head
}

// 打印链表
func PrintCircularLinkedList(head *Node) {
	if head == nil {
		return
	}

	cur := head
	for cur != nil {
		fmt.Printf("%d ", cur.Value)
		cur = cur.Next
		if cur == head { // 遍历到头节点时停止
			break
		}
	}
	fmt.Println()
}

// 解决约瑟夫问题
func JosephusProblem(n, m int) int {
	if n <= 0 || m <= 0 {
		return -1
	}

	// 创建循环链表
	head := CreateCircularLinkedList(n)
	PrintCircularLinkedList(head)

	// 找到要删除的节点前一个节点
	prev := head
	for prev.Next != head {
		prev = prev.Next
	}
	PrintCircularLinkedList(prev)

	count := 0
	for head.Next != head { // 当链表只剩一个节点时停止
		count++
		if count == m {
			// 删除节点
			fmt.Println(head.Value)
			prev.Next = head.Next
			count = 0
		} else {
			prev = head
		}
		head = head.Next
	}

	return head.Value
}

func main() {
	//fmt.Println("Enter the number of people:")
	//var n int
	//fmt.Scanln(&n)
	//
	//fmt.Println("Enter the counting number:")
	//var m int
	//fmt.Scanln(&m)

	//result := JosephusProblem(n, m)
	result := JosephusProblem(7, 3)
	fmt.Printf("The last person remaining is: %d\n", result)
}
