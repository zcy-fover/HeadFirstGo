package main

import "fmt"

var myParamStruct struct{
	name string
	age int
	schoolName string
	grade float64
	gender bool	//true:男，false：女
}

type myTypeStruct struct {
	name string
	age int
	schoolName string
	grade float64
	gender bool	//true:男，false：女
	family 	//内嵌 struct，此时可以直接在myTypeStruct中使用 family 中的属性
}

type family struct {
	father string
	mather string
}

type MyType string
type StrName string

func main() {
	myParamStruct.name = "zcy"
	myParamStruct.age = 22
	myParamStruct.schoolName = "小学"
	myParamStruct.grade  = 12.33
	myParamStruct.gender = true
	fmt.Println(myParamStruct)

	var studentOne myTypeStruct
	studentOne.name = "李四"
	studentOne.age = 21
	studentOne.schoolName = "小学"
	studentOne.grade  = 32.33
	studentOne.gender = false
	fmt.Println(studentOne)
	changeAge(studentOne)
	fmt.Println(studentOne)
	modifyStudent(&studentOne)
	fmt.Println(studentOne)
	fmt.Println((*&studentOne).name)

	studentTwo :=  myTypeStruct{
		name: "赵六",
		age: 33,
		grade: 33.22,
		schoolName: "aa",
		gender: true,
	}
	studentTwo.father = "AA"
	studentTwo.mather = "BB"
	fmt.Println(studentTwo)

	var strName StrName
	strName = StrName("zcc")
	typeTest := MyType("aaa")
	fmt.Println(strName, typeTest)

	value := MyType("zcy")
	value.typeMethod()
	valueTwo := MyType("lisi")
	valueTwo.typeMethod()
	value.typeMethodPointer()
	(&valueTwo).typeMethodPointer()
}

func changeAge(student myTypeStruct)  {
	student.age++
	fmt.Println(student)
}

func modifyStudent(student *myTypeStruct)  {
	student.name = "李四-王五"
	fmt.Println(student)
}

func (m MyType) typeMethod()  {
	fmt.Println("hello ", m)
}

func (m *MyType) typeMethodPointer()  {
	fmt.Println("hello -> Hi", m)
}