package main

import (
	"fmt"
	"os"
)

var smr studentMgr

func showMenu() {
	fmt.Println("欢迎进入学生查询系统")
	fmt.Println(`
		1、查询所有学生
		2、添加学生
		3、修改学生
		4、删除学生
		5、退出
		`)
}
func main() {
	smr = studentMgr{
		allStudent: make(map[int64]student, 100),
	}
	for {
		showMenu()
		fmt.Print("请输入序号:")
		var choice int

		fmt.Scanln(&choice)
		fmt.Println("您输入的是：", choice)
		switch choice {
		case 1:
			smr.showStudent()
		case 2:
			smr.addStudent()
		case 3:
			fmt.Println("还没完成")
		case 4:
			fmt.Println("未完成")
		case 5:
			os.Exit(1)
		default:
			fmt.Println("请输入1~5以内的数字")


		}
	}

}
