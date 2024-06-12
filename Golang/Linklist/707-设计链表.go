/*
*@Author: Wenqiang
*@Date:   2023/4/28
*@Time:   10:48
 */
package main

type MyLinkedList struct {
	val  int
	next *MyLinkedList
}

// MyLinkedList()初始化MyLinkedList对象。
func Constructor(value int, next *MyLinkedList) MyLinkedList {
	node := MyLinkedList{}
	node.val = value
	node.next = next
	return node
}

// int get(int index) 获取链表中下标为 index 的节点的值。如果下标无效，则返回 -1
func (this *MyLinkedList) Get(index int) int {
	if index < 0 {
		return -1
	}
	value := 0
	for i := 0; i < index; i++ {
		value = this.val
		this = this.next
		if this == nil {
			return -1
		}

	}
	return value
}

// void addAtHead(int val) 将一个值为 val 的节点插入到链表中第一个元素之前。
// 在插入完成后，新节点会成为链表的第一个节点。
func (this *MyLinkedList) AddAtHead(val int) {
	preHeadNode := MyLinkedList{}
	preHeadNode.next = this
	preHeadNode.val = val

}

// void addAtTail(int val) 将一个值为 val 的节点追加到链表中作为链表的最后一个元素。
func (this *MyLinkedList) AddAtTail(val int) {
	newNode := MyLinkedList{}
	for this.next != nil {
		this = this.next
	}
	this.next = &newNode
	newNode.val = val
	newNode.next = nil
}

//void addAtIndex(int index, int val) 将一个值为 val 的节点插入到链表中下标为 index 的节点之前。
//如果 index 等于链表的长度，那么该节点会被追加到链表的末尾。如果index比长度更大该节点将不会插入到链表中。
//

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	length := 0
	for j := 0; this.next != nil; j++ {
		length += 1
		this = this.next
	}
	if index == length {
		for j := 0; this.next != nil; j++ {
			this = this.next
		}
		node := MyLinkedList{}
		node.val = val
		this.next = &node
		node.next = nil

	} else if index > length {
		return
	} else {
		for i := 0; i < index-1; i++ {
			this = this.next
		}
		nodeInsert := MyLinkedList{val: val, next: this.next}
		this.next = &nodeInsert
	}

}

// void deleteAtIndex(int index) 如果下标有效，则删除链表中下标为 index 的节点。
func (this *MyLinkedList) DeleteAtIndex(index int) {
	length := 0
	for this.next != nil {
		length += 1
		this = this.next
	}
	if index > length {
		return
	} else {
		for i := 0; i < index-2; i++ {
			this = this.next
		}
		this.next = this.next.next

	}
}
