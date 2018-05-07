package myio

import (
	"flag"
	"fmt"
	"os"
	"io"
	"bufio"
	"io/ioutil"
	"path/filepath"
)
func check(e error) {
	if e != nil {
		panic(e)
	}
}

var path string= "/home/zw/go/src/GoFirst/myio/"
func MyIo(){
	//mycmdline()
	//mywritefile()
	//myreadFile()
	//mycopyfile()
	//  mybufiocopyfile()
	//通过buf进行文件拷贝
	   //ioutil
	   //myioutilcopyfile()
	   //  mygetFilelist(path)
	   MyFile2()

	   }



	func myioutilcopyfile(){

		b, err := ioutil.ReadFile(path+"input")//读文件
		if err != nil { panic(err) }

		err = ioutil.WriteFile(path+"output", b, 0644)//写文件
		if err != nil { panic(err) }
	}


func mybufiocopyfile(){
	fi,err :=os.Open(path+"input")
	if err != nil { panic(err) }
	defer fi.Close()
	r := bufio.NewReader(fi)
	//创建一个读取缓冲流

	fo,err := os.Create(path+"output")
	if err != nil { panic(err) }
	defer fo.Close()
	w := bufio.NewWriter(fo)
	buf := make([]byte , 1024)
	for {
		n,err := r.Read(buf)
		if err != nil && err != io.EOF{panic(err)}
		if n == 0{break}
		//如果没有字节数目就退出
	     if n2,err := w.Write(buf[:n]);err != nil{
	     	panic(err)
		 }else if n2 != n{
		 	panic("error in writing")
		 }
	}
	if err = w.Flush();err!=nil{panic(err)}
}
	func mycopyfile(){
		fi,err :=os.Open(path+"input")
		if err != nil { panic(err) }
		defer fi.Close()
		fo, err := os.Create(path+"output")//创建输出*File
		if err != nil { panic(err) }
		defer fo.Close()
		buf := make([]byte,1024)
		for {
			n, err := fi.Read(buf)//从input.txt读取
			if err != nil && err != io.EOF { panic(err) }
			if n == 0 { break }

			if n2, err := fo.Write(buf[:n]); err != nil {//写入output.txt,直到错误
				panic(err)
			} else if n2 != n {
				panic("error in writing")
			}
		}
		}

func myreadFile()  {
	userFile :="/home/zw/go/src/GoFirst/myio/myfile"
	fin,err := os.Open(userFile)
	defer fin.Close()
	if err != nil{
		fmt.Println(userFile,err)
		return
	}
	buf := make([]byte,1024)//创建一个初始容量为1024的slice,作为缓冲容器
	for{
		//循环读取文件数据到缓冲容器中,返回读取到的个数
		n,_ := fin.Read(buf)
		if 0==n{
			break //如果读到个数为0,则读取完毕,跳出循环
		}
		//从缓冲slice中写出数据,从slice下标0到n,通过os.Stdout写出到控制台
		os.Stdout.Write(buf[:n])
	}
}
func mywritefile(){
	userFile :="/home/zw/go/src/GoFirst/myio/myfile"
	fout,err :=os.Create(userFile)
	defer fout.Close()
	if err != nil{
		fmt.Println(userFile,err)
		return
	}
	fout.WriteString("HelloWorld")
	fout.Write([]byte("abcd!\n"))//强转成byte slice后再写入

}
func mycmdline()  {
	//flagtest1()
	args := os.Args //获取用户输入的所有参数
	if args == nil || len(args) <2{//如果用户没有输入,或参数个数不够,则调用该函数提示用户
		fmt.Println("you name?");
		fmt.Println("you age?");
		return
	}
	//fmt.Println(args[0])//第一个参数编译go的位置
	name := args[1] //获取输入的第一个参数
	age := args[2]  //获取输入的第二个参数
	fmt.Println("your name is:",name,"\nyour age is:",age)
}
func flagtest1(){
	married := flag.Bool("married", false, "Are you married?")
	age := flag.Int("age", 22, "How old are you?")
	name := flag.String("name", "", "What your name?")

	var address string
	flag.StringVar(&address, "address", "GuangZhou", "Where is your address?")

	flag.Parse() //解析输入的参数

	fmt.Println("输出的参数married的值是:", *married)//不加*号的话,输出的是内存地址
	fmt.Println("输出的参数age的值是:", *age)
	fmt.Println("输出的参数name的值是:", *name)
	fmt.Println("输出的参数address的值是:", address)
	}



	func mygetFilelist(path string) {
		err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if ( f == nil ) {
			return err
			}
		if f.IsDir() {
			return nil
			}
			println(path)
			return nil
			})
		if err != nil {
		fmt.Println(err)
		}
	}




	func MyFile2() {

	dat, err := ioutil.ReadFile(path+"myfile")
	check(err)
	fmt.Print(string(dat))

	f, err := os.Open(path+"myfile")
	check(err)

	b1 := make([]byte, 6)
	//utf83个字节
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close()

}