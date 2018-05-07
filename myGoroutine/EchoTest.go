package myGoroutine

import (
	//"net"
	//"io"
	"net"
	"time"
	"strings"
	"fmt"
	"bufio"
	"os"
	//"go/ast"
)
func echo(c net.Conn,shout string,delay time.Duration){
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
func handleConn1(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}

func EchoTest(){
	conn, err := net.Dial("tcp", "localhost:8000")
	check(err)
	defer conn.Close()
	go mustCopy(os.Stdout,conn)
	mustCopy(conn,os.Stdin)
}
//func handleConn(c net.Conn) {
//	io.Copy(c, c) // NOTE: ignoring errors
//	c.Close()
//}