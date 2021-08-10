## Go 语法

### 语法基础

```go
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
    var numA int = 4 //通常类型声明可以省略
    fmt.Println(numA)
}
```

#### 零值

- float：0，int：0，string：空字符串，boolean：false

    ```go
    var myFloat float64
    var myString string
    fmt.Println(myFloat, myString)
    ```

#### 短变量

- 短变量声明：在同一个代码域中不能进行不能对同一个变量重复声明
- 短变量在声明时类型也是确定了，后续不能进行跨类型赋值，可以同类型赋值

    ```go
    myVar := 123
    myVar = 234
    fmt.Println(myVar)
    ```

#### 常量

- 使用 `const` 关键字定义，不是 `var`
- 必须在声明常量时进行赋值，并且不可以改变常量的值
- 常量没有 `:=` 语法

    ```go
    const StudentName string = "zhang"
    ```

#### 命名规则

对于变量、函数、类型的命名规则

- 【强制】名称必须以字母开头，可以有任意数量的字母或数字
- 【强制】如果变量、函数、类型是以大写字母开头 ，则认为它是可以导出的（可以在 main 包或者其他包中被引用），可以在当前包之外的包中被访问；如果是小写字母开头，则认为是未导出，只能在当前包中使用
- 【强制】命名时避免和 Go 的保留关键字重复，会造成对 Go 本身的类型无法使用
- 【约定】遵守驼峰式 命名
- 【约定】当名称的 含义在上下文中很明显时，可以用缩写来代替，例如：用 i 代替 index
- 【约定】包名应该全部小写，含义相当明确时可缩写
- 【约定】多个单词的包名应该全部小写，不是下划线或者驼峰式
- 变量命名不要与包名冲突

#### 转换

- 数学运算和比较运算要求包含的值具有相同的类型，类型不同需要运算或者比较时需要进行转换

    ```go
    var numInt = 3
    var numFloat = 4.6
    fmt.Println(numInt * numFloat) //Invalid operation: numInt*numFloat (mismatched types int and float64)
    fmt.Println(float64(numInt) * numFloat)
    ```

#### 命令

- `go fmt`：自动重新格式化源文件以便使用 Go 标准格式
- `go build`：将 go 源代码编译成计算机可执行的二进制文件
- `go run`：编译并运行一个程序，而不将可运行文件保存在当前目录

### 方法

方法是与特定类型的值关联的函数

- 时间方法

  ```go
  var now = time.Now()
  fmt.Println(now.Year())
  fmt.Println(now.String())
  ```

- 键盘输入

  ```
  fmt.Print("请输入内容：")
  reader := bufio.NewReader(os.Stdin)  //从标准输入（键盘）读取
  input, _ := reader.ReadString('\n')  //以字符串形式返回用户所有输入内容；换行符前的所有内容将被读取
  fmt.Println(input)
  ```

  > **_**：上面代码块中出现的“_”，表示**空白标识符**，空白标识符接收的值会被丢弃掉；Go 不允许定义变量却不使用，这种情况使用空白标识符来处理；

  > 正常情况下要对程序的异常（错误）返回进行处理，否则可能会对后面程序的运行造成意外情况

  ```go
  fmt.Print("请输入内容：")
  reader := bufio.NewReader(os.Stdin)
  input, err := reader.ReadString('\n')
  if err != nil {
  	log.Fatal(nil)	//打印错误并且终止程序
  }
  fmt.Println(input)
  ```

#### 块和变量的作用域

变量的作用域由其声明所在块和嵌套在该块中的其他块组成，声明的变量可以在其作用域任何地方被访问，在域外无法访问

### 函数

#### 函数定义和参数

```go
func calculateArea(width float64, height float64) (float64, error) {
    if width < 0 || height < 0 {
        return 0, fmt.Errorf("输入异常：width{%.2f}, height{%.2f}", width, height)
    }
    return width * height,  nil
}
```

Go 函数可以有多个返回值，Go 是一种**值传递**语言，函数的形式参数从调用中接收实参的副本

```go
func main() {
    aa := 5
    paramChange(aa)
    fmt.Println(aa)
}

func paramChange(num int)  {
    num =  10
    fmt.Println(num)
}
//运行后输出
//10
//5
```

即在调用的函数体中不会改变出入参数的值，这种和 `java` 的引用传递不同

#### 指针

可以利用 **&** 符号获取变量的地址，即变量在内存中的地址，也称为指针

```go
func paramAddress(num int)  {
    fmt.Println(&num)
}
//输出：0xc00001e098
```

##### 指针类型

指针类型表示为 ***变量类型** ，列入指向一个 int 类型变量的指针类型是 *int，声明的指针变量也只能保存一种类型的值的指针，例如将。int 指针赋值给 float  指针会编译错误

