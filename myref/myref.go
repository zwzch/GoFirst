package myref

import (
	"fmt"
	"strconv"
	//"reflect"
	"reflect"
	//"unsafe"
)

type person struct {
	int
	name string
	}
	func(p person)Dis(){
		fmt.Println(p)
	}
	type student struct {
		person
		age int
		Sex int
	}
	func (s student)Sayhello(toname string)(string,int){ //如果方法不是大写第一个字母的公开方法，是无法被反射获取的

		return s.name+" say hello to "+toname,1
	}

	func (s *student)Dis(){ //*student不属于student 仅属于*student对象
		fmt.Println(s)
	}
	func prints(i int) string {
		fmt.Println("i =",i)
		return strconv.Itoa(i)
	}






	/*
	没有运行行期类型对象,实例也没有附加字段用用来表明身身份。只有转换成接口口时,才会在其
	itab 内部存储与该类型有关的信息,Reflect 所有操作都依赖于此。
	*/
	//type User struct {
	//	Username string
	//}
	//type Admin struct {
	//	User
	//	title string
	//}


type User struct {
	Username string
	age int
}
type Admin struct {
	User
	title string
}

type User2 struct {
	Name string `field:"username" type:"nvarchar(20)"`
	Age int `field:"age" type:"tinyint"`
}
func (*User)ToString()  {}
func (Admin)test(){}
var (
	Int= reflect.TypeOf(0)
	String = reflect.TypeOf("")
)



