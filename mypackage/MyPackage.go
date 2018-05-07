package mypackage
/*
	编译工工具对源码⺫目目录有严格要求,每个工工作空间 (workspace) 必须由 bin、pkg、src 三
	个目录组成。
	可在 GOPATH 环境变量列表中添加多个工工作空间,但不能和 GOROOT 相同。
	编码:源码文文件必须是 UTF-8 格式,否则会导致编译器出错。
	结束:语句以 ";" 结束,多数时候可以省略。*/
	//注释:支支持 "//"、"/**/" 两种注释方方式,不能嵌套。
	//命名:采用用 camelCasing ⻛风格,不建议使用用下划线。
/*
	所有代码都必须组织在 package 中。
	• 源文文件头部以 "package <name>" 声明包名称。
	• 包由同一目录下的多个源码文文件组成。
	• 包名类似 namespace,与包所在目录名、编译文件名无关。
	• 目录名最好不用用 main、all、std 这三个保留名称。
	• 可执行行文件必须包含 package main,入口函数 main。
*/

/*
	• public: 首首字母母大大写,可被包外访问。
	• internal: 首首字母母小小写,仅包内成员可以访问。
*/
import (
	//"fmt"
	//"net/http"

)
//func handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, `<meta name="go-import"
//	content="test.com/qyuhen/test git https://github.com/qyuhen/test">`)
//}
//func Hello(){
//	print("Hello")
//}

/*
	初始化函数:
	• 每个源文件都可以定义一个或多个初始化函数。
	• 编译器不保证多个初始化函数执行行次序。
	• 初始化函数在单一线程被调用,仅执行行一次。
	• 初始化函数在包所有全局变量初始化后执行行。
	• 在所有初始化函数结束后才执行行 main。
	• 无法调用用初始化函数。
	因为无法保证初始化函数执行行顺序,因此全局变量应该直接用用 var 初始化。
*/
//var now = time.Now()

//func init()  {
//	fmt.Printf("now: %v\n", now)
//}
//func init() {
//	fmt.Printf("since: %v\n", time.Now().Sub(now))
//}
//函数是后边的
//func main(){
//	fmt.Println("main:", int(time.Now().Sub(now).Seconds()))
//}

//不应该滥用用初始化函数,仅适合完成当前文文件中的相关环境设置。
func MyPackage(){
	//http.HandleFunc("/qyuhen/test", handler)
	//http.ListenAndServe(":80", nil)
	//Hello()
	//fmt.Printf("since: %v\n", time.Now().Sub(now))


	//fmt.Println("init:", int(time.Now().Sub(now).Seconds()))
	//w := make(chan bool)
	//go func() {
	//	time.Sleep(time.Second * 3)
	//	w <- true
	//}()
	//<-w
	//fmt.Println("init:", int(time.Now().Sub(now).Seconds()))


	/*文档*/
	/*
	  扩展工工具 godoc 能自自动提取注释生生成帮助文文档。
	• 仅和成员相邻 (中间没有空行行) 的注释被当做帮助信息。
	• 相邻行会合并成同一段落,用用空行行分隔段落。
	• 缩进表示示格式化文本,比比如示示例代码。
	• 自动转换 URL 为链接。
	• 自动合并多个源码文文件中的 package 文文档。
	• 无法显式 package main 中的成员文文档。
	*/
	/*
	• 建议用用专⻔门的 doc.go 保存 package 帮助信息。
	• 包文文档第一一整句 (中英文文句号结束) 被当做 packages 列表说明。
	*/

	}