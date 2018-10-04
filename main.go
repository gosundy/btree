package main

import (
	"fmt"
	"math/rand"
	"time"
)

const Jie =3
var SUM=0
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
rand.Seed(time.Now().Unix())
data:=make([]int,0)
for i:=0;i<1000;i++{
	n:=rand.Intn(1000000000)
	data=append(data,n)
}
fmt.Println(data)
//data:=[]int{81,87,47,59,81,18,25,40,56,0}
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
 fmt.Println(SUM)
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
				if v==head.next.data{
					return nil
				}
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
				if v==head.next.data{
					return nil
				}
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
						node.next = head.next
						head.next = node
						break
					}
					head = head.next
				}
				if head.next==nil{
					head.next=node
				}
				//将新节点有分支的赋值与右边节点的左分支，因为遍历的时候是遍历的左分支
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
					if midNode.left!=nil{
						head.right=midNode.left
					}
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
		SUM=SUM+1
		head=head.next
	}
	 travelBt(head.right)
}
