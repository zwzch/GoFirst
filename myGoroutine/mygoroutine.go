package myGoroutine

import (
	"math"
	//"sync"
	"time"
	"fmt"
	//"runtime"
	//"os"
	//"unicode"
	"math/rand"

	//"sync"
)

/*Goroutine*/
/*
	Go 在语言言层面面对并发编程提供支支持,一一种类似协程,称作 goroutine 的机制。

	只需在函数调用用语句前添加 go 关键字,就可创建并发执行行单元。开发人人员无无需了解任何
	执行行细节,调度器会自自动将其安排到合适的系统线程上执行行。goroutine 是一一种非非常轻量
	级的实现,可在单个进程里里执行行成千上万的并发任务。

	事实上,入入口口函数 main 就以 goroutine 运行。另有与之配套的 channel 类型,用用以实
	现 "以通讯来共享内存" 的 CSP 模式。相关实现细节可参考本书第二部分的源码剖析。
*/
func sum(id int){
	var x  int64
	for i := 0; i < math.MaxUint32; i++ {
		x += int64(i)
	}
	println(id, x)
}
func MyGoroutine()  {
	//print("xxx")
	//go func() {
	//	print("HelloWorld")
	//}()
	/*
		调度器不能保证多个 goroutine 执行行次序,且进程退出时不会等待它们结束。
		默认情况下,进程启动后仅允许一个系统线程服务于 goroutine。可使用用环境变量或标准
		库函数 runtime.GOMAXPROCS 修改,让调度器用用多个线程实现多核并行行,而而不仅仅是
		并发。
	*/
	//wg := new(sync.WaitGroup)
	//wg.Add(2)
	//for i:=0;i<2;i++{
	//	go func(id int) {
	//		defer wg.Done()
	//		sum(id)
	//	}(i)
	//}
	//wg.Wait()

	/*
	WaitGroup总共有三个方法：Add(delta int),Done(),Wait()。简单的说一下这三个方法的作用。
    Add:添加或者减少等待goroutine的数量
    Done:相当于Add(-1)
    Wait:执行阻塞，直到所有的WaitGroup数量变成0
	类似于协程池
	*/
	//var wg sync.WaitGroup
	//for i:=0;i>5;i=i+1{
	//	wg.Add(1)
	//	go func(n int) {
	//		defer wg.Add(-1)
	//		EchoNumber(n)
	//	}(i)
	//}
	//wg.Wait()
	//调用 runtime.Goexit 将立即终止止当前 goroutine 执行行,调度器确保所有已注册 defer
	//延迟调用用被执行。

	//wg :=new(sync.WaitGroup)
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	defer print("A.defer")
	//	func() {
	//		defer fmt.Println("B.defer")
	//		runtime.Goexit()
	//		//fmt.Println("B")
	//	}()
	//	//fmt.Println("A")
	//}()
	//wg.Wait()

	//和协程 yield 作用用类似,Gosched 让出底层线程,将当前 goroutine 暂停,放回队列等
	//待下次被调度执行行。
	//wg := new (sync.WaitGroup)
	//wg.Add(2)
	//go func() {
	//	//fmt.Println("first1")
	//	defer wg.Done()
	//	for i:=0;i<6;i++{
	//		fmt.Println(i)
	//		if i==3{
	//			runtime.Gosched()
	//		}
	//	}
	//}()
	//
	//go func() {
	//	//fmt.Println("first2")
	//	defer wg.Done()
	//	println("Hello, World!")
	//}()
	//
	//wg.Wait()

	/*
		CHANNEL
		引用用类型 channel 是 CSP 模式的具体实现,用用于多个 goroutine 通讯。其内部实现了
		同步,确保并发安全。
		默认为同步模式,需要发送和接收配对。否则会被阻塞,直到另一一方方准备好后被唤醒。
	*/

	//data:=make(chan int)
	//exit := make(chan bool)
	//go func() {
	//	for d := range data{
	//		fmt.Print(d)
	//	}
	//	fmt.Print("recv over")
	//	exit <- true
	//}()
	//
	//data<-1
	//data<-2
	//data<-3
	//close(data)
	////关闭队列
	//fmt.Print("send over")
	//fmt.Println(<-exit)

	//data := make(chan int,3)
	//exit := make(chan bool)
	//data<-1
	//data<-2
	//data<-3
	//go func() {
	//	for d:= range data{// 在缓冲区未空前,不会阻塞。
	//		fmt.Print(d)
	//	}
	//	exit<-true
	//}()
	//data<-4
	//data<-5
	//close(data)
	//<-exit

	//var a, b chan int = make(chan int), make(chan int, 3)
	//for {
	//	if d, ok := <-data; ok {
	//		fmt.Println(d)
	//	} else {
	//		break
	//	}
	//}
	//除用用 range 外,还可用用 ok-idiom 模式判断 channel 是否关闭。
//	向 closed channel 发送数据引发 panic 错误,接收立立即返回零值。而而 nil channel,无无
//	论收发都会被阻塞。
	//内置函数 len 返回未被读取的缓冲元素数量,cap 返回缓冲区大大小小。
	//d1 := make(chan int)
	//d2:=make(chan int,3) 可以设置容量大小
	//d2<-1
	//fmt.Println(len(d1), cap(d1))
	//fmt.Println(len(d2), cap(d2))
	//单向
	//可以将 channel 隐式转换为单向队列,只收或只发。
	//c := make(chan int, 3)
	//var send chan<- int  = c
	//var recv <-chan int = c
	//send <- 1
	//fmt.Print(<-recv)
	/*
		选择channel
		如果需要同时处理多个 channel,可使用用 select 语句。它随机选择一一个可用用 channel 做
		收发操作,或执行行 default case。
	*/
	//a,b := make(chan  int,3),make(chan int)
	//go func() {
	//	v,ok,s :=0,false,""
	//	for {
	//		select {
	//		// 随机选择可用用 channel,接收数据。
	//		case v, ok = <-a: s = "a"
	//		case v, ok = <-b: s = "b"
	//		}
	//		if ok {
	//			fmt.Println(s, v)
	//		} else {
	//			os.Exit(0)
	//		}
	//	}
	//}()
	//
	//for i:=0;i<5;i++{
	//	select {
	//	case a<-i:
	//	case b<-i:
	//	}
	//}

	//用简单工厂模式打包并发任务和 channel。
	//fmt.Print(<-NewTest())
	//用用 channel 实现信号量 (semaphore)。

	//wg:=sync.WaitGroup{}
	//wg.Add(3)
	//sem := make(chan int, 1)
	//for i := 0; i < 3; i++ {
	//	go func(id int) {
	//		defer wg.Done()
	//		sem <- 1
	//		// 向 sem 发送数据,阻塞或者成功。
	//		for x := 0; x < 3; x++ {
	//			fmt.Println(id, x)
	//		}
	//		<-sem
	//		// 接收数据,使得其他阻塞 goroutine 可以发送数据。
	//	}(i)
	//}
	//wg.Wait()


	//var wg sync.WaitGroup
	//quit := make(chan bool)
	//for i:=0;i<2;i++{
	//	wg.Add(1)
	//	go func(id int) {
	//		defer wg.Done()
	//		task := func() {
	//			fmt.Println(id,time.Now().Nanosecond())
	//			time.Sleep(time.Second)
	//		}
	//		for{
	//			select {
	//			case <-quit:
	//				return
	//			default:
	//				task()
	//
	//			}
	//		}
	//	}(i)

	//}
	//time.Sleep(time.Second*5)
	//close(quit)
	//wg.Wait()

	//用 select 实现超时 (timeout)。
	//w:=make(chan bool)
	//c:=make(chan int,2)
	//go func() {
	//	select {
	//	case v:=<-c:fmt.Println(v)
	//	case <-time.After(time.Second*3):fmt.Println("timeout.")
	//	}
	//	w<-true
	//}()
	//<-w
	//c<-2
	//time.After()?
	//
	//tchan := time.After(time.Second*3)
	//fmt.Printf("tchan type=%T\n",tchan)
	//fmt.Println("mark 1")
	//fmt.Println("tchan=",<-tchan)
	//fmt.Println("mark 2")

	//channel 是第一一类对象,可传参 (内部实现为指针) 或者作为结构成员。
	//process := func(req *Request) {
	//	x := 0
	//	for _, i := range req.data {
	//		x += i
	//	}
	//	req.ret <- x
	//}
	//req := NewRequest(10,20,30)
	//process(req)
	//fmt.Println(<-req.ret)
	}
type Request struct {
	data []int
	ret chan int
}
func NewRequest(data ...int) *Request {
	return &Request{ data, make(chan int, 1) }
}
func NewTest() chan int{
	c:=make(chan int)
	rand.Seed(time.Now().UnixNano())
	go func(){
		fmt.Println("f===========")
		time.Sleep(time.Second)
		fmt.Println("m===========")
		c <- rand.Int()
		fmt.Println("l===========")
	}()
	return c
	}

func EchoNumber(i int) {
	time.Sleep(3e9)
	fmt.Println(i)
}