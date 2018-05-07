package http

import (
	//"fmt"
	"net/http"
	"io"
	//"unicode"
	"log"
	"fmt"
	//"unicode"
	"time"
)

func check(e error)  {
	if e != nil{
		panic(e)
	}
}

func MyHttp()  {
	/*
		http
	*/

	//myHttp1()
	//myhttp2()
	//httpFirst()
	/*
	接收request的过程中，最重要的莫过于路由（router），即实现一个Multiplexer器。
	Go中既可以使用内置的mutilplexer --- DefautServeMux，也可以自定义。
	Multiplexer路由的目的就是为了找到处理器函数（handler），后者将对request进行处理，同时构建response。
	*/
	//Clinet -> Requests ->  [Multiplexer(router) -> handler  -> Response -> Clinet
	//路由分配handler
	/*
		因此，理解go中的http服务，最重要就是要理解Multiplexer和handler，Golang中的Multiplexer基于ServeMux结构，
		同时也实现了Handler接口。
	*/
	/*
		hander函数： 具有func(w http.ResponseWriter, r *http.Requests)签名的函数
		handler处理器(函数): 经过HandlerFunc结构包装的handler函数，它实现了ServeHTTP接口方法的函数。调用handler处理器的ServeHTTP方法时，即调用handler函数本身。
		handler对象：实现了Handler接口ServeHTTP方法的结构
	*/

	//myhttp2()

	//myHttp3()
	//myHttp4()
	myHttp5()
	}

func httpFirst(){
	xxx :=func(w http.ResponseWriter,r *http.Request) {
		fmt.Println("helloworld")
		io.WriteString(w,"Hello world!")
		}
	http.HandleFunc("/", xxx)
	http.ListenAndServe("127.0.0.1:8080",nil)
}


func myHttp1(){
	//类似与C 传入函数处理 request和response有封装的过程
	sayhello := func(resp http.ResponseWriter,req *http.Request ) {
		fmt.Println("爷来咯")
		io.WriteString(resp,"Hello world!")
	}
	http.HandleFunc("/",sayhello)
	err:=http.ListenAndServe(":8080",nil)
	if err!=nil{
		log.Fatal(err)
	}

	}

type myHandler struct {

}

func myhttp2(){
	mux:=http.NewServeMux()
	mux.Handle("/",&myHandler{})
	err:=http.ListenAndServe(":8080",mux)
	check(err)
}
func (this *myHandler)ServeHTTP(resp http.ResponseWriter,req *http.Request)  {
	io.WriteString(resp,"Hello world!")
}

type myHandler3 struct {

}
func (this *myHandler3)ServeHTTP(resp http.ResponseWriter,req *http.Request)  {
	io.WriteString(resp,"Hello world!")
}
func sayhello21(resp http.ResponseWriter,req *http.Request){
	io.WriteString(resp,"Hello world! 21")
}


func myHttp3()  {
	mux := http.NewServeMux()
	//http路由
	mux.Handle("/",&myHandler3{})
	//Handle 传入struct
	mux.HandleFunc("/hello",sayhello21)
	//HandleFunc传入函数
	err:=http.ListenAndServe(":8080",mux)
	if err != nil{
		log.Fatal(err)
	}
}



type myHandler22 struct {

}

func (this *myHandler22)ServeHTTP(resp http.ResponseWriter,req *http.Request)  {
	io.WriteString(resp,"Hello world!")
}

func sayhello22(resp http.ResponseWriter,req *http.Request){
	io.WriteString(resp,"Hello world! 21")
}
func myHttp4(){
	mux:=http.NewServeMux()
	mux.Handle("/",&myHandler22{})
	mux.HandleFunc("/hello",sayhello22)
	wd:="/home/zw/go/src/GoFirst/myio/static"
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(wd))))
	err := http.ListenAndServe(":8080",mux)
	check(err)
}


type myHandler5 struct {

}

func (this *myHandler5)ServeHTTP(resp http.ResponseWriter,req *http.Request)  {
	io.WriteString(resp,"Hello world!")
}

func myHttp5(){
	server := http.Server{
		Addr:":8080",
		Handler:&myHandler5{},
		ReadTimeout:5*time.Second,
		}
		err:=server.ListenAndServe()
		check(err)
}



/*

	通过map配置参数
*/
var mux map[string]func(http.ResponseWriter,*http.Request)
func MyHttp31() {
	server:=http.Server{
		Addr:":8080",
		Handler:&myHandler31{},
		ReadTimeout:5*time.Second,
	}

	mux=make(map[string]func(http.ResponseWriter,*http.Request))
	mux["/hello"]=sayhello31

	err:=server.ListenAndServe()

	if err!=nil{
		log.Fatal(err)
	}
}

type myHandler31 struct {

}

func (this *myHandler31)ServeHTTP(resp http.ResponseWriter,req *http.Request)  {
	if h,ok:=mux[req.URL.String()];ok{
		h(resp,req)
		return
	}
}

func sayhello31(resp http.ResponseWriter,req *http.Request){
	io.WriteString(resp,"Hello world! 31")
}
