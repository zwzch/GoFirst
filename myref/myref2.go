package myref

import (
	//"fmt"
	"unsafe"
	//"strconv"
	"reflect"


)

var x struct {
	a bool
	b int16
	c []int
}
func Float64bits(f float64) uint64 { return *(*uint64)(unsafe.Pointer(&f)) }
func MyRef2()  {
	/*
	unsafe.Sizeof函数返回操作数在内存中的字节大小，参数可以是任意类型的表达式，但是它并不会对表达式进行求值。一个Sizeof函数调用是一个对应uintptr类型的常量表达式，
	因此返回的结果可以用作数组类型的长度大小，或者用作计算其他的常量。
	*/
	//fmt.Println(unsafe.Sizeof(float64(0)))
	//fmt.Println(unsafe.Sizeof(x))
	//fmt.Printf("%#016x\n", Float64bits(1.0))
	/*
	一个unsafe.Pointer指针也可以被转化为uintptr类型，然后保存到指针型数值变量中（译注：这只是和当前指针相同的一个数字值，并不是一个指针），
	然后用以做必要的指针数值运算。（第三章内容，uintptr是一个无符号的整型数，足以保存一个地址）这种转换虽然也是可逆的，
	但是将uintptr转为unsafe.Pointer指针可能会破坏类型系统，因为并不是所有的数字都是有效的内存地址。
	*/
	//pb := (*int16)(unsafe.Pointer(
	//	uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	//*pb = 42
	//fmt.Println(x.b) // "42"

		//fmt.Println(Sprintf(123))

	//反射是由 reflect 包提供的。 它定义了两个重要的类型, Type 和 Value. 一个 Type 表示一个Go类型. 它是一个接口, 有许多方法来区分类型以及检查它们的组成部分, 例如一个结构体的成员或一个函数的参数等. 唯一能反映 reflect.Type 实现的是接口的类型描述信息(§7.5), 也正是这个实体标识了接口值的动态类型.
	//	函数 reflect.TypeOf 接受任意的 interface{} 类型, 并以reflect.Type形式返回其动态类型:

	//t := reflect.TypeOf(3)
	//fmt.Println(t.String())
	//fmt.Println(t)

	/*其中 TypeOf(3) 调用将值 3 传给 interface{} 参数. 回到 7.5节 的将一个具体的值转为接口类型会有一个隐式的接口转换操作, 它会创建一个包含两个信息的接口值: 操作数的动态类型(这里是int)和它的动态的值(这里是3).

	因为 reflect.TypeOf 返回的是一个动态类型的接口值, 它总是返回具体的类型. 因此, 下面的代码将打印 "*os.File" 而不是 "io.Writer". 稍后, 我们将看到能够表达接口类型的 reflect.Type.
	*/
	//var w io.Writer = os.Stdout
	//fmt.Println(reflect.TypeOf(w))
	//v := reflect.ValueOf(3) // a reflect.Value
	//fmt.Println(v)          // "3"
	//fmt.Printf("%v\n", v)   // "3"
	//fmt.Println(v.String())

	//x := 2                   // value   type    variable?
	//a := reflect.ValueOf(2)  // 2       int     no
	//b := reflect.ValueOf(x)  // 2       int     no
	//c := reflect.ValueOf(&x) // &x      *int    no
	//d := c.Elem()            // 2       int     yes (x)
	//fmt.Println(a.CanAddr()) // "false"
	//fmt.Println(b.CanAddr()) // "false"
	//fmt.Println(c.CanAddr()) // "false"
	//fmt.Println(d.CanAddr()) // "true"

	//x := 2
	//d := reflect.ValueOf(&x).Elem()   // d refers to the variable x
	//px := d.Addr().Interface().(*int) // px := &x
	//*px = 3                           // x = 3
	//fmt.Println(x)

	/*
	Set方法将在运行时执行和编译时进行类似的可赋值性约束的检查。以上代码，变量和值都是int类型，但是如果变量是int64类型，那么程序将抛出一个panic异常，所以关键问题是要确保改类型的变量可以接受对应的值：
	*/
	//x := 1
	//rx := reflect.ValueOf(&x).Elem()
	//rx.SetInt(2)                     // OK, x = 2
	//rx.Set(reflect.ValueOf(3))       // OK, x = 3
	//rx.SetString("hello")            // panic: string is not assignable to int
	//rx.Set(reflect.ValueOf("hello")) // panic: string is not assignable to int
	//
	//var y interface{}
	//ry := reflect.ValueOf(&y).Elem()
	//ry.SetInt(2)                     // panic: SetInt called on interface Value
	//ry.Set(reflect.ValueOf(3))       // OK, y = int(3)
	//ry.SetString("hello")            // panic: SetString called on interface Value
	//ry.Set(reflect.ValueOf("hello")) // OK, y = "hello"
	/*
	一个可取地址的reflect.Value会记录一个结构体成员是否是未导出成员，如果是的话则拒绝修改操作。因此，CanAddr方法并不能正确反映一个变量是否是可以被修改的。另一个相关的方法CanSet是用于检查对应的reflect.Value是否是可取地址并可被修改的：
	*/


}


//	func Sprintf(x interface{}) string{
//		type stringer interface {
//			String() string
//		}
//		switch x := x.(type) {
//		case stringer:
//			return x.String()
//		case string:
//			return x
//		case int:
//			return strconv.Itoa(x)
//		case bool:
//			if x{
//				return "true"
//			}
//			return "false"
//
//			default:
//// array, chan, func, map, pointer, slice, struct
//			return "???"
//		}
//	}

