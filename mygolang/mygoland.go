package mygolang

import (
	//"fmt"
	//"strconv"
	//"unsafe"
	//"fmt"
	//"unsafe"
	//"unsafe"
	//"fmt"
	//"reflect"
	//"reflect"
	//"fmt"
	//"go/ast"
)

func Mygolang()  {
	//var i string
	//fmt.Scanln(&i)
	////fmt.Println(i)
	//j:=100
	//if j,_:=strconv.Atoi(i);j<5{ //if支持一个初始化表达式
	//	//占位符表示忽略这个变量 忽略err
	//	fmt.Println(j,"<5")
	//}else if j<10 {
	//	fmt.Println(j,"<10")
	//}else{
	//	fmt.Println(j,"else")
	//}
	//fmt.Println(j)
	//k:=0
	////while（true）
	//for{
	//	if k>=3{
	//		break
	//	}
	//	fmt.Println(k)
	//	k++
	//}
	//
	//for k:=0;k<3;k++{
	//	fmt.Println(k)
	//}


	/*
		默认break
		fallthrough执行下一个
	*/
	//j1:=0
	//switch j1 {
	//case 0:
	//	fmt.Println("00000")
	//	fallthrough
	//case 1:
	//	fmt.Println("11111")
	//	//fallthrough
	//
	//case 2:
	//	fmt.Println("22222")
	//default:
	//	fmt.Println("default")
	//
	//}

	//j1:=8
	//switch j2:=j1; {
	////switch可以赋值
	//case j2<5:
	//	fmt.Println("<5")
	//case j2<10:
	//	fmt.Println("<10")
	//	fallthrough
	//case j2<20:
	//	fmt.Println("<20")
	//default:
	//	fmt.Println("default")
	//
	//}


//
//
//xxx:
//	for i:=1;i<10;i++{
//		for j:=1;j<10;j++{
//			if i+j==5{
//				break xxx
//
//			}
//			fmt.Println(i+j)
//		}
//		fmt.Println("{{{")
//	}
//	fmt.Println("xxx")
//
//	fmt.Println("=======")

//xxx1:
//	//跳到标签在的循环
//	for i:=1;i<3;i++{
//		for j:=1;j<3;j++{
//			if i+j==5{
//				continue xxx1
//			}
//			fmt.Println(i+j)
//		}
//		fmt.Println("{{{1")
//	}
//	fmt.Println("xxx1")

//var i string
//aaa:
//	fmt.Scanln(&i)
//	jj,_:=strconv.Atoi(i)
//	if jj==0{
//		goto aaa
//	}else{
//		goto bbb
//	}
//	fmt.Println("-----")
//bbb:
//	fmt.Println(jj)

















//MyLanguage
//	编译器会将未使用用的局部变量当做错误。
//	s:="abc"
//	println(&s)
//	s, y :="hello", 20;
//	println(&s,y)


	//s, y := "hello", 20 // 重新赋值: 与前 s 在同一一层次的代码块中,且有新的变量被定义。
	//println(&s, y)
	//{
	//	s, z := 1000, 30
	//	// 定义新同名变量: 不在同一一层次代码块。
	//	println(&s, z)
	//	//指向的位置也不一样
	//}


	//const x,y int = 1,2//常量 	// 未使用用局部常量不会引发编译错误。
	//const s = "Hello World"
	//const(
	//	a, b int = 10,100
	//	c bool = false
	//)
	//注意静态代码块小括号

	//const(
	//	s = "abc"
	//	x //如不提供类型和初始化值 和上一常量形同
	//
	//)

	//const   (
	//	a = "abc"
	//	b = len(a)
	//	c = unsafe.Sizeof(b)
	//)
	//println(a,b,c)


	//const(
	//	a byte = 100
	//	b int = 1e20
	//)


	//枚举
	//关键字 iota 定义常量组中从 0 开始按行行计数的自自增枚举值。
	//const(
	//	Sunday = iota
	//	Monday // 1,通常省略后续行行表达式。
	//	Tuesday // 2
	//	Wednesday // 3
	//	Thursday // 4
	//	Friday // 5
	//	Saturday // 6
	//)


	//println(Saturday)



	//const (
	//	_= iota
	//// iota = 0
	//KB int64 =1 << (10 * iota)
	//int64 = 1 << (10 * iota)
	//MB
	//// iota = 1
	//// 与 KB 表达式相同,但 iota = 2
	//GB
	//TB
	//)
	//在同一一常量组中,可以提供多个 iota,它们各自自增⻓长。
	//const(
	//	A,B = iota, iota<<10
	//	C,D
	//)


	//iota打断显式恢复
	//const (
	//	A= iota
	//	B
	//	C= "c"
	//	D
	//	E = iota
	//	// 0
	//	// c
	//	// c,与上一一行行相同。
	//
	//	F
	//	// 4,显式恢复。注意计数包含了 C、D 两行行。
	//	// 5
	//)

	//type Color int
	//const(
	//	Black Color = iota
	//	Red
	//	Blue
	//)
	//println(Blue)


	//a:=[]int {0,0,0}
	//a[1] = 100
	//b:=make([]int,3)
	//b[1] = 10


	//内置函数 new 计算类型大大小小,为其分配零值内存,返回指针。
	//而而 make 会被编译器翻译成具体的创建函数,由其分配内存和初始化成员结构,返回对象而而非非指针。
	//c:=new([]int)
	//c[1] = 10


	//不支支持隐式类型转换,即便是从窄向宽转换也不行行。
	//var b byte = 100
	//var  n int  = b
	//var n int = int(b)
	//显式转换

/*

	字符串是不可变值类型,内部用用指针指向 UTF-8 字节数组。
	• 默认值是空字符串 ""。
	• 用用索引号访问某字节,如 s[i]。
	• 不能用用序号获取字节元素指针,&s[i] 非非法。
	• 不可变类型,无无法修改字节数组。
	• 字节数组尾部不包含 NULL。
*/

	//s:="abc"
	//println(s[0] == '\x61', s[1] == 'b', s[2] == 0x63)
	//使用下标访问字符串

	//s := `a
	//b\r\n\x00
	//c`
	//println(s)
	//``不转义的字符串

	//s:="hello,world"
	//s1:=s[:5]
	//s2:=s[6:]
	//s3:=s[1:5]
	//println(s1,s2,s3)

	//单引号字符常量表示示 Unicode Code Point,支支持 \uFFFF、\U7FFFFFFF、\xFF 格式。
	//对应 rune 类型,UCS-4。
	//fmt.Printf("%T\n", 'a')
	//var c1, c2 rune = '\u6211', '们'
	//println(c1 == '我', string(c2) == "\xe4\xbb\xac")

	//要修改字符串,可先将其转换成 []rune 或 []byte,完成后再转换为 string。无无论哪种转
	//换,都会重新分配内存,并复制字节数组
	//s:="abcd"
	//bs:=[]byte(s)
	////byte切片
	//bs[1] ='B'
	//fmt.Println(string(bs))

	//
	//u:="电脑"
	//us := []rune(u)
	//us[1] = '话'
	//fmt.Println(string(us))

	//s := "abc汉字"
	//for i := 0; i < len(s); i++ {
	//	// byte
	//	fmt.Printf("%c,", s[i])
	//}
	//fmt.Println()
	//for _, r := range s {
	//	// rune
	//	fmt.Printf("%c,", r)
	//}
	//rune unicode单字



	//支持指针类型 *T,指针的指针 **T,以及包含包名前缀的 *<package>.T。
	//• 默认值 nil,没有 NULL 常量。
	//• 操作符 "&" 取变量地址,"*" 透过指针访问⺫目目标对象。
	//• 不支支持指针运算,不支支持 "->" 运算符,直接用用 "." 访问⺫目目标成员。
	//指针无法操作
	//type data struct {
	//	a int
	//}
	//结构体
	//var d  = data{1234}
	//var p *data;
	//p = &d;
	//fmt.Printf("%p, %v\n", p, p.a)
	//指针和值
	//指针不能操作

	//x:=0x12345678
	//p:=unsafe.Pointer(&x)
	//n := (*[4]byte)(p)
	////转化为四个字节的数组
	////unsafe 转化
	//for i := 0; i < len(n); i++ {
	//	fmt.Printf("%X ", n[i])
	//}


		//fmt.Println(test())


		//d:= struct {
		//	s string
		//	x int
		//}{"abc",100}
		//
		//p := uintptr(unsafe.Pointer(&d))//// *struct -> Pointer -> uintptr
		//p +=  unsafe.Offsetof(d.x)//unsafe指针可以操作
		//p2 := unsafe.Pointer(p)
		//px := (*int)(p2)
		//*px = 200
		//fmt.Printf("%#v\n", d)
		//注意:GC 把 uintptr 当成普通整数对象,它无无法阻止止 "关联" 对象被回收


		//var a struct{x int `a`}
		//var b struct {x int `ab`}
		//fmt.Println(a.x,b.x)

		//type bigint int64
		//var x bigint = 100
		//println(x)

	//新类型不是原类型的别名,除拥有相同数据存储结构外,它们之间没有任何关系,不会持
	//有原类型任何信息。除非非⺫目目标类型是未命名类型,否则必须显式转换。
	//	type bigint int64
	//	var x bigint = 100
	//	fmt.Println(reflect.TypeOf(x))

	//x := 1234
	//var b bigint = bigint(x)
	//var b2 int64 = int64(b)
	//必须显示转换


	//var s myslice = []int{1, 2, 3}
	// 未命名类型,隐式转换。
	//var s2 []int = s




	//var a= struct {
	//	x int
	//}{100}
	//var b  = []int{1,2,3}
	//b := []int{
	//	1,
	//	2 }
	// ok



	//r := '我'
	//fmt.Printf("%q的类型为：%t 二进制为：%b\n", r, r, r)
	//
	//rType := reflect.TypeOf(r).Kind()
	//fmt.Printf("r的实际类型为：%s\n", rType)
//IF
	//x :=0
	//if n:="abc";x>0{
	//	fmt.Println(n[2])
	//}else if x <0{
	//	fmt.Println(n[1])
	//}else {
	//	fmt.Println(string(n[0]))
	//}



	//FOR
	//s:="abc"
	//for i,n:= 0, len(s);i<n;i++{
	//	fmt.Println(s[i])
	//}
	//for循环

	//n := len(s)-1
	////println(n)
	//for n>=0 {
	//	println(s[n])
	//	n--
	//}
	//while循环


	//for {
	//	println(s)
	//}
	//while true

	//类似迭代器操作,返回 (索引, 值) 或 (键, 值)。
	//s :="abc"
	//for i:= range s{
	//	println(s[i])
	//}

	//for _, c := range s {
	//	// 忽略 index。
	//	println(c)
	//}


	//for range s{
	//	fmt.Println(s)
	//}
	//忽略全部返回值

	//m:=map[string]int{"a":1,"b":2}
	//for k,v :=range m{
	//	fmt.Println(k,v)
	//}

	//range 会复制对象
	//会生成一个新的对象
	//a:=[3]int{0,1,2}
	//for i,v :=range a{
	//	if i==0{
	//		a[1],a[2] = 999,999
	//		fmt.Println(a)
	//	}
	//	a[i] = v+100
	//}
	//fmt.Println(a)
	//


	//使用应用类型 底层数据不会被复制

	//s :=[] int{1,2,3,4,5}
	////slice底层是struct
	//for i,v := range s{
	//	//i是索引
	//	if i == 0 {
	//		s = s[:3]
	//		//从多少开始切多少个
	//		s[2] = 100
	//	}
	//	fmt.Println(i,v)
	//}




	//SWITCH
	//x := []int{1,2,3}
	//i:=1
	//switch i {
	//case x[1]:
	//	fmt.Println("a")
	//case 1,3:
	//	fmt.Println("b")
	//	fallthrough
	//	//使用下一个分支
	//default:
	//	fmt.Println("c")
	//
	//}
	//可以有条件表达式
	//var i int
	//for {
	//	fmt.Println(i)
	//	i++
	//	if i>2{
	//		goto BREAK
	//	}
	//}
	//BREAK:
	//	fmt.Println("break")
		//EXIT: 不使用无法定义
	//多层循环跳出
	//break 可用用于 for、switch、select,而而 continue 仅能用用于 for 循环。


	//不支支持 嵌套 (n ested)、重载 (overload) 和 默认参数 (default parameter)。
	//• 无无需声明原型。
	//• 支支持不定⻓长变参。
	//• 支支持多返回值。
	//• 支支持命名返回参数。
	//• 支支持匿名函数和闭包。


	}

