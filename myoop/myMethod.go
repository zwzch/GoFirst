package myoop

//import "fmt"

/*
	面向对象
	• 只能为当前包内命名类型定义方方法。
	• 参数 receiver 可任意命名。如方方法中未曾使用用,可省略参数名。
	• 参数 receiver 类型可以是 T 或 *T。基类型 T 不能是接口口或指针。
	• 不支支持方方法重载,receiver 只是参数签名的组成部分。
	• 可用用实例 value 或 pointer 调用用全部方方法,编译器自自动转换。
*/
//type Queue struct {
//	elements[] interface{}
//	//interface数组
//}
//
//func NewQueue() *Queue  {
//	return &Queue{make([]interface{},10)}
//	//创建对象实例
//	}
//func (*Queue) Push(e interface{}) error {
//	// 省略 receiver 参数名。
//	panic("not implemented")
//}
//func (xx *Queue) length() int {
//	// receiver 参数名可以是 self、this 或其他。
//	return len(xx.elements)
//}

//方法不过是一一种特殊的函数,只需将其还原,就知道 receiver T 和 *T 的差别


func MyMethod(){
	//d:=Data{}
	//p:=&d
	//fmt.Println("Data")
	}
