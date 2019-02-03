package main

import (
	"dataStructure/MyLinkList/LinkList"
	"fmt"
)

type Student struct {
	Name string
	Age int
	sex string
}

func main() {
	s1:=Student{"A",18,"male"}
	s2:=Student{"B",19,"female"}
	s3:=Student{"C",20,"male"}
	s4:=Student{"D",21,"male"}
	s5:=Student{"E",22,"female"}
	s6:=Student{"WGY",28,"male"}
	//s7:=Student{"sy",26,"female"}


	myLink:=new(LinkList.LinkList)
	//necessary
	myLink.Header=new(LinkList.LinkNode)
	myLink.Create(s1,s2,s3,s4,s5,s6)
	myLink.Print()
	fmt.Println(myLink.GetLen())
	//myLink.InsertByHeader(s6)
	//myLink.Print()
	//fmt.Println(myLink.GetLen())
	//myLink.InsertByTail(s6)
	//myLink.Print()
	//fmt.Println(myLink.GetLen())
	//myLink.InsertByPos(s6,0)
	//myLink.Print()
	//fmt.Println(myLink.GetLen())
	//myLink.DeleteByValue(s2)
	//myLink.Print()
	//fmt.Println(myLink.GetLen())
	myLink.DeleteByPos(6)
	myLink.Print()
	fmt.Println(myLink.GetLen())



}
