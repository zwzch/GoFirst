package myGoroutine

import (
	"fmt"
	"time"
	"net"
	"log"
	"io"
	"os"

)
func check(err error){
	if err != nil {
		log.Fatal(err)
	}
}
/*
考虑一下，一个函数在线性程序中可以正确地工作。如果在并发的情况下，这个函数依然可以正确地工作的话，
那么我们就说这个函数是并发安全的，并发安全的函数不需要额外的同步工作。我们可以把这个概念概括为一个特定类型的一些方法和操作函数，
对于某个类型来说，如果其所有可访问的方法和操作都是并发安全的话，那么类型便是并发安全的。
，比如死锁(deadlock)、活锁(livelock)和饿死(resource starvation)。我们没有空去讨论所有的问题，这里我们只聚焦在竞争条件上。
*/
func xxx()  {
	fmt.Println("xxx")
}
func MyGoroutine2()  {
	//go xxx()
	//time.Sleep(time.Millisecond)
	//go spinner(100 * time.Millisecond)
	//const n = 45
	//fibN := fib(n) // slow
	//fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)


	/*
	然后主函数返回。主函数返回时，所有的goroutine都会被直接打断，程序退出。除了从主函数退出或者直接终止程序之外，
	没有其它的编程方法能够让一个goroutine来打断另一个的执行，但是之后可以看到一种方式来实现这个目的，
	通过goroutine之间的通信来让一个goroutine请求其它的goroutine，并让被请求的goroutine自行结束执行。
	*/
	/*
	网络编程是并发大显身手的一个领域，由于服务器是最典型的需要同时处理很多连接的程序，这些连接一般来自于彼此独立的客户端。
	在本小节中，我们会讲解go语言的net包，这个包提供编写一个网络客户端或者服务器程序的基本组件，无论两者间通信是使用TCP，
	UDP或者Unix domain sockets。在第一章中我们使用过的net/http包里的方法，也算是net包的一部分。
	*/


	/*
	Listen函数创建了一个net.Listener的对象，这个对象会监听一个网络端口上到来的连接，在这个例子里我们用的是TCP的localhost:8000端口。listener对象的Accept方法会直接阻塞，\
	直到一个新的连接被创建，然后会返回一个net.Conn对象来表示这个连接。
	handleConn函数会处理一个完整的客户端连接。在一个for死循环中，用time.Now()获取当前时刻，然后写到客户端。由于net.Conn实现了io.Writer接口，我们可以直接向其写入内容。
	这个死循环会一直执行，直到写入失败。最可能的原因是客户端主动断开连接。这种情况下handleConn函数会用defer调用关闭服务器侧的连接，然后返回到主函数，继续等待下一个连接请求。


	*/

	tcpClient()
	//tcpServer()
	}
func tcpServer()  {
	listenser,err := net.Listen("tcp","localhost:8000")
	check(err)
	for {
		conn,err := listenser.Accept()
		check(err)
		go 	handleConn(conn)
	}
}
	func tcpClient(){
		conn, err := net.Dial("tcp", "localhost:8000")
		check(err)
		defer conn.Close()
		mustCopy(os.Stdout,conn)
	}
func mustCopy(dst io.Writer,src io.Reader){
	_,err  := io.Copy(dst, src)
	//writer reader
	check(err)

	}
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
func handleConn(c net.Conn){
	defer c.Close()
	for{
		_,err :=io.WriteString(c,time.Now().Format("15:04:05\n"))
		check(err)
		time.Sleep(1*time.Second)
	}
}
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}