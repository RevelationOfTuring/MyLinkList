package LinkList

import "fmt"

type LinkNode struct {
	Data interface{}
	Next *LinkNode
}

type LinkList struct {
	Header *LinkNode
	len    int
}

func (l *LinkList) Create(data ...interface{}) {
	if l == nil {
		return
	}
	//assisted pointer
	pCur := l.Header
	lenData := len(data)

	for i := 0; i < lenData; i++ {
		newNode := &LinkNode{data[i], nil}
		pCur.Next = newNode
		pCur = newNode
		l.len++
	}
}

func (l *LinkList) Print() {
	if l == nil {
		return
	}
	//头指针里面是不存储任何data的，所以遍历数据最开始应该从header的next开始
	for i := l.Header.Next; i != nil; i = i.Next {
		fmt.Println(i.Data)
	}
}

func (l *LinkList) GetLen() int {
	if l == nil {
		return -1
	}
	return l.len
}

func (l*LinkList) InsertByHeader(data interface{}){
	if l==nil {
		return
	}
	if data==nil{
		return
	}
	newNode:=&LinkNode{data,nil}
	newNode.Next=l.Header.Next
	l.Header.Next=newNode
	l.len++
}

func (l*LinkList) InsertByTail(data interface{}){
	if l==nil{
		return
	}
	if data==nil{
		return
	}
	pCur:=l.Header
	for {
		if pCur.Next!=nil{
			pCur=pCur.Next
		}else{
			newNode:=new(LinkNode)
			newNode.Data=data
			pCur.Next=newNode
			l.len++
			break
		}
	}
}

func (l*LinkList)InsertByPos(data interface{},pos int){
	if l==nil{
		return
	}
	if data==nil {
		return
	}
	if pos<=1{
		l.InsertByHeader(data)
		return
	}
	if pos>l.len{
		l.InsertByTail(data)
		return
	}
	pCur:=l.Header
	for i:=1;i<=pos-1;i++{
		pCur=pCur.Next
	}

	newNode:=new(LinkNode)
	newNode.Data=data
	newNode.Next=pCur.Next
	pCur.Next=newNode
	l.len++
}

func (l*LinkList)DeleteByValue(data interface{}){
	if l==nil {
		return
	}
	if data==nil{
		return
	}
	pCur:=l.Header
	var pPre *LinkNode
	for pCur!=nil {
		if pCur.Data!=data{
			pPre=pCur
			pCur=pCur.Next
		}else{
			pPre.Next=pCur.Next
			l.len--
			break
		}
	}
}

func (l*LinkList)DeleteByPos(pos int){
	if l==nil{
		return
	}
	if pos<1{
		pos=1
	}
	if pos>l.len{
		pos=l.len
	}
	pCur:=l.Header
	for i:=1;i<=pos-1;i++{
		pCur=pCur.Next
	}
	pCur.Next=pCur.Next.Next
	l.len--
}