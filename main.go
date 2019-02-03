package main

import (
	"dataStructure/MyLinkList/LinkList"
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

	myLink:=new(LinkList.LinkList)
	//necessary
	myLink.Header=new(LinkList.LinkNode)
	myLink.Create(s1,s2,s3,s4,s5)
	myLink.Print()

}
