package ticketMaster

import (
	"sync/atomic"
	"sync"
	//"strconv"
	"errors"
	"net"
	"fmt"
	"os"
	"container/list"
	"strconv"
)

/*
写个管理资源的服务器，资源可以是火车票，选修课，或者钱之类的，
总之是有限的，不可重复分配的。要有持久层。再写个客户端，多放在几个机器上，
不停的请求资源。基本要求，服务器资源分配不能错，100个座，卖了250张票，是要拖出去打的。
在这个基础上，提高性能。代码里尝试不同的锁，cas，数据结构；设计上考虑缓存位置，如何与持久层同步；
网络上尝试不同的容器协议，对吞吐量的影响；架构上考虑达到单机上限后，如何水平扩展，如何宕机重启，如何负载平衡，如何发现服务...
*/

type ticket struct {
	ticketId int64
	ticketInfo	string
	}
type Manger struct{
	 //ticketMax atomic.
	 MaxTicketsNum int64
	 rwLock 	sync.RWMutex
	 sessionTickets *list.List
	 saledTicketsNum	int64
	 }

func NewManger(MaxSize int64) (*Manger,error){
	//返回的是元组
	manger := &Manger{
			MaxTicketsNum:	MaxSize,
			saledTicketsNum: 0,
			sessionTickets: list.New(),
	}
	manger.MaxTicketsNum =MaxSize
	for i:=0;i<int(manger.MaxTicketsNum);i++{
		tt:=&ticket{
			ticketId: int64(i),
			ticketInfo:"info",
		}

		manger.sessionTickets.PushBack(tt)
		//manger.sessionTickets<-tt

	}

	if manger.sessionTickets!=nil{
	return manger,nil}
	return nil,errors.New("新建manger失败")
	}
func (self*Manger)getTicket() (*ticket,error)  {
	self.rwLock.Lock()
	defer self.rwLock.Unlock()
	value:=atomic.LoadInt64(&self.saledTicketsNum)
	//fmt.Println("xxx")
	if value>=self.MaxTicketsNum{
		return nil,errors.New("sorry，票已经售空")
	}
	//fmt.Println("xxx")
	tt:=self.sessionTickets.Back()
	//fmt.Println(tt.Value.(*ticket).ticketId)
	self.sessionTickets.Remove(tt)
	atomic.AddInt64(&self.MaxTicketsNum,-1)
	return tt.Value.(*ticket),nil
	}

func checkErrserver(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
func TStart()  {
	manger,_:=NewManger(100)

	service := "127.0.0.1:5000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkErrserver(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkErrserver(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn,manger)
	}

	}

func handleClient(conn net.Conn,manger *Manger) {
	defer conn.Close()
	var buf [512]byte
	for {
		_, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		//rAddr := conn.RemoteAddr()
		//fmt.Println("Receive from client", rAddr.String(), string(buf[0:n]))
		//fmt.Println(string(buf[0:n]))
		ticket,merr := manger.getTicket()
		//fmt.Println(merr)
		//fmt.Println(ticket.ticketId)

		if merr==nil{
			backInfo:="票号"+strconv.Itoa(int(ticket.ticketId))+"info"+ticket.ticketInfo
			fmt.Println(backInfo)
			_, err2 := conn.Write([]byte(backInfo))
			if err2 != nil {
				return
			}
		}else {
			_, err2 := conn.Write([]byte("对不起 票已售完"))
			if err2 != nil {
				return
			}
		}

	}
}

func QueueTest()  {
	l:=list.New()
	l.PushBack(&ticket{
		ticketId: int64(11),
		ticketInfo:"info",
	})
	tt:=l.Front()
	fmt.Println(tt.Value.(*ticket).ticketId)

	}

