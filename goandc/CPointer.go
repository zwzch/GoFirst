package goandc

import "C"
import (
	//"unsafe"
	//"time"
	//"runtime"
	//"fmt"
)

//type data struct {
//	x [1024 * 100]byte
//}
//func test() uintptr {
//	p := &data{}
//	return uintptr(unsafe.Pointer(p))
//}

/*
对象内存分配会受编译参数影响。举个例子子,当函数返回对象指针时,必然在堆上分配。
可如果该函数被内联,那么这个指针就不会跨栈帧使用用,就有可能直接在栈上分配,以实
现代码优化⺫目目的。因此,是否阻止止内联对指针输出结果有很大大影响。
除正常指针外,指针还有 unsafe.Pointer 和 uintptr 两种形态。其中 uintptr 被 GC 当
做普通整数对象,它不能阻止止所 "引用用" 对象被回收。


*/

type Data struct {
	d [1024 * 100]byte
	o *Data
}
func MyPointer(){
	//const N = 10000
	//cache := new([N]uintptr)
	////cyuyan
	//for i:=0;i<N;i++{
	//	cache[i] = test()
	//	time.Sleep(time.Millisecond)
	//}

	/*
	由于可以用用 unsafe.Pointer、uintptr 创建 "dangling pointer" 等非非法指针,所以在使
	用用时需要特别小小心心。另外,cgo C.malloc 等函数所返回指针,与 GC 无无关。
	指针构成的 "循环引用用" 加上 runtime.SetFinalizer 会导致内存泄露。
	*/
	//var a,b Data
	//a.o = &b
	//b.o = &a
	//runtime.SetFinalizer(&a, func(d *Data) { fmt.Printf("a %p final.\n", d) })
	//runtime.SetFinalizer(&b, func(d *Data) { fmt.Printf("b %p final.\n", d) })
	//time.Sleep(time.Millisecond)

	}