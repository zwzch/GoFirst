package tomcat
import ( "github.com/garyburd/redigo/redis"
	"fmt"
	"net"
	"time"
	"strings"
	"io"
	"sync"
	"github.com/pkg/errors"
	//"encoding/asn1"
	//"strconv"
	//"reflect"
)
func RedisTest()  {
	//c,err := redis.Dial("tcp","127.0.0.1:6379")
	//checkErr(err)
	//defer c.Close()

	//v,err:=c.Do("SET","name","red")
	//fmt.Println(v)
	//v,err = redis.String(c.Do("GET","name"))
	//fmt.Println(v)
	//

	//列表
	//c.Do("lpush","redlist","qqq")
	//c.Do("lpush","redlist","www")
	//c.Do("lpush","redlist","eee")
	//values, _ := redis.Values(c.Do("lrange", "redlist", "0", "100"))
	// for _, v := range values {
	//	     fmt.Println(string(v.([]byte)))
	//	 }
		go subscribe()
		go subscribe()
		go subscribe()
		go subscribe()

	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	 if err != nil {
		     fmt.Println(err)
		     return
		 }
	 defer c.Close()
	 for {
		     var s string
		     fmt.Scanln(&s)
		     _, err := c.Do("PUBLISH", "redChatRoom", s)
		     if err != nil {
			         fmt.Println("pub err: ", err)
			         return
			     }
		 }
	}
	func subscribe(){
		c,err:=redis.Dial("tcp","127.0.0.1:6379")
		checkErr(err)
		defer c.Close()
		psc:= redis.PubSubConn{c}
		psc.Subscribe("redChatRoom")
		for {
			switch v:=psc.Receive().(type) {
			 case redis.Message:
				             fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
			 case redis.Subscription:
				 fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
			 case error:
				 fmt.Println(v)
				 return
			}
		}
	}

func Acceept(listener net.Listener)(net.Conn ,error){
	var tempDelay time.Duration
	for{
		conn,err :=listener.Accept()
		if err !=nil{
			if ne, ok := err.(net.Error); ok && ne.Temporary(){
				if tempDelay == 0{
					tempDelay = 5 * time.Millisecond
				}else {
					tempDelay *= 2
				}
				if max:=1 * time.Second;tempDelay>max{
					tempDelay  = max
				}
				time.Sleep(tempDelay)
				continue
			}
			if strings.Contains(err.Error(),"use of closed network connection"){
				return nil,io.EOF
			}
			return nil,err

		}
		return conn,nil
	}
}



//频繁的创建和关闭连接，对系统会造成很大负担
//所以我们需要一个池子，里面事先创建好固定数量的连接资源，需要时就取，不需要就放回池中。
//但是连接资源有一个特点，我们无法保证连接长时间会有效。
//比如，网络原因，人为原因等都会导致连接失效。
//所以我们设置一个超时时间，如果连接时间与当前时间相差超过超时时间，那么就关闭连接。

//只要类型实现了ConnRes接口中的方法，就认为是一个连接资源类型
type ConnRes interface {
	Close() error
}

type Factory func()(ConnRes,error)

type Conn struct{
	conn ConnRes
	time time.Time
}

type ConnPool struct{
	mu sync.Mutex
	conns chan*Conn
	Factory Factory
	closed bool
	connTimeOut time.Duration
}


func NewConnPool(factory Factory,cap int,connTimeOut time.Duration)(*ConnPool, error){
	if cap<=0{
		return nil,errors.New("cap不能小于0")
	}
	if connTimeOut <=0 {
		return nil,errors.New("timeout不能小于0")
	}
	cp:=&ConnPool{
		mu:	sync.Mutex{},
		conns:	make(chan*Conn,cap),
		Factory: factory,
		closed:false,
		connTimeOut:connTimeOut,
	}
	for i:=0;i<cap;i++{
		connRes ,err :=cp.Factory();
		if err!=nil {
			return nil,errors.New("factory出错")
		}
		cp.conns<-&Conn{conn: connRes, time: time.Now()};
	}
	return cp,nil
}
func (cp *ConnPool) Get()(ConnRes,error){
	if cp.closed{
		return nil,errors.New("连接池已关闭")
	}
	for{
		select{
		case connRes,ok:= <-cp.conns:
			{
				if !ok{
					return nil,errors.New("连接池已关闭")
				}

				//判断连接中的时间，如果超时，则关闭
				//继续获取
				if time.Now().Sub(connRes.time) > cp.connTimeOut {
					connRes.conn.Close()
					continue
				}
				return connRes.conn,nil
			}
		default:
			{
				//新创建资源返回
				connRes, err := cp.Factory()
				if err != nil {
					return nil, err
				}
				return connRes, nil
			}

		}
	}
}

func (cp *ConnPool)Put(conn ConnRes)error{
	if cp.closed {
		return errors.New("连接池已关闭")
	}

	select {
	//向通道中加入连接资源
	case cp.conns <- &Conn{conn: conn, time: time.Now()}:
		{
			return nil
		}
	default:
		{
			//如果无法加入，则关闭连接
			conn.Close()
			return errors.New("连接池已满")
		}
	}
}
func (cp *ConnPool) Close(){
	if cp.closed{
		return
	}
	cp.mu.Lock()
	cp.closed = true
	close(cp.conns)
	for conn :=range cp.conns{
		conn.conn.Close()
	}
	cp.mu.Unlock()
}
func (cp *ConnPool) len() int {
	return len(cp.conns)
}
func ConnTest(){
	cp,_:=NewConnPool(func() (ConnRes, error) {
		return redis.Dial("tcp","127.0.0.1:6379")
		},10,time.Second*10)
	//
	conn1,_:=cp.Get()

	//fmt.Println("cp len : ", cp.len())
	//v,_:=conn1.(redis.Conn).Do("SET","name","red")
	//fmt.Println(v)
	//v,err = redis.String(c.Do("GET","name"))
	//fmt.Println(v)
	//if RedisExist(conn1.(redis.Conn),"aaa") ==1 {
	//	fmt.Println("true")
	//}

	//RedisPut(conn1.(redis.Conn),"heihei","niangmen")
	fmt.Println(RedisGet(conn1.(redis.Conn),"heihei"))
	}




	func RedisExist(conn redis.Conn,key string)int64{
		flag,_:=conn.Do("Exists",key)
		//flag ,_= strconv.Atoi(flag)
		//fmt.Println(flag)
		//fmt.Println(reflect.TypeOf(flag))
		return flag.(int64)
	}

	func RedisGet(conn redis.Conn,key string)  string{
		v,err := redis.String(conn.Do("GET",key))
		//fmt.Println(v)
		checkErr(err)
		return v
	}
	func RedisPut(conn redis.Conn,key string,value string){
		_,err:=conn.(redis.Conn).Do("SET",key,value)
		checkErr(err)
	}


