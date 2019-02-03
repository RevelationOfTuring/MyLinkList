package LinkList

import "fmt"

type LinkNode struct {
	Data interface{}
	Next *LinkNode
}

type LinkList struct {
	Header *LinkNode
	Len int
}

func (l*LinkList)Create(data ...interface{}){
	if l==nil {
		return
	}
	//assisted pointer
	pCur:=l.Header
	lenData:=len(data)

	for i:=0;i<lenData;i++{
		newNode:=&LinkNode{data[i],nil}
		pCur.Next=newNode
		pCur=newNode
		l.Len++
	}
}

func (l*LinkList)Print(){
	if l==nil{
		return
	}
	//头指针里面是不存储任何data的，所以遍历数据最开始应该从header的next开始
	for i:=l.Header.Next;i!=nil;i=i.Next{
		fmt.Println(i.Data)
	}
}