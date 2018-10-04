package main

import (
	"fmt"
)

const Jie =3
type BtreeNode struct{
	NodeNum int
	head *Node
}
type Node struct{
	data int
	left *BtreeNode
	right *BtreeNode
	next *Node

}
func main() {
 var root *BtreeNode
 data:=[]int{1,2,3,4,5,6,7,8,9,0,65,34,68,90}
 for _,v:=range data{
 	node:=insert(&root,v)
 	//产生溢出
 	if node!=nil{
 		root=node
	}
 }
 //insert(&root,data[0])
 //insert(&root,data[1])
 //root=insert(&root,data[2])
 //insert(&root,data[3])
 travelBt(root)
}
func insert(btNode **BtreeNode,v int)*BtreeNode{
	if *btNode==nil{
		*btNode=&BtreeNode{}
		btNode:=*btNode
		btNode.head=&Node{}
		btNode.head.next=&Node{data:v}
		btNode.NodeNum=btNode.NodeNum+1
		return nil
	}else{
		head:=(*btNode).head
		//判断是叶子节点
		if head.next.left==nil{
			//先插入节点
			for head.next!=nil{
				if v<head.next.data{
					node:=&Node{data:v}
					node.next=head.next
					head.next=node
					break
				}
				head=head.next
			}
			//插入末尾
			if head.next==nil{
				head.next=&Node{data:v}
			}
			(*btNode).NodeNum=(*btNode).NodeNum+1
			//插入之后判断是否产生上溢出
			//如果节点个数等于阶，则产生上溢出分裂
			if (*btNode).NodeNum==Jie{
				//找到中位数
				mid:=((*btNode).NodeNum+1)/2
				head=(*btNode).head
				for i:=1;i<mid;i++{
					head=head.next
				}
				midNode:=head.next
				leftBtNode:=&BtreeNode{NodeNum:mid-1,head:(*btNode).head}
				rightBtNode:=&BtreeNode{NodeNum:(*btNode).NodeNum-mid,head:&Node{}}
				rightBtNode.head.next=midNode.next
				midNode.left=leftBtNode
				midNode.right=rightBtNode
				head.next=nil
				midNode.next=nil
				midBtNode:=&BtreeNode{NodeNum:1,head:&Node{}}
				midBtNode.head.next=midNode
				return midBtNode
			}
			return nil
			//非叶子节点
		}else{
			var preBtNode *BtreeNode
			head=(*btNode).head
			for head.next!=nil{
				if v<head.next.data{
					preBtNode=insert(&head.next.left,v)
					break
				}
				head=head.next
			}
			if head.next==nil{
				preBtNode=insert(&head.right,v)
			}
			//下游节点产生溢出
			if preBtNode!=nil {
				head := (*btNode).head
				node:=preBtNode.head.next
				//先插入节点
				for head.next != nil {
					if node.data < head.next.data {
						node.next = head.next.next
						head.next = node

						(*btNode).NodeNum = (*btNode).NodeNum + 1
						break
					}
					head = head.next
				}
				if head.next==nil{
					head.next=node
				}
				//左分支进行比较
				if node.next!=nil{
					node.next.left=node.right
					node.right=nil
				}
				(*btNode).NodeNum=(*btNode).NodeNum+1
				//插入之后判断是否产生上溢出
				//如果节点个数等于阶，则产生上溢出分裂
				if (*btNode).NodeNum == Jie {
					//找到中位数
					mid := ((*btNode).NodeNum + 1) / 2
					head = (*btNode).head
					for i := 1; i < mid; i++ {
						head = head.next
					}
					midNode:=head.next
					leftBtNode:=&BtreeNode{NodeNum:mid-1,head:(*btNode).head}
					rightBtNode:=&BtreeNode{NodeNum:(*btNode).NodeNum-mid,head:&Node{}}
					rightBtNode.head.next=midNode.next
					midNode.left=leftBtNode
					midNode.right=rightBtNode
					head.next=nil
					midNode.next=nil
					midBtNode:=&BtreeNode{NodeNum:1,head:&Node{}}
					midBtNode.head.next=midNode
					return midBtNode
				}
				return nil
			}

		}

	}
	return nil
}
func travelBt(root *BtreeNode){
	if root==nil{
		return
	}
	head:=root.head
	for head.next!=nil{
		travelBt(head.next.left)
		fmt.Println(head.next.data)
		head=head.next
	}
	travelBt(head.right)
}