type Data struct {
}
func (*Data) String() string {
	return ""
}
func (*Data)Test(x,y int)(int,int){
	return x+100,y+100
}
//func (*Data)Sum(s string
//)
func refTest(){


	//var u Admin
	//t := reflect.TypeOf(u)
	//for i, n := 0, t.NumField(); i < n; i++ {
	//	f := t.Field(i)
	//	fmt.Println(f.Name, f.Type)
	//	//名称和类型
	//
	//}

	//同样,value-interface 和 pointer-interface 也会导致方方法集存在差异。
	//var u Admin
	//methods := func(t reflect.Type) {
	//	for i, n := 0, t.NumMethod(); i < n; i++ {
	//		m := t.Method(i)
	//		fmt.Println(m.Name)
	//	}
	//}
	//fmt.Println("--- value interface ---")
	//methods(reflect.TypeOf(u))
	//fmt.Println("--- pointer interface ---")
	//methods(reflect.TypeOf(&u))


	//var u Admin
	//t := reflect.TypeOf(u)
	//f, _ := t.FieldByName("title")
	//fmt.Println(f.Name)
	//f, _ = t.FieldByName("User")
	//// 访问嵌入入字段。
	//fmt.Println(f.Name)
	//
	//f, _ = t.FieldByName("Username")
	//// 直接访问嵌入入字段成员,会自自动深度查找。
	//fmt.Println(f.Name)
	//f = t.FieldByIndex([]int{0, 1})
	//// Admin[0] -> User[1] -> age
	//fmt.Println(f.Name)

	//字段标签可实现简单元数据编程,比比如标记 ORM Model 属性。


	//var u User2
	//t := reflect.TypeOf(u)
	//f, _ := t.FieldByName("Name")
	//fmt.Println(f.Tag)
	//fmt.Println(f.Tag.Get("field"))
	//fmt.Println(f.Tag.Get("type"))
	//获取元数据的标签



	//c := reflect.ChanOf(reflect.SendDir, String)
	//fmt.Println(c)
	//m := reflect.MapOf(String, Int)
	//fmt.Println(m)
	//s := reflect.SliceOf(Int)
	//fmt.Println(s)
	//t := struct{ Name string }{}
	//p := reflect.PtrTo(reflect.TypeOf(t))
	//fmt.Println(p)

	//var d *Data
	//t := reflect.TypeOf(d)
	//// 没法直接获取接口类型,好在接口本身身是个 struct,创建
	//// 一个空指针对象,这样传递给 TypeOf 转换成 interface{}
	//// 时就有类型信息了。。
	//it := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	//fmt.Println(t.Implements(it))


	//u := &Admin{User{"Jack", 23}, "NT"}
	//v := reflect.ValueOf(u).Elem()
	//fmt.Println(v.FieldByName("title").String()) // 用用转换方方法获取字段值
	//fmt.Println(v.FieldByName("age").Int()) // 直接访问嵌入入字段成员
	//fmt.Println(v.FieldByIndex([]int{0, 1}).Int()) // 用用多级序号访问嵌入入字段成员

	/*
	除具体的 Int、String 等转换方方法,还可返回 interface{}。只是非非导出字段无无法使用用,需
	用 CanInterface 判断一一下。
	*/
	//u :=User{"Jack",23}
	//v := reflect.ValueOf(u)
	//f := v.FieldByName("age")
	//if f.CanInterface() {
	//	fmt.Println(f.Interface())
	//} else {
	//	fmt.Println(f.Int())
	//}


	//v := reflect.ValueOf([]int{1, 2, 3})
	//for i, n := 0, v.Len(); i < n; i++ {
	//	fmt.Println(v.Index(i).Int())
	//}
	//fmt.Println("---------------------------")
	//v = reflect.ValueOf(map[string]int{"a": 1, "b": 2})
	//for _, k := range v.MapKeys() {
	//	fmt.Println(k.String(), v.MapIndex(k).Int())
	//}
	/*
	需要注意,Value 某些方方法没有遵循 "comma ok" 模式,而而是返回 ZeroValue,因此需
	要用用 IsValid 判断一一下是否可用用。
	*/
	//var p *int
	//var x interface{} = p
	//fmt.Println(x == nil)
	//v := reflect.ValueOf(p)
	//fmt.Println(v.Kind(), v.IsNil())

	//将对象转换为接口口,会发生生复制行行为。该复制品只读,无无法被修改。所以要通过接口口改变
	//目标对象状态,必须是 pointer-interface。
	//就算是指针,我们依然没法将这个存储在 data 的指针指向其他对象,只能透过它修改⺫目目
	//标对象。因为⺫目目标对象并没有被复制,被复制的只是指针

	//u := User{"Jack", 23}
	//v := reflect.ValueOf(u)
	//p := reflect.ValueOf(&u)
	//fmt.Println(v.CanSet(), v.FieldByName("Username").CanSet())
	//fmt.Println(p.CanSet(), p.Elem().FieldByName("Username").CanSet())


	//u := User{"Jack", 23}
	//p := reflect.ValueOf(&u).Elem()
	//p.FieldByName("Username").SetString("Tom")
	//f := p.FieldByName("age")
	//fmt.Println(f.CanSet())
	//// 判断是否能获取地址。
	//if f.CanAddr() {
	//	age := (*int)(unsafe.Pointer(f.UnsafeAddr()))
	//	// age := (*int)(unsafe.Pointer(f.Addr().Pointer()))
	//	// 等同
	//	*age = 88
	//}
	//// 注意 p 是 Value 类型,需要还原成接口口才能转型。
	//fmt.Println(u, p.Interface().(User))
	//先判断是否可以设置值

	//s:=make([]int, 0, 10)
	//v:=reflect.ValueOf(&s).Elem()
	//v.SetLen(2)
	//v.Index(0).SetInt(100)
	//v.Index(1).SetInt(200)
	//fmt.Println(v.Interface(),s)
	//v2:=reflect.Append(v,reflect.ValueOf(300))
	//v2 = reflect.AppendSlice(v2, reflect.ValueOf([]int{400, 500}))
	//fmt.Println(v2.Interface())

	//m:=map[string]int{"a":1}
	//v = reflect.ValueOf(&m).Elem()
	//v.SetMapIndex(reflect.ValueOf("a"), reflect.ValueOf(100)) // update
	//v.SetMapIndex(reflect.ValueOf("b"), reflect.ValueOf(200)) // add
	//fmt.Println(v.Interface(), m)

	}

	func Myref(){
		/*
			Myref
		*/
		//s:=student{person:person{int:1,name:"aaa"},age:22}
		//t:=reflect.TypeOf(s)
		//v:=reflect.ValueOf(s)
		//属性和值是分离的
		//fmt.Println("type",t,"value",v)
		//
		//for i:=0;i<t.NumField();i++{
		//	key:=t.Field(i)
		//	val:=v.Field(i)
		//	fmt.Println(key.Name,key.Type,"|",val)
		//}
		//
		//
		//for i:=0;i<t.NumMethod();i++{  //仅能获取非指针对象的引用方法
		//	m:=t.Method(i)
		//	fmt.Println(m.Name,m.Type)
		//}
		//只有非指针对象

		//指针对象获取的是地址的值
		//s:=student{person:person{int:1,name:"aaa"},age:22}
		//t1:=reflect.TypeOf(&s)
		//v1:=reflect.ValueOf(&s)
		//fmt.Println(t1)
		//fmt.Println(v1)
		//if k:=t1.Kind();k==reflect.Struct{
		//	for i:=0;i<t1.NumField();i++{
		//		key:=t1.Field(i)
		//		val:=v1.Field(i)
		//		fmt.Println(key.Name,key.Type,"|",val)
		//	}
		//}
		//
		//for i:=0;i<t1.NumMethod();i++{ //可以获取所有方法，但是不能用指针获取属性
		//	m:=t1.Method(i)
		//	fmt.Println(m.Name,m.Type)
		//}



		//s:=student{person:person{int:1,name:"aaa"},age:22}
		//t:=reflect.TypeOf(s)
		//v:=reflect.ValueOf(s)
		//fmt.Println(t.FieldByName("person"))
		//fmt.Println(t.FieldByIndex([]int{0}))
		//fmt.Println(t.FieldByIndex([]int{0,0}),t.FieldByIndex([]int{0,1}))
		//m2,_ := t.MethodByName("Sayhello")
		//fmt.Println(m2)

		//x:=123
		//vx:=reflect.ValueOf(&x)
		//vx.Elem().SetInt(333)
		//fmt.Println(x)
		//
		//s11:=student{person:person{int:1,name:"aaa"},age:22,Sex:1}
		//fmt.Println(s11)
		//v11:=reflect.ValueOf(&s11)
		//v11.Elem().FieldByName("age").CanSet()//小写是只读
		//fmt.Println(v11.Elem().FieldByName("Sex").CanSet())//大写才是读写
		//if v11.Elem().FieldByName("age").CanSet(){
		//	v11.Elem().FieldByName("age").SetInt(99)
		//}
		//if v11.Elem().FieldByName("Sex").CanSet(){
		//	v11.Elem().FieldByName("Sex").SetInt(2)
		//}
		//fmt.Println(s11)
		//小写是只读 大写是读写

		//通过传入的参数调用函数
		//fv := reflect.ValueOf(prints)
		//params := make([]reflect.Value,1)
		//params[0] = reflect.ValueOf(20)
		//rs := fv.Call(params)
		//fmt.Println("result:",rs[0].Interface().(string))

		//params1 := make([]reflect.Value,1)  //参数
		//params1[0] = reflect.ValueOf("ppp")   //参数设置为20
		//b := reflect.ValueOf(s11).MethodByName("Sayhello").Call(params1)
		//fmt.Println(b[0],"|",b[1])
		//
		//b1 := reflect.ValueOf(&s11).MethodByName("Dis").Call([]reflect.Value{})
		//fmt.Println(b1)

		//str := "this is string"
		//fmt.Println(reflect.TypeOf(str))

		//a := new(person)
		//a.name = "xxxx"
		//typ := reflect.TypeOf(a)
		//fmt.Println(typ)//指针
		//fmt.Println(typ.Elem())//对象
		//
		//fmt.Println(typ.NumMethod())                   // 1
		//fmt.Println(typ.Method(0))                     // {GetName  func(*myref.person) <func(*myref.person) Value> 0}
		//fmt.Println(typ.Name())                        // ""
		//fmt.Println(typ.PkgPath())                     // ""
		//fmt.Println(typ.Size())                        // 8
		//fmt.Println(typ.String())                      // *myref.person
		//fmt.Println(typ.Elem().String())               // myref.person
		//fmt.Println(typ.Elem().FieldByIndex([]int{0})) // {int go.builtin int  0 [0] true}
		//fmt.Println(typ.Elem().FieldByName("name"))    // {name iotestgo/myref string  8 [1] false} true
		//
		//fmt.Println(typ.Kind() == reflect.Ptr)                              // true
		//fmt.Println(typ.Elem().Kind() == reflect.Struct)


		//mp := make(map[string]int)
		//mp["test1"] = 1
		//fmt.Println(reflect.TypeOf(mp).Key()) //string
		//
		//arr := [1]string{"test"}
		//fmt.Println(arr)
		//fmt.Println(reflect.TypeOf(arr).Len()) // 1
		refTest()
	}