可以利用 ***指针变量** 获取指针指向的值，还可以通过 * 改变指针指向的值，此处是指针（内存地址）不变，但是该内存处的值被改变，所有引用该内存地址的变量的值都会被改变。这里可以做到在上述函数参数传递中改变原参数的值

```go
func paramPoint() {
    var myInt  = 10
    myIntPointer :=  &myInt
    fmt.Println(myIntPointer)
    fmt.Println(*myIntPointer) //输出 10
    *myIntPointer =  20
    fmt.Println(*myIntPointer) //输出 20
    fmt.Println(myInt) // 输出 20
    var  myFloat float64
    var myFloatPointer *float64
    myFloatPointer  = &myFloat
    fmt.Println(myFloatPointer)
    //myFloatPointer = &myInt，编译异常
}
```

##### 函数指针

```go
func main() {
    aa := 5
    fmt.Println(*funcPoint(&aa))  //输出 50
    fmt.Println(aa) //输出 10
}

func funcPoint(numPoint *int) *int {
    var result = *numPoint * 10
    *numPoint = 10
    return &result
}
```

### 数组

`var 变量名 [数组大小]数据类型{字面量}`，与变量一样，数组在创建时会给数组中每一项初始化为对应数据类型的零值

```go
var ageArray [2]int
ageArray[1] =  15
fmt.Println(ageArray[0])	//输出：0
var nameArray = [2]string{"aa", "bb"}
fmt.Println(nameArray[0])	//输出：aa
fmt.Printf("%#v\n", nameArray)	//输出：[2]string{"aa", "bb"}

for i := 0; i < len(nameArray); i++ {
    fmt.Print(nameArray[i], "--")
}

for index, value := range nameArray {
    fmt.Println(index, "---", value)
}

for _, value := range nameArray {
    fmt.Println(value)
}
```

使用`for ... range` 遍历数组，index 保存了索引，value 保存了值；这种方式不会引起无效数组的访问，当下文代码块不需要使用index 或者 value 时，可以用`_`空白标识符。

### 切片

`var 变量名 []数据类型{字面量}`，与数组不同的是切片在声明时，不指定大小。声明切片变量不会创建初始化该切片，一般使用内建函数 `make()` 来创建切片。

```go
var mySlice = []int{1, 2, 3}
notes := make([]string, 10)
notes[0] = "a"
notes[1] = "b"
fmt.Println(mySlice)  //[1 2 3]
fmt.Println(notes)	//[a b        ]
var myArray = []int{1, 2, 3, 4, 5, 6,  7, 8, 9}
sliceArray := myArray[1:5]	//通过数组创建切片
fmt.Println(sliceArray)    //[2 3 4 5]
sliceArray = append(sliceArray, 11, 12, 13)
fmt.Println(sliceArray)		//	[2 3 4 5 11 12 13]

sliceArray = append(sliceArray, 11, 12, 13)
fmt.Println(sliceArray)		//	[2 3 4 5 11 12 13]

newSlice := append(mySlice, 4, 5, 6)
fmt.Println(mySlice)	//[1 2 3]
fmt.Println(newSlice)	//[1 2 3 4 5 6]
```

通过数据创建的切片，切片是底层数组内容的视图，对于数组的修改会反映给所有的切片

Go 通过内建函数 append() 在一个切片尾部追加一个或者多个值，返回包含了老元素与新元素的新切片。切片的底层数组不能增长大小，如果在尾部添加元素数组空间不够时，会自动开辟新的空间将所有元素拷贝过来，返回新的切片指向地址。所以一般要用原切片变量接收 append() 函数的返回值，或者用其他变量接收。

切片变量的零值`nil`。

#### 可变长函数参数

```go
func paramFun(paramOne string, paramTwo ...string) {

}
```

函数可以同时设置一个或者多个可变长参数，仅函数定义的最后一个参数可以是可变长

### 映射

`var 变量名 map[键数据类型]值数据类型`{字面量}，与切片相同，声明映射变量不会初始化创建映射变量 ，需要调用 `make()` 函数。

```
var myMap map[string]int
myMap = make(map[string]int)
initMap := map[string]int{"a": 1, "b": 2}
```

#### 零值

映射变量的零值是 nil，初始化后的映射访问一个不存在的键时，得到零值时对应数据类型的零值。这种情况会造成无法判断该舰是已经存在于 map 中（不存在或者已经存在键值就是默认值）

```go
value, ok := initMap["a"]
delete(initMap, "b")
```

此时解决这个问题，在访问键值时返回了两个参数，第一个时该键对应的值，第二个时布尔值，键存在返回 true，不存在返回 false

### struct 结构体

可以定义 struct 的变量，也可以定义 struct 的类型

