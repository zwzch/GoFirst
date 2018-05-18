package tomcat

import (
	"sync"
	"sync/atomic"
)

const  sessionMapNum  = 32

var globalSessionId uint64
type Manager struct {
	sessionMaps [sessionMapNum]sessionMap
	//sessionmaps 储存多个sessionMap
	disposeOnce  sync.Once
	disponseWait	sync.WaitGroup
	}
type sessionMap struct {
	//sessionMap储存多个session
	sync.RWMutex
	sessions map[uint64]*Session
	disposed bool
}
func NewManger() *Manager{
		manager := &Manager{}
		//给sessionMaps初始化内存
		for i :=0;i<len(manager.sessionMaps);i++{
			manager.sessionMaps[i].sessions =  make(map[uint64]*Session)

		}
		return manager}

type Session struct {
	id uint64
	manger *Manager
	sendChan chan interface{}
	recvMutex	sync.Mutex
	sendMutex 	sync.RWMutex

	closeFlag	int32
	closeChan	chan int
	closeMutex	sync.Mutex
	State interface{}
	}

	func (session *Session) ID() uint64{
		return session.id
	}
	func (session *Session) IsClosed(){

	}
func NewSession(sessionChanSize int) *Session{
	return newSession(nil,sessionChanSize)
	}
func newSession(manager *Manager,sendChanSize int) *Session{
	session:=&Session{
		manger:manager,
		closeChan:make(chan int),
		id:        atomic.AddUint64(&globalSessionId, 1),//session全局id

	}
	if sendChanSize>0{
		session.sendChan = make(chan interface{}, sendChanSize)

	}
	return session
	}




