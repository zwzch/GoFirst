package myfunc

import (
	"strconv"
	"fmt"
	"os"
	"errors"
)

func Myfunc() {
	//res1,_:=myfun1(1,2,"3")
	//fmt.Println(res1)
	//res2,_:=myfun2(4,5,"6")
	//fmt.Println(res2)
	//res3,_:=myfun3("6",1,2,3,4,5)
	//fmt.Println(res3)
	//a:=1
	//b:=2
	//c:=[]int{1,2}
	//myfun31(c,a,b)
	//fmt.Println(a,b,c)

	//var add = func(a,b int)(int){
	//	return  a+b
	//}
	//fmt.Println(add(1,2))
	//var opt= func(a,b int,myfunc func(int,int)(int))(int) { //参数是函数
	//	return myfunc(a,b)
	//}
	//fmt.Println(opt(1,2,add))

	//var cl = func(x int) (func(int) (int)) { //返回值是函数，典型案例，闭包
	//	return func(y int) (int) {
	//		x++
	//		return x + y
	//	}
	//}
	//mycl1:=cl(1)
	//m1:=mycl1(1)
	//m2:=mycl1(1)
	//mycl2:=cl(1)
	//m3:=mycl2(1)
	//fmt.Println(m1,m2,m3)







	/*
		myfunc!!!!
		函数是第一一类对象,可作为参数传递。建议将复杂签名定义为函数类型,以便于阅读。
		变参本质上就是 slice。只能有一一个,且必须是最后一一个。
	*/
	//
	//s:=[]int{1,2,3}
	//fmt.Println(test("sum: %d",s...))

	//不能用用容器对象接收多返回值。只能用用多个变量,或 "_" 忽略。
	//x,_ := test2()
	//fmt.Println(x)
	//多返回值可直接作为其他函数调用用实参。
	//命名返回参数可看做与形参类似的局部变量,最后由 return 隐式返回。
	//fmt.Println(add(1,2))
	//命名返回参数允许 defer 延迟调用用通过闭包读取和修改。根据defer的栈顺序
	//最上面的最后执行
	//fmt.Println(add2(1,2))

	// --- function variable ---
	//匿名函数
	//fn:= func() {fmt.Println("Hello World")}
	//fn()
	// --- function collection ---
	//fns := [](func(x int) int){
	//	func(x int) int { return x + 1 },
	//	func(x int) int { return x + 2 },
	//}
	//fmt.Println(fns[0][100])

	//fc :=make(chan func() string , 2)
	//fc <-func() string{
	//	return "Hello World"
	//}
	//println((<-fc)())

	//闭包复制的是原对象指针,这就很容易解释延迟引用用现象。
	//test3()()

	/*
	关键字 defer 用用于注册延迟调用用。这些调用用直到 ret 前才被执行行,通常用用于释放资源或错
	误处理。
	*/
	//test4()
	//多个 defer 注册,按 FILO 次序执行行。哪怕函数或某个延迟调用用发生生错误,这些调用用依旧
	//会被执行行。
	//test5(0)

	//test
	//滥用用 defer 可能会导致性能问题,尤其是在一一个 "大大循环" 里里。
	//test7()
	//test8()

	//test9(1,2)
	//funcb22()


	//switch z,err:=div(10,0);err {
	//case nil:
	//	print(z)
	//case ErrDivByZero:
	//	panic(err)
	//}
	//?
	//如何区别使用用 panic 和 error 两种方方式?
	//惯例是:导致关键流程出现不可修复性错误的
	//使用用 panic,其他使用用 error。


	}
var ErrDivByZero = errors.New("division by zero")
	func div(x, y int) (int, error) {
		if y == 0 { return 0, ErrDivByZero }
	return x / y, nil
	}
	type error interface{
		Error() string
	}
	func funcb22(){
		defer func() {
			err:=recover()
			if err !=nil{
				fmt.Println(err)
			}
		}()
		panic("B error")
	}
func test9(x,y int){
	var z int
	func (){
		defer func() {
			if recover() != nil{
				z =0
			}
		}()
		z = x / y
		return
	}()
	println("x/y=",z)
}
func test8() {
	defer func() {
		fmt.Println(recover())
	}()
	defer func() {
		panic("defer panic")
	}()
	panic("test panic")
}

func test7(){
	defer func(){
		if err:= recover();err!=nil{
			println(err.(string))
		}
	}()
	panic("panic error")
}

func test6() {
	x, y := 10, 20
	defer func(i int) {
		println("defer:", i, y)
		// y 闭包引用用

	}(x)
	// x 被复制
	x += 10
	y += 100
	println("x =", x, "y =", y)
}
func test5(x int) {
	defer println("a")
	defer println("b")
	defer func() {
		println(100 / x)
		// div0 异常未被捕获,逐步往外传递,最终终止止进程。
	}()
	defer println("c")
}
	func test4() error  {
		f,err :=os.Create("test.txt")
		if err!=nil{return err}
		defer f.Close()// 注册调用用,而而不是注册函数。必须提供参数,哪怕为空。类似于final
		f.WriteString("Hello World")
		return nil


	}
	func test3() func(){
		x:=100
		fmt.Printf("x (%p) = %d\n", &x, x)
		return func() {
			fmt.Printf("x (%p) = %d\n", &x, x)
		}
	}
	func add2(x,y int)(z int){
		//(z = x + y) -> (call defer) -> (ret)
		defer func() {
			z-=50
		}()
		defer func() {
			z+=100

		}()
		z=x+y
		return
	}
	func add(x,y int)(z int){
		z = x+y
		return
	}
	func test2()(int,int){
		return 1,2
	}
	func test(s string, n ...int)string {
	var x int
	for _, i := range n {
		x += i
	}
	return fmt.Sprint(s,x)
}
//func test(x,y int, s string)(int, string)  {
//	n := x+y
//	return n,fmt.Sprint(s,n)
//}
func myfun1(a,b int, c string) (int,error) {
	d,err:=strconv.Atoi(c)
	return a+b+d,err
}
func myfun2(a,b int,c string)(e int,err error){ //私有函数
	d,err:=strconv.Atoi(c)
	e= a+b+d
	//将返回值也提前定义
	return
}
func myfun3(c string,a... int)(e int,err error){ //私有函数
	d,err:=strconv.Atoi(c)
	for _,v:=range a{
		d=d+v
	}
	e=d
	return
}

func myfun31(c []int,a... int){ //私有函数
	c[0]=123
	a[0]=123
}