```go
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
    family 	//内嵌 struct，此时可以直接在myTypeStruct中使用 family 中的属性，匿名字段；当然也可以定义变量名来访问
}

type family struct {
	father string
	mather string
}

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
    
    studentTwo :=  myTypeStruct{	//使用短变量和字面量初始化创建struct类型变量
		name: "赵六",
		age: 33,
		grade: 33.22,
		schoolName: "aa",
		gender: true,
	}
    studentTwo.father = "AA"
	studentTwo.mather = "BB"
	fmt.Println(studentTwo)
}

func changeAge(student myTypeStruct)  {
	student.age++
	fmt.Println(student)
}

func modifyStudent(student *myTypeStruct)  {
	student.name = "李四-王五"
	fmt.Println(student)
}
```

在函数的参数中使用  struct 类型，如果是指针访问其中的字段注意格式：`(*指针变量).fieldName`，“&变量”表示指向了变量的地址即指针，“*指针变量”中 * 相当于指针运算符可以获取指针指向的值；通常可以省略 *，直接通过指针访问属性字段

如果要在其他的包中使用定义的  struct，则类型名称和对应需要导出的字段名称首字母都要大写

### 定义类型

#### 基础类型定义

```go
type MyType string
type StrName string

var strName StrName
strName = StrName("zcc")
typeTest := MyType("aaa")
```

可以把任何基础类型的值转换为定义的类型

#### 定义方法

```
func (m MyType) typeMethod()  {
	fmt.Println("hello ", m)
}

value := MyType("zcy")
value.typeMethod() 	//输出：hello zcy
valueTwo := MyType("lisi")
valueTwo.typeMethod()	输出：hello li si
```

`func (接收器参数名 接收器类型) typeMethod() {}` 方法定义是在函数名称前增加一个接收器参数类型和接收器参数名，一个方法被定义了某个类型后，可以被这个类型创建的所有变量调用。类似于其他语言中的 this 或 self 关键字(可隐式使用)，但在 Go 中是显式声明调用。几点要求：

- 接收器参数名一般使用接收器类型名称的首字母小写
- 方法和接收器参数类型必须要定义在同一个包中，不能跨包定义。确保了不会为一些基础数据类型定义新方法

##### 指针类型接收器参数

```
func (m *MyType) typeMethodPointer()  {
	fmt.Println("hello ", m)
}

value := MyType("zcy")
value.typeMethod() 	//输出：hello zcy
valueTwo := MyType("lisi")
&valueTwo.typeMethod()	输出：hello li si

value.typeMethodPointer()
(&valueTwo).typeMethodPointer()
```

与函数参数传递类似，如果不使用指针接收器接收的是拷贝值不会改变原值。接收器参数定义时可以利用指针类型的接收器参数。在调用时可以省略显式声明指针变量调用，Go 会做自动转换，当然也可以定义指针变量调用，但是代码中应该保持风格统一

### 封装嵌入

自定义类型后，内部数据并不想让包外程序直接访问到，破坏内部封闭原则。需要将部分数据封装，提供统一的对外访问方法。

```go
package calendar

import (
	"errors"
	"fmt"
)

type Date struct {
	year int
	month int
	day int
}

func (d *Date) SetDate(year int, month int, day int) error {
	err := d.SetDay(day)
	if err != nil {
		return err
	}
	err = d.SetMonth(month)
	if err != nil {
		return err
	}
	err = d.SetYear(year)
	if err != nil {
		return err
	}
	return nil
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("输入年无效")
	}
	d.year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("输入月份无效")
	}
	d.month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("输入天数无效")
	}
	d.day = day
	return nil
}

func (d *Date) Year() int {
	return d.year
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Day() int {
	return d.day
}

func (d *Date) Display()  {
	fmt.Println(*d)
}
```

例如上述代码如果 Date 的属性是可以导出的，则会破坏数据的有效性校验。未导出的变量、struct字段、函数、方法、等仍然能够被相同包的导出的函数或方法访问。Go 的未导出数据是包内可见，所以要做到包内数据安全。

```go
package main

import (
	"HeadFirstGo/src/chapterEight/calendar"
	"fmt"
	"log"
)

func main() {
	date := calendar.Date{}
	err := date.SetDate(2021, 8, 10)
	if err != nil {
		log.Fatal(err)
	}
	date.Display()

	event := calendar.Event{}
	event.Title = "测试标题"
	err = event.SetDate(2, 3, 5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Year(), event.Month(), event.Day())

	event.DisplayEvent()
}
```

这里也可以定义 get 方法输出封装的属性值，一般情况下使用属性字段名首字母大写来定义 get 方法名，不是 Get 开头。

```go
package calendar

import "fmt"

type Event struct {
	Title string
	Date	//嵌入匿名Date类型，Event可以拥有 Date 属性字段
}

//这里体现了包内非导出数据的可见性
var eventTest = Event{"aa", Date{1, 2, 4}}

func (e Event) DisplayEvent() {
	fmt.Println(eventTest)
}
```


