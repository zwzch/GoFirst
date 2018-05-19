package myGoroutine

import (
	"sync/atomic"
	"runtime"
	"time"
	"fmt"
	//"sync"
	"sync"
	//"strconv"
)

func MyAtomic(){
	//CAS

	var ops uint64  = 0
	for i:=0;i<50;i++{
		go func() {
			//ops=ops+1
			//for {
				atomic.AddUint64(&ops,1)
			//	// 这个函数用于时间片切换
			//	//可以理解为高级版的time.Sleep()
			//	//避免前面的for循环将CPU时间片都卡在一个线程里，使得其它线程没有执行机会
				runtime.Gosched()
			//	}
		}()
	}
	time.Sleep(time.Second)
	opsFinal :=atomic.LoadUint64(&ops)
	fmt.Println("ops",opsFinal)
}


func MyMutx(){
	var mutex sync.Mutex
	count:=0
	for r:=0;r<50;r++{
		go func() {
			mutex.Lock()
			count = count+2
			mutex.Unlock()
			runtime.Gosched()
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(count)
}

func MyOnce(){
	o:=&sync.Once{}
	go do(o)
	go do(o)
	time.Sleep(time.Second*2)
	}
	func do(o *sync.Once){
		fmt.Println("Start do")

		o.Do(func() {

			fmt.Println("Doing something...")

		})
		/*
			once.do里的方法只会被执行一次
		*/

		func(){
			fmt.Println("not once")
		}()

		fmt.Println("Do end")
	}




type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	//如果automic变量是1 直接返回
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}

	o.m.Lock()
	defer o.m.Unlock()

	if o.done == 0 {
		f()
		atomic.StoreUint32(&o.done, 1)
	}
}

type Once2 struct {
	done int32
}

func (o *Once2) Do2(f func()) {
	if atomic.LoadInt32(&o.done) == 1 {
		return
	}
	// Slow-path.
	if atomic.CompareAndSwapInt32(&o.done, 0, 1) {
		f()
	}
}


func MyOnce2() {

	o := &Once{}

	go do2(o)

	go do2(o)

	time.Sleep(time.Second * 2)

	fmt.Println("========================1=======================")
	//========================================================================

	o3 := &Once2{}

	go do3(o3)

	go do3(o3)

	time.Sleep(time.Second * 2)

}

func do2(o *Once) {

	fmt.Println("Once Start do")

	o.Do(func() {

		fmt.Println("Once Doing something...")

	})

	func(){
		fmt.Println("Once not once")
	}()

	fmt.Println("Once Do end")

}

func do3(o *Once2) {

	fmt.Println("Once2 Start do")

	o.Do2(func() {

		fmt.Println("Once2 Doing something...")

	})

	func(){
		fmt.Println("Once2 not once")
	}()

	fmt.Println("Once2 Do end")

}




var m *sync.RWMutex
//读写锁

func MyRWLock(){
	m = new(sync.RWMutex)
	go read(1)
	go read(2)

	go write(3)
	go read(4)
	go write(5)
	time.Sleep(time.Second*4)
}
func read(i int){
	fmt.Println(i,"read start")
	m.RLock()
	//
	fmt.Println(i,"reading")
	time.Sleep(time.Second*1)
	m.RUnlock()
	fmt.Println(i, "read end")
}

func write(i int)  {

	fmt.Println(i, "write start")



	m.Lock()

	fmt.Println(i, "writing")

	time.Sleep(1 * time.Second)

	m.Unlock()



	fmt.Println(i, "write end")
}

func MyCountDownLunch(){
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i:=0;i<10;i++{
		go func(i int) {
			time.Sleep(time.Second)
			fmt.Println("ok",i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("轮到我上场了")
}
func MySelect(){
	//c1:=make(chan int)
	//go func(){
	//	for v:=range c1{
	//		fmt.Println(v)
	//	}
	//	}()
	//
	//for i:=0;i<5;i++{
	//
	//	select {
	//		//随机选择一个
	//	case c1<-0:
	//		//fmt.Println(111)
	//	case c1<-1:
	//		//fmt.Println(222)
	//
	//		}
	//}

	//go func() {
	//	for i:=0;i<5;i++{
	//		c1<-i
	//	}}()
	//for  {
	//
	//
	//select {
	//	//只有执行一次 没有循环
	//case  j:= <-c1:
	//	fmt.Println("xxx",j)
	//case i:=<-c1:
	//	fmt.Println("lll",i)
	//
	//}}

	//c11:=make(chan int)
	//c12:=make(chan int)
	//go func() {
	//	c11<-1
	//}()
	//go func() {
	//	c12<-2
	//}()
	//for i:=0;i<2;i++{
	//	select{
	//	case j:=<-c11:
	//		fmt.Println("c11: "+strconv.Itoa(j))
	//	case j:=<-c12:
	//		fmt.Println("c12: "+strconv.Itoa(j))
	//	}
	//}


	//使用default避免死锁错误
	//a:=make(chan string)
	//select{
	//case str:=<-a:
	//	fmt.Println(str)
	//default:
	//	fmt.Println("非阻塞") //否则就是死锁错误
	//}
	//select{
	//case a<-"xxx":
	//	fmt.Println("发送了xxx")
	//default:
	//	fmt.Println("非阻塞") //否则就是死锁错误
	//}

	//
	//runtime.GOMAXPROCS(runtime.NumCPU()) //多核处理
	//c1:=make(chan bool,10)
	//for i:=0;i<10;i++{
	//	go func(i int) {
	//		sum:=0
	//		for j:=0;j<1000000;j++{
	//			sum+=j
	//		}
	//		fmt.Println(i)
	//		c1<-true
	//	}(i)
	//}
	//
	//for i:=0;i<10;i++{
	//	<-c1
	//}



	//ticker := time.NewTicker(time.Millisecond * 500)
	//
	//go func() {
	//	for t := range ticker.C {
	//		fmt.Println("Tick at", t)
	//	}
	//}()
	//time.Sleep(time.Millisecond * 1600)
	//ticker.Stop()
	//fmt.Println("Ticker stopped")

	}

