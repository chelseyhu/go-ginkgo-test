package main

import "fmt"

//1.先建立一个student的结构体
type student struct {
	ID    int8
	Name  string
	Class string
}

//student的构造函数
func newStudent(id int8, name string, class string) *student {
	return &student{
		ID:    id,
		Name:  name,
		Class: class,
	}
}

//2.建立一个studentManager的结构体
type studentManager struct {
	studentmgr []*student
}

//studentmanager的构造函数
func newStudentManager() *studentManager {
	return &studentManager{
		studentmgr: make([]*student, 0, 10),
	}
}

//studentmgr类的方法包含添加和修改以及展示
func (s *studentManager) add() {
	newstu := getInformation()
	s.studentmgr = append(s.studentmgr, newstu)
	return
}

func (s *studentManager) modify() {
	newstu := getInformation()
	for i, v := range s.studentmgr {
		if v.ID == newstu.ID {
			//*v = &newstu //或者用索引的方式去赋值
			s.studentmgr[i] = newstu
			return
		}
	}
	fmt.Println("学生信息有误，找不到该学生")
}
func (s *studentManager) show() {
	for _, v := range s.studentmgr {
		fmt.Printf("编号：%d 姓名：%s 班级：%s\n", v.ID, v.Name, v.Class)
	}
}
func getInformation() *student {

	var (
		id    int8
		name  string
		class string
	)
	fmt.Println("请按要求输入学生信息：")
	fmt.Print("请输入学生编号：")
	fmt.Scanf("%d\n", &id)
	fmt.Print("请输入学生姓名：")
	fmt.Scanf("%s\n", &name)
	fmt.Print("请输入学生班级：")
	fmt.Scanf("%s\n", &class)
	newstu := newStudent(id, name, class)
	return newstu
}
