package tomcat

import (
	"sync"
	"sync/atomic"
	"errors"
)

const  sessionMapNum  = 32
var SessionClosedError = errors.New("Session Closed")
var SessionBlockedError = errors.New("Session Blocked")

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

func (manager *Manager) GetSession(sessionID uint64) *Session {
	smap := &manager.sessionMaps[sessionID%sessionMapNum]
	smap.RLock()
	defer smap.RUnlock()
	session, _ := smap.sessions[sessionID]
	return session
}

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
	func (session *Session) IsClosed() bool{
		return atomic.LoadInt32(&session.closeFlag) == 1
	}
	func (session *Session)Close() error{
		if atomic.CompareAndSwapInt32(&session.closeFlag,0,1){
			close(session.closeChan)
			if session.sendChan!=nil{
				session.sendMutex.Lock()
				close(session.sendChan)
				//if clear, ok := session.codec.(ClearSendChan); ok {
				//	clear.ClearSendChan(session.sendChan)
				//}

				session.sendMutex.Unlock()
			}

		}
		return SessionClosedError
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







