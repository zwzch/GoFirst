package myoop

import (
	"fmt"
	//"unsafe"
	//"reflect"
)

/*
	接口是一个或多个方方法签名的集合,任何类型的方方法集中只要拥有与之对应的全部方方法,
	就表示示它 "实现" 了该接口口,无无须在该类型上显式添加接口声明。
	所谓对应方方法,是指有相同名称、参数列表 (不包括参数名) 以及返回值。当然,该类型还
	可以有其他方方法。
	• 接口命名习惯以er 结尾,结构体。
	• 接口只有方法签名,没有实现。
	• 接口没有数据字段。
	• 可在接口中嵌入入其他接口。
	• 类型可实现多个接口。
*/
//type Stringer interface {
//	String() string
//}
////方法的规定
//type Printer interface {
//	Stringer
//	//接口中嵌套接口
//	Printer()
//	}
//type User struct {
//
//	id int
//	name string
//}
//
//func (self *User) String() string {
//	return fmt.Sprintf("user %d, %s", self.id, self.name)
//}
//func (self *User) Print() {
//	fmt.Println(self.String())
//}

/*
	空接口口 interface{} 没有任何方方法签名,也就意味着任何类型都实现了空接口口。其作用用类
	似面面向对象语言言中的根对象 object。
*/

//func Print(v interface{}){
//	fmt.Printf("%T: %v\n", v, v)
//}

//type Tester struct {
//	s interface{
//		String() string
//	}
//}
//type User struct {
//	id	int
//	name string
//}
//func (self *User) String() string {
//	return fmt.Sprintf("user %d, %s", self.id, self.name)
//}
func MyInterface() {

	//fmt.Println("xxx")
	//var t Printer = (*User{{1, "Tom"}}).Print()
	// *User 方方法集包含 String、Print。
	//Print(1)
	//Print("HelloWorld")
	//匿名接口口可用用作变量类型,或结构成员。
	//t:=Tester{&User{1,"Tom"}}
	//fmt.Print(t)
	//接口口对象由接口口表 (interface table) 指针和数据指针组成。

	//struct Iface
	//{
	//	Itab* tab;
	//	void* data;
	//};
	//struct Itab
	//{
	//	InterfaceType* inter;
	//	Type* type;
	//	void (*fun[])(void);
	//};
	/*
	接口口表存储元数据信息,包括接口口类型、动态类型,以及实现接口口的方方法指针。无无论是反
	射还是通过接口口调用用方方法,都会用用到这些信息。
	数据指针持有的是⺫目目标对象的只读复制品,复制完整对象或指针。
	*/
	//type User struct {
	//id int
	//name string}
	//u:=User{1,"Tom"}
	//var i  interface{}=u
	//u.id = 2
	//u.name = "Jack"
	//fmt.Printf("%v\n", u)
	//fmt.Printf("%v\n", i.(User))

	//var a interface{} = nil // tab = nil, data = nil
	//var b interface{} = (*int)(nil) // tab 包含 *int 类型信息, data = nil
	//type iface struct {
	//	itab, data uintptr
	//}
	//ia := *(*iface)(unsafe.Pointer(&a))
	//ib := *(*iface)(unsafe.Pointer(&b))
	//fmt.Println(a == nil, ia)
	//fmt.Println(b == nil, ib, reflect.ValueOf(b).IsNil())
	//interface包含tab和data
	//
	//var o interface{} =&User{1,"Tom"}
	//if i,ok :=o.(fmt.Stringer);ok{
	//	 fmt.Print(i)
	//}
	//u:=o.(*User)
	//fmt.Print(u)
	//interface的表的遍历

	//var o interface{} = &User{1,"Tom"}
	//switch v := o.(type) {
	//case nil:
	//	// o == nil
	//	fmt.Println("nil")
	//case fmt.Stringer:
	//	// interface
	//	fmt.Println(v)
	//case func() string:
	//	// func
	//	fmt.Println(v())
	//case *User:
	//	// *struct
	//	fmt.Printf("%d, %s\n", v.id, v.name)
	//default:
	//	fmt.Println("unknown")
	//}
	//
	//超集接口口对象可转换为子子集接口口,反之出错。
	//var o Printer = &User{1, "Tom"}
	//var s Stringer = o
	//fmt.Println(s.String())
	//接口技巧

	//type Tester interface {
	//	Do()
	//}
	//type FuncDo func()
	//func (self FuncDo) Do() { self() }
	}
//type Stringer interface {
//	String() string
//}
//type Printer interface {
//	String() string
//	Printer()
//}
//type User struct {
//	id   int
//	name string
//}
//func (self *User) String() string {
//	return fmt.Sprintf("%d, %v", self.id, self.name)
//}
//func (self *User) Print() {
//	fmt.Println(self.String())
//}