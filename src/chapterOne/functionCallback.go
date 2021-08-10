//文件中的所有其余代码都属于 main 包
package main

//引入 fmt 包
import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

//main 程序运行的入口，程序运行时首先运行
//go 静态类型
func main() {
	fmt.Println("hello world")
	fmt.Println(math.Floor(2.67))
	fmt.Println(strings.Title("hello zcy"))
	fmt.Println('a')
	fmt.Println(reflect.TypeOf(5.2))
	fmt.Println(reflect.TypeOf(1))
	var numA = 4.6 //通常类型声明可以省略
	fmt.Println(numA)
	//基本数据类型的 零值
	//float：0，int：0，string：空字符串，boolean：false
	var myFloat float64
	var myString string
	fmt.Println(myFloat, myString)
	//短变量声明：在同一个代码域中不能进行不能对同一个变量重复声明
	//短变量在声明时类型也是确定了，后续不能进行跨类型赋值，可以同类型赋值
	myVar := 123
	myVar = 234
	fmt.Println(myVar)

	//
	var numInt = 3
	var numFloat = 4.6
	//fmt.Println(numInt * numFloat)
	fmt.Println(float64(numInt) * numFloat)
}
