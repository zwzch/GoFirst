package ticketMaster

import (
	"net"
	"fmt"
	//"os"
	//"bufio"
	//"strings"
	"log"
	"time"
)


//连接服务器
func connectServer() {
	//接通
	conn, err := net.Dial("tcp", "localhost:5000")
	checkError(err)
	fmt.Println("连接成功！")
	//输入
	//inputReader := bufio.NewReader(os.Stdin)
	//fmt.Println("你是谁？")
	//name, _ := inputReader.ReadString('\n')
	//
	//trimName := strings.Trim(name, "\r\n")
	//conn.Write([]byte(trimName + " 接入了\n "))
	for {
		//fmt.Println("我们来聊天吧！按quit退出")
		//读一行
		fmt.Println("send...")
		var buf [512]byte
		n, err := conn.Write([]byte("Hello server!"))
		checkError(err)
		n, err = conn.Read(buf[0:])
		fmt.Println("Reply from server ", string(buf[0:n]))
		time.Sleep(time.Second*3)
		//input, _ := inputReader.Read(buf[0:])

		////如果quit就退出
		//if trimInput == "quit" {
		//	fmt.Println("再见")
		//	conn.Write([]byte(trimName + " 退出了 "))
		//	return
		//}
		////写出来
		//_, err = conn.Write([]byte(trimName + " says " + trimInput))
	}
}

//检查错误
func checkError(err error) {
	if err != nil {
		log.Fatal("an error!", err.Error())
	}
}

//主函数
func TClient() {
	//连接servser
	connectServer()
}