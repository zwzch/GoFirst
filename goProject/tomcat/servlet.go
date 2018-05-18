package tomcat

import (
	//"net/http"
	"net"
	"fmt"
	//"strings"
	"strings"
	"strconv"
)
type responseImp interface {
	setMethod()
	setParameter()
	printToWeb(sth string)
}
type requestImp interface {

	//uri string
	getParameter()
	getUri()
	parseInput()}



	//req
	type Request struct {
		requestImp
		uri string
		serverport int
		httpversion string
		method string
		input_all string
		parmetersmap map[string]string
		}

func (self *Request)getUri()  string{
	return self.uri
}
func (self *Request)parseInput(inputs []string){
	//tt:= [...]string
	self.parmetersmap = make(map[string]string)
	for _,element :=range inputs[1:]{
		if (element=="\r"){break}
		//fmt.Println(element)
		var1:=strings.Split(element,":")
		//fmt.Println(len(var1))
		//fmt.Println(var1[0]," ---->",var1[1])
		self.parmetersmap[var1[0]]=strings.Trim(var1[1]," ")
		}

	//fmt.Println(self.parmetersmap["hostname"])
	//for k,v := range self.parmetersmap{
		//fmt.Println(k)
		//fmt.Println(v)
	//}
	}
func (self *Request)getParameter() map[string]string  {
	return self.parmetersmap
}




//resp
type Response struct {

	responseImp
	staticPath string
	resq Request
	_conn net.Conn
	}

func  (self *Response)redictHtml(htmlName string){
	var headers =
		"HTTP/1.1 200 OK\n" +
			"Date: Sat, 31 Dec 2005 23:59:59 GMT\n" +
			"Content-Type: text/html;charset=UTF-8\n" +
			"Content-Length: 121\n"+"\n"

	htmlbody := HtmlBody(self.staticPath,htmlName)
	allInfo:=headers+htmlbody
	fmt.Println(allInfo)
	self._conn.Write([]byte(headers+htmlbody))

}
func (self *Response)printToWeb(sth string){
	var headers =
		"HTTP/1.1 200 OK\n" +
			"Date: Sat, 31 Dec 2005 23:59:59 GMT\n" +
			"Content-Type: text/html;charset=ISO-8859-1\n" +
			"Content-Length: 121\n" +
			"\n" +
			"<html>\n" +
			"<head>\n" +
			"<title>"+sth+"</title>\n" +
			"</head>\n" +
			"<body>\n" +
			"<!-- body goes here -->\n" +
				sth+
			"</body>\n" +
			"</html>"
			//print(headers)
	self._conn.Write([]byte(headers))
}
func checkErr(e error){
	if e != nil{
		panic(e)
	}
}


//func
type handlers struct {
	handlers map[string]func(*Response ,*Request)
}



type BootStrap struct {
	myhandlers handlers
	parseRe		parseResult
	//startServer(addr string)
}

func (self *BootStrap)Init(path string )  *BootStrap{
	self.myhandlers = handlers{}
	self.myhandlers.handlers = make(map[string]func(*Response ,*Request))
	self.parseRe = parseResult{}
	self.parseRe = *XmlParse(path)
	return self
	}

func (self *BootStrap)Register(partten string, myhandler func(response *Response ,request *Request) ) *BootStrap{
	self.myhandlers.handlers[partten] = myhandler
	return self
	}

func (self*BootStrap)StartServer(){
	addr:=self.parseRe.ip+":"+strconv.Itoa(self.parseRe.port)

	//http.ListenAndServe(addr,nil)
	tcpAddr,err:=net.ResolveTCPAddr("tcp4",addr)
	checkErr(err)
	listener , err :=net.ListenTCP("tcp",tcpAddr)
	checkErr(err)
	for{
		conn,err :=listener.Accept()
		if err != nil{
			fmt.Println(err)
			continue
		}
		go exec_handler(conn,self)
	}
	}


func exec_handler(conn net.Conn,selfBoot *BootStrap){

	defer conn.Close()
	var req = new(Request)
	var resp = new(Response)
	resp._conn = conn
	resp.staticPath = selfBoot.parseRe.staticpath
	var buf[512] byte
	//fmt.Println()
	n,err := conn.Read(buf[0:])
	checkErr(err)
	//rAddr:= conn.RemoteAddr()
	//fmt.Println( string(buf[0:n]))
	req.input_all = string(buf[0:n])
	inputs:= strings.Split(string(buf[0:n]), "\n")
	var tt = strings.Split(inputs[0]," ")
	req.method =tt[0]
	req.uri = tt[1]
	req.httpversion = tt[2]
	req.parseInput(inputs)
	var uriKey = strings.TrimPrefix(req.getUri(),"/")
	 doNow :=selfBoot.myhandlers.handlers[uriKey]
	if (doNow!=nil){
	doNow(resp,req)}
	 //fmt.Println(rAddr.String())
	//headerInput,bodyInput := input[0],input[1]
	//fmt.Println("Header",headerInput)
	//fmt.Println("body",bodyInput)
	//for i:=1;i<len(inputs);i++{
	//	fmt.Println("===>",inputs[i])
	//}
	}

func FinalTest(){
	boot :=new(BootStrap)
	sayHello:= func(response *Response ,request *Request) {

		response.printToWeb("hello")
	}
	lalala:= func(response *Response ,request *Request) {
		//response.printToWeb("lalala")
		response.redictHtml("myhtml.html")

		}
	boot.Init("/home/zw/go/src/GoFirst/goProject/tomcat/study.xml").Register("xxx",sayHello).Register("lalala",lalala).StartServer()
	}