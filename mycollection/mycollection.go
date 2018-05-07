package mycollection

import "fmt"

//import "fmt"
func MyCollection() {
	//var a [10]int = [10] int{1,2,3,4,5,6,7,8,9,10}
	////fmt.Println(a)
	//for i,v:=range a{
	//	fmt.Println(i,v)
	//	v=888
	//临时变量
	//}
	//fmt.Println(a)

	//for i,v:=range a{
	//	fmt.Println(i,v)
	//	a[i] = 888
	//}
	//fmt.Println(a)

	//var b =[...]int{1,2,3,4,5,6,7,8,9,10}
	//fmt.Println(b)

	//var c = new ([10]int)
	//

	//1、make只能用来分配及初始化类型为slice，map，chan的数据；new可以分配任意类型的数据
	//
	//2、new分配返回的是指针，即类型*T；make返回引用，即T；
	//
	//3、new分配的空间被清零，make分配后，会进行初始化。effective go举了一个例子，见：http://golang.org/doc/effective_go.html#allocation_make
	//for i:=1;i<11;i++{
	//	c[i-1] = i
	//}
	//fmt.Println(c,*c)
	//
	//var ptr *[10]int = c
	//fmt.Println(ptr)
	//fmt.Println(ptr[1],c[1],(*c)[1],(*ptr)[1])

	//var ptrnum = new ([10]*int)
	//for i:=0;i<10;i++{
	//	ptrnum[i] = &a[i]
	//	//存a里元素的地址
	//
	//}
	//fmt.Println(ptrnum,*ptrnum)
	//fmt.Println(*ptrnum[1],*(*ptrnum)[1])

	//Slience
	//var a1 [] int  = make([]int,3,10)//长度3 容量10
	//fmt.Println(a1,len(a1),cap(a1))

	//var b1 = [10]int{1,2,3,4,5,6,7,8,9,10}
	//ab1:=b1[1:5]
	//ab2:=b1[2:6]
	//起始点 终止点
	//ab1[1] = 999
	//
	//fmt.Println(ab1,ab2,b1)
	//ab1=append(ab1,1,2,3,4,5,6,7,8,9,10)
	//ab1[1]=888
	//fmt.Println(ab1,ab2,b1)

	//var c1 =new([]int)
	//fmt.Println(c1)
	////c1=append(c1,1,2,3,4,5)　//由于c1是指向[]int的指针，然而没有调用make的情况下，系统不可能给他一个数组去指，因此更谈不上追加，因为他是野指针
	//c1=&ab1
	//var tmp []int
	//tmp=append(*c1,1,2,3,4,5)
	//c1=&tmp
	//fmt.Println(c1)

	//var d1=[]int{1,2,3,4,5}
	//var e1=[]int{10,20,30}
	//copy(d1,e1)
	//fmt.Println(d1,e1)
	////后面复制到前面
	//
	//var d11=[]int{1,2,3,4,5}
	//var e11=[]int{10,20,30}
	//copy(e11,d11)
	//fmt.Println(d11,e11)

	//
	//
	//var a2 map[int]string=make(map[int]string)
	//a2[0]="aaa"
	//fmt.Println(a2)
	//
	//fmt.Println("======")
	//
	//var b2=map[int]string{0:"aaa",1:"bbb",2:"ccc"}
	//
	////两个值
	//for k,v:=range b2{
	//	fmt.Println(k,v)
	//	b2[k]="ok"
	//}
	//fmt.Println(b2)
	//
	//fmt.Println("======")
	//
	//var c2 =make(map[int]map[int]string)
	////map套map
	//c2[0]=make(map[int]string)
	//c2[0][0]="aaa"
	//c2[0][1]="bbb"
	//c2[1]=make(map[int]string)
	//c2[1][0]="ccc"
	//c2[1][1]="ddd"
	//fmt.Println(c2)
	//
	//fmt.Println("======")
	//
	//var d2=map[int]string{0:"aaa",1:"bbb",2:"ccc"}
	//var e2=new(map[int]string)
	//e2=&d2
	//fmt.Println(e2)

	/*
		myData
	*/
	/*
	Array
	• 数组是值类型,赋值和传参会复制整个数组,而而不是指针。
	• 数组⻓长度必须是常量,且是类型的组成部分。[2]int 和 [3]int 是不同类型。
	• 支支持 "=="、"!=" 操作符,因为内存总是被初始化过的。
	• 指针数组 [n]*T,数组指针 *[n]T。*/
	//a:=[3]int{1,2}
	//未初始化元素值为0
	//print(a[2])
	//b:=[...]int{1,2,3,4}
	//推断数组长度
	//c := [5]int{2: 100, 4:200} // 使用用索引号初始化元素。
	//d := [...]struct {
	//	name string
	//	age
	//	uint8
	//}{
	//	{"user1", 10}, // 可省略元素类型。
	//	{"user2", 20}, // 别忘了最后一一行行的逗号。
	//}

	//值拷⻉贝行行为会造成性能问题,通常会建议使用用 slice,或数组指针。

	//a:=[2]int{}
	//fmt.Printf("a: %p\n", &a)
	//test(a)
	//fmt.Println(a)

	//a := [2]int{}
	//fmt.Println(len(a),cap(a))

	/*
		Slice
	• 引用用类型。但自自身身是结构体,值拷⻉贝传递。
	• 属性 len 表示示可用用元素数量,读写操作不能超过该限制。
	• 属性 cap 表示示最大大扩张容量,不能超出数组限制。
	• 如果 slice == nil,那么 len、cap 结果都等于 0。
	*/
	//data := [...]int{0,1,2,3,4,5,6}
	//slice := data[1:4:5]
	////// [low : high : max]
	//fmt.Println(slice)
	//创建表达式使用用的是元素索引号,而而非非数量。
	//data := [...]int{0,1,2,3,4,5}
	//s:=data[2:4]
	//读写操作实际⺫目目标是底层数组,只需注意索引号的差别
	//s[0]+=100
	//s[1]+=200
	//fmt.Println(s)
	//fmt.Println(data)

	//可直接创建 slice 对象,自自动分配底层数组。
	//s1 :=[]int{0,1,2,3,8:100}
	//fmt.Println(s1, len(s1), cap(s1))

	//s2 := make([]int,6,8)
	//fmt.Println(s2,len(s2),cap(s2))
	//长度容量
	//s3 := make([]int, 6)
	//// 省略 cap,相当于 cap = len。
	//fmt.Println(s3, len(s3), cap(s3))
	//s := []int{0,1,2,3}
	//p:=&s[2]
	//*p+=100
	//fmt.Println(s)

	//data := [][]int{
	//	[]int{1, 2, 3},
	//	[]int{100, 200},
	//	[]int{11, 22, 33, 44},
	//}
	//fmt.Println(data)
	//d := [5]struct{
	//	x int
	//}{}
	//s := d[:]
	//d[1].x = 10
	//s[2].x = 20
	//fmt.Println(d)
	//fmt.Printf("%p, %p\n", &d, &d[0])

	/*reslice*/

	//s:=[]int{0,1,2,3,4,5,6,7,8,9}
	//s1:=s[2:5]
	//s2:=s1[2:6:7]
	//s3:=s2[3:6]
	//新对象依旧指向原底层数组。
	//fmt.Println(s,s1,s2)

	//s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//s1 := s[2:5]
	//// [2 3 4]
	//s1[2] = 100
	//s2 := s1[2:6]
	//// [100 5 6 7]
	//s2[3] = 200
	//fmt.Println(s)

	/*
		append
	*/
	//s :=make([]int,0,5)
	//fmt.Printf("%p\n", &s)
	//s2:=append(s,1)
	//fmt.Printf("%p\n", &s2)
	//append不是一个数组
	//简单点说,就是在 array[slice.high] 写数据。
	//data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//s := data[:3]
	//s2 := append(s, 100, 200)
	//// 添加多个值。
	//fmt.Println(data)
	//fmt.Println(s)
	//fmt.Println(s2)
	//data := [...]int{0, 1, 2, 3, 4, 10: 0}
	//s := data[:2:3]
	//s = append(s, 100, 200) // 一一次 append 两个值,超出 s.cap 限制。
	//fmt.Println(s, data) // 重新分配底层数组,与原数组无无关。
	//fmt.Println(&s[0], &data[0]) // 比比对底层数组起始指针。

	//s := make([]int, 0, 1)
	//c := cap(s)
	//for i := 0; i < 50; i++ {
	//	s = append(s, i)
	//	if n := cap(s); n > c {
	//		fmt.Printf("cap: %d -> %d\n", c, n)
	//		c = n
	//	}
	//}

	//copy
	//函数 copy 在两个 slice 间复制数据,复制⻓长度以 len 小小的为准。两个 slice 可指向同一一
	//底层数组,允许元素区间重叠。

	//data := [...]int{0,1,2,3,4,5,6,7,8,9}
	//s:=data[8:]
	//s2:=data[:5]
	//copy(s2,s)
	//fmt.Println(s2)
	//引用的拷贝

	/*
	MAP
	*/
	//m:=map[int]struct{
	//	name string
	//	age int
	//}{
	//	1: {"user1", 10},
	//	2: {"user2", 20},
	//}
	//fmt.Println(m[1].name)
	//预先给 make 函数一一个合理元素数量参数,有助于提升性能。因为事先申请一一大大块内存,
	//	可避免后续操作时频繁扩张。
	//m := make(map[string]int, 1000)
	//m:=map[string]int{
	//	"a":1,
	//}
	//if v,ok :=m["a"];ok{
	//	fmt.Println(v)
	//}

	//fmt.Println(m["c"])
	//m["b"]=2
	//delete(m,"c")
	//fmt.Println(len(m))
	//for k,v := range m{
	//	fmt.Println(k,v)
	//}
	//type user struct {
	//	name string
	//}
	//m := map[int]user{
	//	1:{"user1"},
	//}
	//m[1].name = "Tom"		error

	//u:=m[1]
	//u.name = "Tom"
	//m[1] = u
	//m2:=map[int] *user{
	//	1:&user{"user1"}
	//}
	//
	//m2[1].name = "JACK"
	//返回指针地址 可以修改原对象
	//可以在迭代时安全删除键值。但如果期间有新增操作,那么就不知道会有什么意外了。?

	//for i:=0;i<5;i++{
	//	m:=map[int]string{
	//		0: "a", 1: "a", 2: "a", 3: "a", 4: "a",
	//	    5: "a", 6: "a", 7: "a", 8: "a", 9: "a",
	//	}
	//	for k:= range m{
	//		m[k+k] = "x"
	//		delete(m,k)
	//	}
	//	fmt.Println(m)
	//}

	/*
		struct
	*/
	//n1:=Node{
	//	id:1,
	//	data:nil,
	//}
	//n2 := Node{
	//	id:
	//	2,
	//	data: nil,
	//	next: &n1,
	//}
	//fmt.Println(n1,n2)
	//
	//type File struct {
	//	name string
	//	size int
	//	attr struct {
	//		perm int
	//		owner int
	//	}
	//}
	//f:=File{
	//	name:"test.txt",
	//	size:1025,
	//}
	//f.attr.owner = 1
	//f.attr.perm = 0755
	//
	//var attr = struct {
	//	perm int
	//	owner int
	//}{2,0775}
	//f.attr = attr
	//fmt.Println(f)

	//type User struct {
	//	id int
	//	name string}
	//m:=map[User] int{
	//	User{1,"Tom"}:100,
	//}
	//fmt.Println(m)

	//var null struct{}
	//set := make(map[string]struct{})
	//set["a"] = nul
	//空结构 "节省" 内存,比比如用用来实现 set 数据结构,或者实现没有 "状态" 只有方方法的 "静
	//态类"。
	/*	匿名字段  */
	//可以像普通字段那样访问匿名字段成员,编译器从外向内逐级查找所有层次的匿名字段,
	//直到发现目目标或出错。
	//type Resource struct {
	//	id int
	//}
	//type User struct {
	//	Resource
	//	name string
	//}
	//type Manager struct {
	//	User
	//	title string
	//	}
	//var m Manager
	//m.id =1
	//m.name = "Jack"
	//m.title = "Admin"
	//fmt.Println(m)

	//外层同名字段会遮蔽嵌入入字段成员,相同层次的同名字段也会让编译器无无所适从。解决方方
	//法是使用用显式字段名。
	//type Resource struct {
	//	id int
	//	name string
	//}
	//type Classify struct {
	//	id int
	//}
	//type User struct {
	//	Resource
	//	// Resource.id 与 Classify.id 处于同一一层次。
	//	Classify
	//	name string
	//	// 遮蔽 Resource.name。
	//}
	//u:=User{
	//	Resource{1, "people"},
	//	Classify{100},
	//	"Jack",
	//}
	//fmt.Println(u)

	//type Resource struct {
	//	id int
	//	name string
	//}
	//type Classify struct {
	//	id int
	//}
	//type User struct {
	//	Resource
	//	// Resource.id 与 Classify.id 处于同一一层次。
	//	Classify
	//	name string
	//	// 遮蔽 Resource.name。
	//}
	//u:=User{
	//	Resource{1, "people"},
	//	Classify{100},
	//	"Jack",
	//}
	//fmt.Println(u)


	}
type Node struct {
	_ int
	id int
	data *byte
	next *Node
}
func test(x [2]int){
	fmt.Printf("x: %p\n", &x)
	x[1] = 1000
}