package main

import "fmt"

type student struct {
	id int64
	name string

}
type studentMgr struct {
	allStudent map[int64]student
}



func (s studentMgr) showStudent(){
	for _,stu := range s.allStudent{
		fmt.Printf("学号:%d 姓名：%s\n",stu.id,stu.name)
	}

}
func (s studentMgr) addStudent(){
	var (
		stuid int64
		stuname string
	)
	fmt.Print("请输入学生id")
	fmt.Scanln(&stuid)

	fmt.Print("请输入学生姓名")
	fmt.Scanln(stuname)
	newStu:=student{
		id:stuid,
		name:stuname,
	}
	s.allStudent[newStu.id]=newStu



}