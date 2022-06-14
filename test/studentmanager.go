package main

import (
	"fmt"
	"os"
)

//1.先建立一个student的结构体
//2.建立一个studentManager的结构体
//3.根据输入，进入对应的操作
func showManu() {
	fmt.Println("欢迎来到学生管理系统")
	fmt.Println("1.添加学生")
	fmt.Println("2.编辑学生")
	fmt.Println("3.查看所有学生")
	fmt.Println("4.退出系统")
}

func main() {
	sm := newStudentManager()
	for {
		showManu()
		var input int
		fmt.Print("请输入你的选择：")
		fmt.Scanf("%d\n", &input)
		fmt.Printf("你的选择是：%d\n", input)
		switch input {
		case 1: //添加学生
			sm.add()
		case 2: //编辑学生
			sm.modify()
		case 3: //查看所有学生
			sm.show()
		case 4: //退出系统
			os.Exit(0)

		default:
			continue
		}
	}

}
